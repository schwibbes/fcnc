package action

import ()

type ActionResult struct {
	Success bool
}

type Action interface {
	Execute() ActionResult
}

var (
	factory = make(map[string]ActionBuilder)
)

func init() {
	factory["echo"] = func(tokens []string) Action {
		return NewEchoAction(tokens)
	}
}

type ActionBuilder func(tokens []string) Action

func NewAction(tokens []string) Action {
	return factory[tokens[0]](tokens)
}
