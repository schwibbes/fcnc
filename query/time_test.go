package query

import (
	. "github.com/schwibbes/fcnc/util"
	"testing"
)

func TestTime(t *testing.T) {
	res := time{}.Get()
	t.Logf("", res)
	AssertTrue(res.Ok, t)
}
