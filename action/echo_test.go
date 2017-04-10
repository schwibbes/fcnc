package action

import (
	"github.com/schwibbes/fcnc/util"
	"testing"
)

func TestEcho(t *testing.T) {
	action := EchoAction{"hello"}
	res := action.Execute()
	util.AssertTrue(res.Success, t)
}
