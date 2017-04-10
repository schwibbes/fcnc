package task

import (
	"github.com/schwibbes/fcnc/action"
	"github.com/schwibbes/fcnc/query"
)

type Task struct {
	Vars       map[string]string
	Conditions []query.Query
	Actions    []action.Action
	Asserts    []query.Query
}

func (self Task) Run() {
}

func run(actions []action.Action) {
	for _, a := range actions {
		a.Execute()
	}
}

func load(content []byte) (result []Task, err error) {

	p := &Parser{asByes: content, mode: TASK, taskNum: 0}
	p.parse()
	return p.Result, nil
}

func NewTask(decoded string) (result *Task) {
	result = new(Task)
	result.Vars = make(map[string]string)
	result.Actions = make([]action.Action, 0)
	result.Conditions = make([]query.Query, 0)
	result.Asserts = make([]query.Query, 0)
	return
}
