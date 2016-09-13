package log

import "testing"

func TestLog(t *testing.T) {
	SetLevel(LevelInfo)

	Debug("hello debug")
	Info("hello info")

	Errorf("hello %s", "error")

	EnableColor(false)

	Fatal("oops: color diabled")

	Info("hello info 2")
}

func TestLogFile(t *testing.T) {
	SetLogFile("test.log")

	Debug("i'm debug")
	Info("i'm info")
	Errorf("-->%s<--", "i'm errorf")
}
