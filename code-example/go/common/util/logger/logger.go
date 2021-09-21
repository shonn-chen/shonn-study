package logger

import (
	"log"
	"strings"
	"time"
)

const (
	LEVEL_DEBUG = "DEBUG"
	LEVEL_INFO  = "INFO"
	LEVEL_ERROR = "ERROR"
	LEVEL_WARN  = "WARN"
)

type Option struct {
	Level *string
}

type Logger struct {
	level string
}

var defaultLogger Logger

func init() {
	defaultLogger = Logger{level: "INFO"}
}

func DefaultLogger() *Logger {
	return &defaultLogger
}

func Init(option *Option) {
	if option == nil {
		return
	}
	if option.Level != nil {
		switch strings.ToUpper(*option.Level) {
		case LEVEL_DEBUG:
			defaultLogger.level = LEVEL_DEBUG
		case LEVEL_INFO:
			defaultLogger.level = LEVEL_INFO
		case LEVEL_ERROR:
			defaultLogger.level = LEVEL_ERROR
		default:
		}
	}
}

func (l *Logger) Info(msg string, fields ...*field) {
	switch l.level {
	case LEVEL_INFO, LEVEL_DEBUG:
		l.print("INFO", msg, fields...)
	default:
	}
}

func (l *Logger) Debug(msg string, fields ...*field) {
	switch l.level {
	case LEVEL_DEBUG:
		l.print("DEBUG", msg, fields...)
	default:
	}
}

func (l *Logger) Error(msg string, fields ...*field) {
	switch l.level {
	case LEVEL_INFO, LEVEL_DEBUG, LEVEL_ERROR, LEVEL_WARN:
		l.print("ERROR", msg, fields...)
	default:
	}
}

func (l *Logger) Fatal(msg string, fields ...*field) {
	l.print("FATAL", msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...*field) {
	switch l.level {
	case LEVEL_WARN, LEVEL_INFO, LEVEL_DEBUG:
		l.print("WARN", msg, fields...)
	default:
	}
}

func (l *Logger) convertToArgs(level string, msg string, fields ...*field) []interface{} {
	args := make([]interface{}, 0, len(fields)+3)
	args = append(args, time.Now().String(), level, msg)
	for _, f := range fields {
		args = append(args, f.String())
	}
	return args
}

func (l *Logger) print(level string, msg string, fields ...*field) {
	fmtStr := "%s|%s|msg=%s|" + strings.Repeat("%s, ", len(fields))
	log.Printf(fmtStr, l.convertToArgs(level, msg, fields...)...)
}

func Info(msg string, fields ...*field) {
	defaultLogger.Info(msg, fields...)
}

func Debug(msg string, fields ...*field) {
	defaultLogger.Debug(msg, fields...)
}

func Error(msg string, fields ...*field) {
	defaultLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...*field) {
	defaultLogger.Fatal(msg, fields...)
}

func Warn(msg string, fields ...*field) {
	defaultLogger.Warn(msg, fields...)
}
