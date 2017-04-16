package task

import (
	"github.com/schwibbes/fcnc/action"
	"github.com/schwibbes/fcnc/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestSimpleTask(t *testing.T) {

	b, err := ioutil.ReadFile("simple.task")
	require.Empty(t, err)

	result, err := load(b)
	require.Equal(t, 1, len(result))

	// task1
	task1 := result[0]
	log.Printf("task-1 -> %v", result[0])

	// task1.%var
	require.Equal(t, 1, len(task1.Vars))
	assert.Equal(t, "/tmp/foo", task1.Vars["path"])

	// task1.?run
	require.Equal(t, 2, len(task1.Conditions))
	cond1 := task1.Conditions[0].(query.StateQuery)
	assert.Equal(t, query.STARTUP, cond1.Expected)
	cond2 := task1.Conditions[1].(query.EchoQuery)
	assert.Equal(t, "foo", cond2.Msg)

	// task1.!run
	require.Equal(t, 1, len(task1.Actions))
	act1 := task1.Actions[0].(action.EchoAction)
	require.Equal(t, "bar", act1.Msg)

	// task1.?asserts
	require.Equal(t, 1, len(task1.Asserts))
	assert1 := task1.Asserts[0].(query.EchoQuery)
	require.Equal(t, "string in quotes", assert1.Msg)

}

func testSampleTask(t *testing.T) {

	b, err := ioutil.ReadFile("sample.task")
	require.Empty(t, err)

	result, err := load(b)
	require.Equal(t, 3, len(result))

	// task1
	task1 := result[0]
	log.Printf("task-1 -> %v", result[0])
	log.Printf("task-2 -> %v", result[1])
	log.Printf("task-3 -> %v", result[2])

	// task1.%var
	require.Equal(t, 1, len(task1.Vars))
	require.Equal(t, "/tmp/foo", task1.Vars["path"])

	// task1.?run
	require.Equal(t, 2, len(task1.Conditions))
	require.Equal(t, task1.Conditions[0], "state 'STARTUP'")
	require.Equal(t, task1.Conditions[1], "path-exists /tmp/foo")

	// task1.!run
	require.Equal(t, 1, len(task1.Actions))
	//require.Equal(t, task1.Actions[0], "clean-dir /tmp/foo")

	// task1.?asserts
	require.Equal(t, 0, len(task1.Asserts))

}
