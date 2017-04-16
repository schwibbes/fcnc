package query

type QueryResult struct {
	Ok  bool
	Val string
}

var (
	factory = make(map[string]QueryBuilder)
)

func init() {
	factory["echo"] = func(tokens []string) Query {
		return NewEchoQuery(tokens)
	}
	factory["state"] = func(tokens []string) Query {
		return NewStateQuery(tokens)
	}
}

type QueryBuilder func(tokens []string) Query

type Query interface {
	Get() QueryResult
}

func NewQuery(tokens []string) Query {
	key := tokens[0]
	return factory[key](tokens)
}
