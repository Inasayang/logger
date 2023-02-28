package log

import (
	"strings"
	"testing"
)

func TestLog(t *testing.T) {
	Init("./", "test", "warn")
	Infof("info")
	Debugf("debug")
	Errorf("err")
	Warnf("warn")
	Panicf("panic")
}
func TestA(t *testing.T) {
	a := "./"
	t.Log(strings.HasSuffix(a, "/"))
}
