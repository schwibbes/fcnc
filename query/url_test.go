package query

import (
	. "github.com/schwibbes/fcnc/util"
	"testing"
)

func TestUrl(t *testing.T) {
	res := Url{"http://www.google.de/robots.txt"}.Get()
	t.Logf("", res)
	AssertTrue(res.Ok, t)
}
