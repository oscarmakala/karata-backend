package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"quana.co.tz/karata/api"
	"quana.co.tz/karata/core"
	"quana.co.tz/karata/events"
	"time"
)

const (
	moduleName = "karata"

	tickRate = 5

	maxEmptySec = 30

	delayBetweenGamesSec = 5
	turnTimeFastSec      = 10
	turnTimeNormalSec    = 20
)

// Compile-time check to make sure all required functions are implemented.
var _ runtime.Match = &MatchHandler{}

type MatchLabel struct {
	Open                int  `json:"open"`
	IsPrivate           bool `json:"isPrivate"`
	RequiredPlayerCount int  `json:"requiredPlayerCount"`
	PlayerCount         int  `json:"playerCount"`
}

type MatchHandler struct{}

type MatchState struct {
	label      *MatchLabel
	emptyTicks int
	// Number of users currently in the process of connecting to the match.
	joinsInProgress int
	// True if there's a game currently in progress.
	playing bool
	// Ticks until the next game starts, if applicable.
	nextGameRemainingTicks int64
	presences              map[string]runtime.Presence
	requiredPlayerCount    int
	board                  core.Board
	winner                 string
	GameHandler            core.GameHandlerImpl
	// Ticks until they must submit their move.
	deadlineRemainingTicks int64
}

func (ms *MatchState) ConnectedCount() int {
	count := 0
	for _, p := range ms.presences {
		if p != nil {
			count++
		}
	}
	return count
}

func (m *MatchHandler) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	//create nakama service instance

	////create a game service instance with Nakama service as a dependency
	service := &core.GameHandlerImpl{}

	isPrivate, ok := params["isPrivate"].(bool)
	if !ok {
		logger.Error("invalid match init parameter \"isPrivate\"")
		return nil, 0, ""
	}
	label := &MatchLabel{
		Open:                1,
		IsPrivate:           isPrivate,
		RequiredPlayerCount: 2,
		PlayerCount:         0,
	}

	labelJSON, err := json.Marshal(label)
	if err != nil {
		logger.WithField("error", err).Error("match init failed")
		labelJSON = []byte("{}")
	}

	state := &MatchState{
		label:               label,
		presences:           make(map[string]runtime.Presence),
		requiredPlayerCount: 2,
		GameHandler:         *service,
	}

	return state, tickRate, string(labelJSON)
}

func (m *MatchHandler) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	logger.Debug("MatchJoinAttempt entered->()")
	s := state.(*MatchState)

	// Check if it's a user attempting to rejoin after a disconnect.
	if presence, ok := s.presences[presence.GetUserId()]; ok {
		if presence == nil {
			s.joinsInProgress++
			return s, true, ""
		} else {
			// User attempting to join from 2 different devices at the same time.
			return s, false, "already joined"
		}
	}

	// Check if match is full.
	if len(s.presences)+s.joinsInProgress >= s.requiredPlayerCount {
		return s, false, "match full"
	}
	// New player attempting to connect.
	s.joinsInProgress++
	return s, true, ""
}

func (m *MatchHandler) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	logger.Debug("MatchJoin entered->()")
	s := state.(*MatchState)
	for _, presence := range presences {
		s.emptyTicks = 0
		s.presences[presence.GetUserId()] = presence
		s.joinsInProgress--

		// Check if we must send a message to this user to update them on the current game state.
		var opCode events.EventType
		var msg api.Message
		if s.playing {
			// There's a game still currently in progress, the player is re-joining after a disconnect. Give them a state update.
			opCode = events.EventTypeUpdate
			msg = &api.UpdateMessage{
				Deck:        s.board.Deck,
				CurrentCard: s.board.CurrentCard,
				CurrentTurn: s.board.ActivePlayerId,
				GameOver:    false,
			}
		} else if s.board.DropZone != nil {
			// There's no game in progress but we still have a completed game that the user was part of.
			// They likely disconnected before the game ended, and have since forfeited because they took too long to return.
			opCode = events.EventTypeNextTurn
			msg = &api.DoneMessage{
				Winner: s.winner,
			}
		}
		// Send a message to the user that just joined, if one is needed based on the logic above.
		if msg != nil {
			buf, err := json.Marshal(msg)
			if err != nil {
				logger.Error("error encoding message: %v", err)
			} else {
				_ = dispatcher.BroadcastMessage(int64(opCode), buf, []runtime.Presence{presence}, nil, true)
			}
		}
	}
	// Check if match was open to new players, but should now be closed.
	if len(s.presences) >= 2 && s.label.Open != 0 {
		s.label.Open = 0
		if labelJSON, err := json.Marshal(s.label); err != nil {
			logger.Error("error encoding label: %v", err)
		} else {
			if err := dispatcher.MatchLabelUpdate(string(labelJSON)); err != nil {
				logger.Error("error updating label: %v", err)
			}
		}
	}
	return s
}

