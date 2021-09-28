package glogger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	ErrorLevel LogLevel = iota
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

var log *Logger = &Logger{timeFormat: time.RFC3339, writer: os.Stdout}

type Logger struct {
	m          sync.RWMutex
	prefix     string
	logLevel   LogLevel
	timeFormat string
	writer     io.Writer
}

type LogLevel uint

func logLevel(level LogLevel) string {
	switch level {
	case ErrorLevel:
		return "ERROR"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBUG"
	case TraceLevel:
		return "TRACE"
	default:
		return "UNKNOWN"
	}
}

func SetLogLevel(level LogLevel) {
	log.m.Lock()
	defer log.m.Unlock()

	log.logLevel = level
}

func SetPrefix(prefix string) {
	log.m.Lock()
	defer log.m.Unlock()

	log.prefix = prefix + " "
}

func SetTimeFormat(timeFormat string) {
	log.m.Lock()
	defer log.m.Unlock()

	log.timeFormat = timeFormat
}

func SetWriter(writer io.Writer) {
	log.m.Lock()
	defer log.m.Unlock()

	log.writer = writer
}

func Info(args ...interface{}) {
	if log.logLevel >= InfoLevel {
		fmt.Fprintln(log.writer, fmtLog(InfoLevel, args...))
	}
}

func Infof(format string, args ...interface{}) {
	if log.logLevel >= InfoLevel {
		fmt.Fprintln(log.writer, fmtLogf(InfoLevel, format, args...))
	}
}

func Panic(args ...interface{}) {
	s := fmtLog(ErrorLevel, args...)
	fmt.Fprintln(log.writer, s)
	panic(s)
}

func Panicf(format string, args ...interface{}) {
	s := fmtLogf(ErrorLevel, format, args...)
	fmt.Fprintln(log.writer, s)
	panic(s)
}

func Fatal(args ...interface{}) {
	s := fmtLog(ErrorLevel, args...)
	fmt.Fprintln(log.writer, s)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	s := fmtLogf(ErrorLevel, format, args...)
	fmt.Fprintln(log.writer, s)
	os.Exit(1)
}

func Error(args ...interface{}) {
	if log.logLevel >= ErrorLevel {
		fmt.Fprintln(log.writer, fmtLog(ErrorLevel, args...))
	}
}

func Errorf(format string, args ...interface{}) {
	if log.logLevel >= ErrorLevel {
		fmt.Fprintln(log.writer, fmtLogf(ErrorLevel, format, args...))
	}
}

func Warn(args ...interface{}) {
	if log.logLevel >= WarnLevel {
		fmt.Fprintln(log.writer, fmtLog(WarnLevel, args...))
	}
}

func Warnf(format string, args ...interface{}) {
	if log.logLevel >= WarnLevel {
		fmt.Fprintln(log.writer, fmtLogf(WarnLevel, format, args...))
	}
}

func Debug(args ...interface{}) {
	if log.logLevel >= DebugLevel {
		fmt.Fprintln(log.writer, fmtLog(DebugLevel, args...))
	}
}

func Debugf(format string, args ...interface{}) {
	if log.logLevel >= DebugLevel {
		fmt.Fprintln(log.writer, fmtLogf(DebugLevel, format, args...))
	}
}

func Trace(args ...interface{}) {
	if log.logLevel >= TraceLevel {
		fmt.Fprintln(log.writer, fmtLog(TraceLevel, args...))
	}
}

func Tracef(format string, args ...interface{}) {
	if log.logLevel >= TraceLevel {
		fmt.Fprintln(log.writer, fmtLogf(TraceLevel, format, args...))
	}
}

func fmtLog(level LogLevel, args ...interface{}) string {
	return fmt.Sprintf("%v: [%s] %s%s", time.Now().Format(log.timeFormat), logLevel(level), log.prefix, fmt.Sprint(args...))
}

func fmtLogf(level LogLevel, format string, args ...interface{}) string {
	return fmt.Sprintf("%v: [%s] %s%s", time.Now().Format(log.timeFormat), logLevel(level), log.prefix,
		fmt.Sprintf(format, args...))
}
