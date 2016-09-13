package log

import "testing"

func TestLog(t *testing.T) {
	Setup(&LogConfig{MinLevel: LevelDebug})

	Debug("hello debug")
	Info("hello info")

	Errorf("hello %s", "error")

	EnableColor(false)

	Fatal("oops: color diabled")

	Info("hello info 2")
}

func TestLogFile(t *testing.T) {
	Setup(&LogConfig{File: "test.log"})

	Debug("i'm debug")
	Info("i'm info")
	Errorf("-->%s<--", "i'm errorf")
}

func TestLogFolder(t *testing.T) {
	Setup(&LogConfig{Folder: "output/log", LogFileByLevel: true})

	Debug("i'm debug")
	Info("i'm info")
	Errorf("-->%s<--", "i'm errorf")
}
