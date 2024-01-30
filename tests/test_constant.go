package tests

const TestDeck = `{"Deck":{"Cards":[{"Suit":"Diamonds","Rank":13},{"Suit":"Spades","Rank":4},{"Suit":"Hearts","Rank":6},{"Suit":"Spades","Rank":10},{"Suit":"Hearts","Rank":9},{"Suit":"Clubs","Rank":3},{"Suit":"Spades","Rank":2},{"Suit":"Clubs","Rank":10},{"Suit":"Spades","Rank":1},{"Suit":"Spades","Rank":12},{"Suit":"Hearts","Rank":3},{"Suit":"Spades","Rank":6},{"Suit":"Hearts","Rank":13},{"Suit":"Clubs","Rank":8},{"Suit":"Diamonds","Rank":2},{"Suit":"Clubs","Rank":12},{"Suit":"Diamonds","Rank":7},{"Suit":"Spades","Rank":7},{"Suit":"Spades","Rank":8},{"Suit":"Clubs","Rank":13},{"Suit":"Clubs","Rank":5},{"Suit":"Hearts","Rank":7},{"Suit":"Hearts","Rank":8},{"Suit":"Diamonds","Rank":12},{"Suit":"Hearts","Rank":5},{"Suit":"Clubs","Rank":9},{"Suit":"Diamonds","Rank":10},{"Suit":"Clubs","Rank":11},{"Suit":"Clubs","Rank":2},{"Suit":"Diamonds","Rank":6},{"Suit":"Diamonds","Rank":5},{"Suit":"Diamonds","Rank":8},{"Suit":"Hearts","Rank":2},{"Suit":"Spades","Rank":3},{"Suit":"Diamonds","Rank":9},{"Suit":"Hearts","Rank":1},{"Suit":"Clubs","Rank":7},{"Suit":"Clubs","Rank":6},{"Suit":"Hearts","Rank":11},{"Suit":"Diamonds","Rank":3},{"Suit":"Hearts","Rank":4}]},"CurrentCard":{"Suit":"Diamonds","Rank":4},"NoOfCardsToDeal":5,"DropZone":[],"Players":[{"UserID":"abc","Hand":[{"Suit":"Spades","Rank":9},{"Suit":"Hearts","Rank":10},{"Suit":"Hearts","Rank":12},{"Suit":"Spades","Rank":5},{"Suit":"Clubs","Rank":1}],"Name":"tom"},{"UserID":"efg","Hand":[{"Suit":"Spades","Rank":11},{"Suit":"Clubs","Rank":4},{"Suit":"Diamonds","Rank":1},{"Suit":"Spades","Rank":13},{"Suit":"Diamonds","Rank":11}],"Name":"Mark"}],"ActivePlayerId":"abc","PlayerOrder":["abc","efg"],"Penalties":[],"CanDefend":false,"DrawnCards":[]}`
const KTestNextTurnWhenCardPlayedIsPlus2NoDefense = `{"Deck":{"Cards":[{"Suit":"Diamonds","Rank":4},{"Suit":"Clubs","Rank":8},{"Suit":"Clubs","Rank":9},{"Suit":"Hearts","Rank":2},{"Suit":"Diamonds","Rank":3},{"Suit":"Diamonds","Rank":10},{"Suit":"Clubs","Rank":13},{"Suit":"Clubs","Rank":1},{"Suit":"Hearts","Rank":8},{"Suit":"Clubs","Rank":7},{"Suit":"Spades","Rank":6},{"Suit":"Diamonds","Rank":11},{"Suit":"Hearts","Rank":1},{"Suit":"Diamonds","Rank":5},{"Suit":"Hearts","Rank":7},{"Suit":"Hearts","Rank":13},{"Suit":"Hearts","Rank":9},{"Suit":"Hearts","Rank":11},{"Suit":"Diamonds","Rank":12},{"Suit":"Diamonds","Rank":7},{"Suit":"Spades","Rank":2},{"Suit":"Spades","Rank":10},{"Suit":"Clubs","Rank":4},{"Suit":"Clubs","Rank":11},{"Suit":"Diamonds","Rank":6},{"Suit":"Spades","Rank":12},{"Suit":"Diamonds","Rank":13},{"Suit":"Hearts","Rank":12},{"Suit":"Spades","Rank":3},{"Suit":"Spades","Rank":9},{"Suit":"Clubs","Rank":5},{"Suit":"Spades","Rank":11},{"Suit":"Hearts","Rank":3},{"Suit":"Spades","Rank":7},{"Suit":"Clubs","Rank":10},{"Suit":"Diamonds","Rank":1},{"Suit":"Spades","Rank":1},{"Suit":"Spades","Rank":13},{"Suit":"Hearts","Rank":5}]},"CurrentCard":{"Suit":"Diamonds","Rank":2},"NoOfCardsToDeal":5,"DropZone":[{"Suit":"Clubs","Rank":2}],"Players":[{"UserID":"abc","Hand":[{"Suit":"Diamonds","Rank":8},{"Suit":"Clubs","Rank":12},{"Suit":"Hearts","Rank":6},{"Suit":"Spades","Rank":4},{"Suit":"Spades","Rank":8}],"Name":"tom"},{"UserID":"efg","Hand":[{"Suit":"Hearts","Rank":10},{"Suit":"Hearts","Rank":4},{"Suit":"Diamonds","Rank":9},{"Suit":"Clubs","Rank":6},{"Suit":"Spades","Rank":5}],"Name":"Mark"}],"ActivePlayerId":"efg","PlayerOrder":null,"Penalties":[],"CanDefend":false,"DrawnCards":[],"PlayerTurnDirection":1}`
const KTestNextTurnWhenCardPlayedIsPlus4NoDefense = `{"Deck":{"Cards":[{"Suit":"Diamonds","Rank":4},{"Suit":"Clubs","Rank":8},{"Suit":"Clubs","Rank":9},{"Suit":"Hearts","Rank":2},{"Suit":"Diamonds","Rank":3},{"Suit":"Diamonds","Rank":10},{"Suit":"Clubs","Rank":13},{"Suit":"Clubs","Rank":1},{"Suit":"Hearts","Rank":8},{"Suit":"Clubs","Rank":7},{"Suit":"Spades","Rank":6},{"Suit":"Diamonds","Rank":11},{"Suit":"Hearts","Rank":1},{"Suit":"Diamonds","Rank":5},{"Suit":"Hearts","Rank":7},{"Suit":"Hearts","Rank":13},{"Suit":"Hearts","Rank":9},{"Suit":"Hearts","Rank":11},{"Suit":"Diamonds","Rank":12},{"Suit":"Diamonds","Rank":7},{"Suit":"Spades","Rank":2},{"Suit":"Spades","Rank":10},{"Suit":"Clubs","Rank":4},{"Suit":"Clubs","Rank":11},{"Suit":"Diamonds","Rank":6},{"Suit":"Spades","Rank":12},{"Suit":"Diamonds","Rank":13},{"Suit":"Hearts","Rank":12},{"Suit":"Spades","Rank":3},{"Suit":"Spades","Rank":9},{"Suit":"Clubs","Rank":5},{"Suit":"Spades","Rank":11},{"Suit":"Hearts","Rank":3},{"Suit":"Spades","Rank":7},{"Suit":"Clubs","Rank":10},{"Suit":"Diamonds","Rank":1},{"Suit":"Spades","Rank":1},{"Suit":"Spades","Rank":13},{"Suit":"Hearts","Rank":5}]},"CurrentCard":{"Suit":"Diamonds","Rank":2},"NoOfCardsToDeal":5,"DropZone":[{"Suit":"Clubs","Rank":2},{"Suit":"Diamonds","Rank":2}],"Players":[{"UserID":"abc","Hand":[{"Suit":"Diamonds","Rank":8},{"Suit":"Clubs","Rank":12},{"Suit":"Hearts","Rank":6},{"Suit":"Spades","Rank":4},{"Suit":"Spades","Rank":8}],"Name":"tom"},{"UserID":"efg","Hand":[{"Suit":"Hearts","Rank":10},{"Suit":"Hearts","Rank":4},{"Suit":"Diamonds","Rank":9},{"Suit":"Clubs","Rank":6},{"Suit":"Spades","Rank":5}],"Name":"Mark"}],"ActivePlayerId":"efg","PlayerOrder":null,"Penalties":[],"CanDefend":false,"DrawnCards":[],"PlayerTurnDirection":1}`
const KTestDrawAfterPenalty = `{"Deck":{"Cards":[{"Suit":"Diamonds","Rank":4},{"Suit":"Clubs","Rank":8},{"Suit":"Clubs","Rank":9},{"Suit":"Hearts","Rank":2},{"Suit":"Diamonds","Rank":3},{"Suit":"Diamonds","Rank":10},{"Suit":"Clubs","Rank":13},{"Suit":"Clubs","Rank":1},{"Suit":"Hearts","Rank":8},{"Suit":"Clubs","Rank":7},{"Suit":"Spades","Rank":6},{"Suit":"Diamonds","Rank":11},{"Suit":"Hearts","Rank":1},{"Suit":"Diamonds","Rank":5},{"Suit":"Hearts","Rank":7},{"Suit":"Hearts","Rank":13},{"Suit":"Hearts","Rank":9},{"Suit":"Hearts","Rank":11},{"Suit":"Diamonds","Rank":12},{"Suit":"Diamonds","Rank":7},{"Suit":"Spades","Rank":2},{"Suit":"Spades","Rank":10},{"Suit":"Clubs","Rank":4},{"Suit":"Clubs","Rank":11},{"Suit":"Diamonds","Rank":6},{"Suit":"Spades","Rank":12},{"Suit":"Diamonds","Rank":13},{"Suit":"Hearts","Rank":12},{"Suit":"Spades","Rank":3},{"Suit":"Spades","Rank":9},{"Suit":"Clubs","Rank":5},{"Suit":"Spades","Rank":11},{"Suit":"Hearts","Rank":3},{"Suit":"Spades","Rank":7},{"Suit":"Clubs","Rank":10},{"Suit":"Diamonds","Rank":1},{"Suit":"Spades","Rank":1},{"Suit":"Spades","Rank":13},{"Suit":"Hearts","Rank":5}]},"CurrentCard":{"Suit":"Diamonds","Rank":2},"NoOfCardsToDeal":5,"DropZone":[],"Players":[{"UserID":"abc","Hand":[{"Suit":"Diamonds","Rank":8},{"Suit":"Clubs","Rank":12},{"Suit":"Hearts","Rank":6},{"Suit":"Spades","Rank":4},{"Suit":"Spades","Rank":8}],"Name":"tom"},{"UserID":"efg","Hand":[{"Suit":"Hearts","Rank":10},{"Suit":"Hearts","Rank":4},{"Suit":"Diamonds","Rank":9},{"Suit":"Clubs","Rank":6},{"Suit":"Spades","Rank":5}],"Name":"Mark"}],"ActivePlayerId":"efg","PlayerOrder":null,"Penalties":[{"UserId":"abc","NumberOfCards":4,"Penalty":"Draw"}],"CanDefend":false,"DrawnCards":[],"PlayerTurnDirection":1}`