func (m *MatchHandler) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	logger.Debug("MatchLeave entered->()")
	s := state.(*MatchState)
	for _, presence := range presences {
		s.presences[presence.GetUserId()] = nil
	}

	var humanPlayersRemaining []runtime.Presence
	for _, presence := range s.presences {
		if presence != nil {
			humanPlayersRemaining = append(humanPlayersRemaining, presence)
		}
	}
	// Notify remaining player that the opponent has left the game
	if len(humanPlayersRemaining) == 1 {
		_ = dispatcher.BroadcastMessage(int64(events.EventTypeOpponentLeft), nil, humanPlayersRemaining, nil, true)
	}

	return s
}

func (m *MatchHandler) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {

	s := state.(*MatchState)
	if s.ConnectedCount()+s.joinsInProgress == 0 {
		s.emptyTicks++
		if s.emptyTicks >= maxEmptySec*tickRate {
			// Match has been empty for too long, close it.
			logger.Info("closing idle match")
			return nil
		}
	}
	t := time.Now().UTC()

	// If there's no game in progress check if we can (and should) start one!
	if !s.playing {
		// Between games any disconnected users are purged, there's no in-progress game for them to return to anyway.
		for userID, presence := range s.presences {
			if presence == nil {
				delete(s.presences, userID)
			}
		}

		// Check if we need to update the label so the match now advertises itself as open to join.
		if len(s.presences) < s.requiredPlayerCount && s.label.Open != 1 {
			s.label.Open = 1
			if labelJSON, err := json.Marshal(s.label); err != nil {
				logger.Error("error encoding label: %v", err)
			} else {
				logger.Debug("EventTypeUpdate label : %s", string(labelJSON))

				if err := dispatcher.MatchLabelUpdate(string(labelJSON)); err != nil {
					logger.Error("error updating label: %v", err)
				}
			}
		}

		// Check if we have enough players to start a game.
		if len(s.presences) < s.requiredPlayerCount {
			return s
		}

		// Check if enough time has passed since the last game.
		if s.nextGameRemainingTicks > 0 {
			s.nextGameRemainingTicks--
			return s
		}

		// We can start a game! Set up the game state and assign the marks to each player.
		s.playing = true
		players := make([]core.PlayerState, 0)
		for userID, presence := range s.presences {
			players = append(players, core.PlayerState{
				UserID: userID,
				Hand:   make([]core.Card, 0),
				Name:   presence.GetUsername(),
			})
		}

		s.board = s.GameHandler.ActionStartGame(players, 5)
		s.deadlineRemainingTicks = calculateDeadlineTicks()
		s.nextGameRemainingTicks = 0

		// Notify the players a new game has started.
		buf, err := json.Marshal(&api.StartMessage{
			Board:    s.board,
			Deadline: t.Add(time.Duration(s.deadlineRemainingTicks/tickRate) * time.Second).Unix(),
		})

		logger.Debug("Notify players : %s", string(buf))
		if err != nil {
			logger.Error("error encoding message: %v", err)
		} else {
			_ = dispatcher.BroadcastMessage(int64(events.EventTypeStart), buf, nil, nil, true)
		}
		return s
	}

	// There's a game in progress. Check for input, update match state, and send messages to clients.
	for _, message := range messages {
		switch events.EventType(message.GetOpCode()) {
		case events.EventTypeCardPlayed:
			msg := &api.UpdateMessage{}
			err := json.Unmarshal(message.GetData(), msg)
			if err != nil {
				// Client sent bad data.
				_ = dispatcher.BroadcastMessage(int64(events.EventTypeRejected), nil, []runtime.Presence{message}, nil, true)
				continue
			}
			err = s.GameHandler.ActionPlayCard(&s.board, &msg.CardPlayed, message.GetUserId())
			if err != nil {
				logger.Error("error ActionPlayCard: %v", err)
			} else {
				_ = dispatcher.BroadcastMessage(int64(events.EventTypeCardPlayed), message.GetData(), nil, nil, true)
			}
		case events.EventTypeNextTurn:
			if message.GetUserId() != s.board.ActivePlayerId {
				_ = dispatcher.BroadcastMessage(int64(events.EventTypeRejected), nil, []runtime.Presence{message}, nil, true)
				return nil
			}
			err := s.GameHandler.ActionNextTurn(&s.board, message.GetUserId(), s.playing)
			if err != nil {
				logger.Error("error ActionNextTurn: %v", err)
				return nil
			}

			buf, err := json.Marshal(&api.UpdateMessage{
				Deck:        s.board.Deck,
				CurrentCard: s.board.CurrentCard,
				CurrentTurn: s.board.ActivePlayerId,
				Penalties:   s.board.Penalties,
			})
			if err != nil {
				logger.Error("error encoding message: %v", err)
			} else {
				_ = dispatcher.BroadcastMessage(int64(events.EventTypeNextTurn), buf, nil, nil, true)
			}
		case events.EventTypeOpponentLeft:
		default:
			// No other opcodes are expected from the client, so automatically treat it as an error.
			_ = dispatcher.BroadcastMessage(int64(events.EventTypeRejected), nil, []runtime.Presence{message}, nil, true)
		}
	}
	return s
}

func calculateDeadlineTicks() int64 {
	return turnTimeNormalSec * tickRate
}

func (m *MatchHandler) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	return state
}

func (m *MatchHandler) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return state, ""
}
