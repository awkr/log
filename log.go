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

const (
	colorWhite = "\x1b[0m"

	colorBlue     = "\x1b[34m"
	colorGreen    = "\x1b[32m"
	colorYellow   = "\x1b[33m"
	colorLightRed = "\x1b[91m"
	colorRed      = "\x1b[31m"
)

// 最小level，默认为LevelDebug
var (
	logLevel    = LevelDebug
	enableColor = true

	levelColors = map[LogLevel]string{
		LevelDebug: colorBlue,
		LevelInfo:  colorGreen,
		LevelWarn:  colorYellow,
		LevelError: colorLightRed,
		LevelFatal: colorRed,
	}
)

func SetLevel(level LogLevel) {
	logLevel = level
}

func EnableColor(flag bool) {
	enableColor = flag
}

func Debug(msg string) {
	if logLevel > LevelDebug {
		return
	}
	if enableColor {
		fmt.Printf("%s %sDEBUG%s %s\n", timestamp(), levelColors[LevelDebug], colorWhite, msg)
	} else {
		fmt.Printf("%s DEBUG %s\n", timestamp(), msg)
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel > LevelDebug {
		return
	}
	if enableColor {
		format = fmt.Sprintf("%s %sDEBUG%s %s\n", timestamp(), levelColors[LevelDebug], colorWhite, format)
		fmt.Printf(format, v...)
	} else {
		format = fmt.Sprintf("%s DEBUG %s\n", timestamp(), format)
		fmt.Printf(format, v...)
	}
}

func Info(msg string) {
	if logLevel > LevelInfo {
		return
	}
	if enableColor {
		fmt.Printf("%s %s%-5s%s %s\n", timestamp(), levelColors[LevelInfo], levelName(LevelInfo), colorWhite, msg)
	} else {
		fmt.Printf("%s %-5s %s\n", timestamp(), levelName(LevelInfo), msg)
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel > LevelInfo {
		return
	}
	if enableColor {
		format = fmt.Sprintf("%s %s%-5s%s %s\n", timestamp(), levelColors[LevelInfo], levelName(LevelInfo), colorWhite, format)
	} else {
		format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName(LevelInfo), format)
	}
	fmt.Printf(format, v...)
}

func Warn(msg string) {
	if enableColor {
		fmt.Printf("%s %s%-5s%s %s\n", timestamp(), levelColors[LevelWarn], levelName(LevelWarn), colorWhite, msg)
	} else {
		fmt.Printf("%s %-5s %s\n", timestamp(), levelName(LevelWarn), msg)
	}
}

func Warnf(format string, v ...interface{}) {
	if enableColor {
		format = fmt.Sprintf("%s %s%-5s%s %s\n", timestamp(), levelColors[LevelWarn], levelName(LevelWarn), colorWhite, format)
	} else {
		format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName(LevelWarn), format)
	}
	fmt.Printf(format, v...)
}

func Error(msg string) {
	if enableColor {
		fmt.Printf("%s %sERROR%s %s\n", timestamp(), levelColors[LevelError], colorWhite, msg)
	} else {
		fmt.Printf("%s ERROR %s\n", timestamp(), msg)
	}
}

func Errorf(format string, v ...interface{}) {
	if enableColor {
		format = fmt.Sprintf("%s %sERROR%s %s\n", timestamp(), levelColors[LevelError], colorWhite, format)
	} else {
		format = fmt.Sprintf("%s ERROR %s\n", timestamp(), format)
	}
	fmt.Printf(format, v...)
}

func Fatal(msg string) {
	if enableColor {
		fmt.Printf("%s %sFATAL%s %s\n", timestamp(), levelColors[LevelFatal], colorWhite, msg)
	} else {
		fmt.Printf("%s FATAL %s\n", timestamp(), msg)
	}
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	if enableColor {
		format = fmt.Sprintf("%s %sFATAL%s %s\n", timestamp(), levelColors[LevelFatal], colorWhite, format)
	} else {
		format = fmt.Sprintf("%s FATAL %s\n", timestamp(), format)
	}
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
