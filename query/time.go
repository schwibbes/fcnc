package query

import (
	tm "time"
)

type time struct {
}

func (t time) Get() QueryResult {

	return QueryResult{true, string(tm.Now().Format("2006-Jan-02 15:30:16"))}
}
