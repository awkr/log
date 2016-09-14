package log

import (
	"fmt"
	"os"
	"path/filepath"
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

var (
	levelNames = map[LogLevel]string{
		LevelDebug: "DEBUG",
		LevelInfo:  "INFO",
		LevelWarn:  "WARN",
		LevelError: "ERROR",
		LevelFatal: "FATAL",
	}

	levelColors = map[LogLevel]string{
		LevelDebug: colorBlue,
		LevelInfo:  colorGreen,
		LevelWarn:  colorYellow,
		LevelError: colorLightRed,
		LevelFatal: colorRed,
	}

	isTerm bool

	logFile *os.File

	shouldLogFolder bool
	logFolderFiles  = make(map[LogLevel]*os.File)
	logFolderFile   *os.File
)

type LogConfig struct {
	MinLevel    LogLevel
	EnableColor bool

	// file and folder are not compatible
	// file has a higher priority. if setted, folder will be ignored
	File string

	Folder         string
	LogFileByLevel bool
}

var (
	cfg *LogConfig
)

func Setup(conf *LogConfig) error {
	if conf == nil {
		conf = &LogConfig{MinLevel: LevelDebug, EnableColor: true} // default
	}

	if cfg != nil {
		conf.EnableColor = cfg.EnableColor
	}

	cfg = conf

	var err error

	if cfg.File != "" {
		if logFile, err = os.OpenFile(cfg.File, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm); err != nil {
			return err
		}
	} else if cfg.Folder != "" {
		if err = os.MkdirAll(cfg.Folder, os.ModePerm); err != nil {
			return err
		}

		if !cfg.LogFileByLevel {
			filename := filepath.Join(cfg.Folder, "log.log")
			if logFolderFile, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm); err != nil {
				return err
			}
		}

		shouldLogFolder = true
	}

	return err
}

func EnableColor(flag bool) {
	if cfg == nil {
		cfg = &LogConfig{MinLevel: LevelDebug}
	}

	cfg.EnableColor = (flag && isTerm)
}

func Debug(msg interface{}) {
	if cfg.MinLevel > LevelDebug {
		return
	}

	doLog(LevelDebug, "%v", msg)
}

func Debugf(format string, v ...interface{}) {
	if cfg.MinLevel > LevelDebug {
		return
	}
	doLog(LevelDebug, format, v...)
}

func Info(msg interface{}) {
	if cfg.MinLevel > LevelInfo {
		return
	}

	doLog(LevelInfo, "%v", msg)
}

func Infof(format string, v ...interface{}) {
	if cfg.MinLevel > LevelInfo {
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
func doLog(level LogLevel, format string, msg ...interface{}) error {
	levelName := levelNames[level]

	if logFile != nil { // log to the specified file
		format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName, format)
		fmt.Fprintf(logFile, format, msg...)
	} else if shouldLogFolder { // log to folder
		if cfg.LogFileByLevel {
			f, ok := logFolderFiles[level]
			if !ok {
				// create log file
				filename := filepath.Join(cfg.Folder, fmt.Sprintf("%s.log", strings.ToLower(levelName)))
				file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
				if err != nil {
					return err
				}

				f = file
				logFolderFiles[level] = f
			}

			format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName, format)
			fmt.Fprintf(f, format, msg...)
		} else {
			if logFolderFile != nil {
				format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName, format)
				fmt.Fprintf(logFolderFile, format, msg...)
			}
		}
	} else { // log to terminal or redirect
		if cfg.EnableColor {
			format = fmt.Sprintf("%s %s%-5s%s %s\n", timestamp(), levelColors[level], levelName, colorWhite, format)
		} else {
			format = fmt.Sprintf("%s %-5s %s\n", timestamp(), levelName, format)
		}

		fmt.Printf(format, msg...)
	}

	return nil
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
