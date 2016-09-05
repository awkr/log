package log

import "testing"

func TestLog(t *testing.T) {
	SetLevel(LevelInfo)

	Debug("hello debug")
	Info("hello info")

	Errorf("hello %s", "error")

	EnableColor(false)

	Fatal("oops")

	Info("hello info 2")
}
