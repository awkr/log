package log

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int16

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// 最小level，默认为LevelDebug
var logLevel = LevelDebug

func SetLevel(level LogLevel) {
	logLevel = level
}

func Debug(msg string) {
	if logLevel > LevelDebug {
		return
	}
	fmt.Printf("%s DEBUG %s\n", timestamp(), msg)
}

func Debugf(format string, v ...interface{}) {
	if logLevel > LevelDebug {
		return
	}
	format = fmt.Sprintf("%s DEBUG %s\n", timestamp(), format)
	fmt.Printf(format, v...)
}

func Info(msg string) {
	if logLevel > LevelInfo {
		return
	}
	fmt.Printf("%s %-5s %s\n", timestamp(), levelName(LevelInfo), msg)
}

func Infof(format string, v ...interface{}) {
	if logLevel > LevelInfo {
		return
	}
	format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName(LevelInfo), format)
	fmt.Printf(format, v...)
}

func Warn(msg string) {
	fmt.Printf("%s %-5s %s\n", timestamp(), levelName(LevelWarn), msg)
}

func Warnf(format string, v ...interface{}) {
	format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName(LevelWarn), format)
	fmt.Printf(format, v...)
}

func Error(msg string) {
	fmt.Printf("%s ERROR %s\n", timestamp(), msg)
}

func Errorf(format string, v ...interface{}) {
	format = fmt.Sprintf("%s ERROR %s\n", timestamp(), format)
	fmt.Printf(format, v...)
}

func Fatal(msg string) {
	fmt.Printf("%s FATAL %s\n", timestamp(), msg)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	format = fmt.Sprintf("%s FATAL %s\n", timestamp(), format)
	fmt.Printf(format, v...)
	os.Exit(1)
}

func levelName(level LogLevel) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func timestamp() string {
	return time.Now().Local().Format("2006/01/02 15:04:05.000")
}
