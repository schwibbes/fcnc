package action

import (
	"log"
)

const (
	ActionID = "ECHO_ACTION"
)

type EchoAction struct {
	Msg string
}

func (self EchoAction) Execute() ActionResult {
	log.Print(self.Msg)
	return ActionResult{true}
}

func NewEchoAction(tokens []string) EchoAction {
	if len(tokens) != 2 {
		panic("Echo needs 1 parameter!")
	}

	return EchoAction{tokens[1]}
}
