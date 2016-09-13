package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
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

	isTerm bool

	shouldLogFile bool
	logFile       *os.File
)

func SetLevel(level LogLevel) {
	logLevel = level
}

func SetLevelStr(level string) {
	level = strings.ToUpper(level)
	SetLevel(levelFromName(level))
}

func SetLogFile(f string) error {
	shouldLogFile = true

	file, err := os.OpenFile(f, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	logFile = file
	return nil
}

func EnableColor(flag bool) {
	flag = flag && isTerm
	enableColor = flag
}

func Debug(msg interface{}) {
	if logLevel > LevelDebug {
		return
	}

	doLog(LevelDebug, "%v", msg)
}

func Debugf(format string, v ...interface{}) {
	if logLevel > LevelDebug {
		return
	}
	doLog(LevelDebug, format, v...)
}

func Info(msg interface{}) {
	if logLevel > LevelInfo {
		return
	}

	doLog(LevelInfo, "%v", msg)
}

func Infof(format string, v ...interface{}) {
	if logLevel > LevelInfo {
		return
	}
	doLog(LevelInfo, format, v...)
}

func Warn(msg interface{}) {
	doLog(LevelWarn, "%v", msg)
}

func Warnf(format string, v ...interface{}) {
	doLog(LevelWarn, format, v...)
}

func Error(msg interface{}) {
	doLog(LevelError, "%v", msg)
}

func Errorf(format string, v ...interface{}) {
	doLog(LevelError, format, v...)
}

func Fatal(msg interface{}) {
	doLog(LevelFatal, "%v", msg)

	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	doLog(LevelFatal, format, v...)

	os.Exit(1)
}

// format is original format
func doLog(level LogLevel, format string, msg ...interface{}) {
	if shouldLogFile {
		if logFile != nil {
			format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName(level), format)
			fmt.Fprintf(logFile, format, msg...)
		}
	} else {
		if enableColor {
			format = fmt.Sprintf("%s %s%-5s%s %s\n", timestamp(), levelColors[level], levelName(level), colorWhite, format)
		} else {
			format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName(level), format)
		}

		fmt.Printf(format, msg...)
	}
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

func levelFromName(level string) LogLevel {
	switch level {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	case "FATAL":
		return LevelFatal
	default:
		return LevelInfo // if no level matched, set INFO as default
	}
}

func timestamp() string {
	return time.Now().Local().Format("2006/01/02 15:04:05.000")
}

func init() {
	isTerm = terminal.IsTerminal(int(os.Stdout.Fd()))
	EnableColor(isTerm)
}
