package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type LogLevel uint32

const (
	LEVEL_INFO  LogLevel = 0
	LEVEL_DEBUG LogLevel = 1
	LEVEL_WARN  LogLevel = 2
	LEVEL_ERROR LogLevel = 3
	LEVEL_FATAL LogLevel = 4
)

const (
	CAT_INFO  = "[INFO]"
	CAT_DEBUG = "[DEBUG]"
	CAT_WARN  = "[WARN]"
	CAT_ERROR = "[ERROR]"
	CAT_FATAL = "[FATAL]"
)

type Logger struct {
	logger   *log.Logger
	level    LogLevel
	category string
}

func New(level string, category string) *Logger {
	return &Logger{
		level:    conevertLogLevel(level),
		category: category,
		logger:   log.New(os.Stdout, category, 0),
	}
}

func NewWithWriter(writer io.Writer, level string, category string) *Logger {
	return &Logger{
		level:    conevertLogLevel(level),
		category: category,
		logger:   log.New(writer, category, 0),
	}
}

func (l *Logger) log(level LogLevel, message string, args ...interface{}) {
	if l.level <= level {
		ts := time.Now().Format("2006-01-02 15:04:05")
		prefix := fmt.Sprintf("%s %s %s : ", conevertLogLevelCategory(level), ts, l.category)
		l.logger.SetPrefix(prefix)
		l.logger.Printf(message, args...)
	}
}

func (l *Logger) Debug(message string, args ...interface{}) {
	l.log(LEVEL_DEBUG, message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log(LEVEL_INFO, message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(LEVEL_WARN, message, args...)
}

func (l *Logger) Error(message string, args ...interface{}) {
	l.log(LEVEL_ERROR, message, args...)
}
