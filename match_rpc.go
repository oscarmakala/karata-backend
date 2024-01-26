package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"quana.co.tz/karata/api"
)

func rpcFindMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	_, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		return "", errNoUserIdFound
	}

	isPrivate := false

	// Get the isPrivate value from the payload if it exists
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		logger.Error("error unmarshaling payload: %v", err.Error())
		return "", errUnmarshal
	}
	if val, ok := data["isPrivate"]; ok {
		isPrivate = val.(bool)
	}

	// List at most 10 matches, not authoritative, and that
	// have between 0 and 1 players currently participating.
	limit := 10
	label := ""
	maxSize := 1

	matchIDs := make([]string, 0, 10)
	matches, err := nk.MatchList(ctx, limit, true, label, nil, &maxSize, "")
	if err != nil {
		logger.Error("error listing matches: %v", err)
		return "", errInternalError
	}

	if len(matches) > 0 {
		// There are one or more ongoing matches the user could join.
		for _, match := range matches {
			matchIDs = append(matchIDs, match.MatchId)
		}
	} else {
		// No available matches found, create a new one.
		params := map[string]interface{}{
			"isPrivate": isPrivate,
		}
		// No available matches found, create a new one.
		matchID, err := nk.MatchCreate(ctx, moduleName, params)
		if err != nil {
			logger.Error("error creating match: %v", err)
			return "", errInternalError
		}
		matchIDs = append(matchIDs, matchID)
	}
	response, err := json.Marshal(&api.RpcFindMatchResponse{MatchIds: matchIDs})
	if err != nil {
		logger.Error("error marshaling response: %v", err.Error())
		return "", errMarshal
	}

	return string(response), nil
}
