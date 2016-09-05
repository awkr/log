package log

import "testing"

func TestLog(t *testing.T) {
	SetLevel(LevelInfo)

	Debug("hello debug")
	Info("hello info")

	Errorf("hello %s", "error")

	Fatal("oops")

	Info("hello info 2")
}
