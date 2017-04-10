package query

type EchoQuery struct {
	Ok  bool
	Msg string
}

func (self EchoQuery) Get() QueryResult {
	return QueryResult{self.Ok, self.Msg}
}

func NewEchoQuery(tokens []string) EchoQuery {
	if len(tokens) != 2 {
		panic("Echo needs 1 parameter!")
	}

	return EchoQuery{true, tokens[1]}
}
