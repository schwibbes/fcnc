package query

import (
	"strings"
)

type StateQuery struct {
	Expected AppState
}

type AppState uint32

const (
	STARTUP AppState = 1 + iota
	ANY
	SHUTDOWN
)

func (self StateQuery) Get() QueryResult {
	state := STARTUP
	return QueryResult{self.Expected == state, "currentstate"}
}

func NewStateQuery(tokens []string) StateQuery {
	if len(tokens) != 2 {
		panic("State needs exactly one parameter")
	}

	switch strings.ToLower(tokens[1]) {
	case "startup":
		return StateQuery{STARTUP}
	case "any":
		return StateQuery{ANY}
	case "shutdown":
		return StateQuery{SHUTDOWN}
	default:
		panic("unhandled case: " + tokens[1])
	}
}
