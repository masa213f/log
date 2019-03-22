package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Level int

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelDefault = LevelInfo
)

var levelString = []string{
	"debug",
	"info",
	"warn",
	"error",
}

type event struct {
	Level     string `json:"level"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message,omitempty"`
}

type JsonLogger struct {
	level    Level
	location *time.Location
	logger   *log.Logger
}

func NewJsonLogger() *JsonLogger {
	logger := &JsonLogger{
		level:    LevelDefault,
		location: time.UTC,
		logger:   log.New(os.Stdout, "", 0),
	}
	return logger
}

// デフォルトのロガー
var std = NewJsonLogger()

func (l *JsonLogger) SetLogLevel(lv Level) error {
	if lv < LevelDebug || LevelError < lv {
		return fmt.Errorf("invalid log level: %d", lv)
	}
	l.level = lv
	return nil
}

func SetLogLevel(lv Level) error {
	return std.SetLogLevel(lv)
}

const (
	ISO8601Nano = "2006-01-02T15:04:05.999999"
)

func (l *JsonLogger) now() string {
	return time.Now().In(l.location).Format(ISO8601Nano)
}

func (l *JsonLogger) output(lv Level, format string, a ...interface{}) error {
	if lv < l.level {
		return nil
	}
	if lv < LevelDebug || LevelError < lv {
		return fmt.Errorf("invalid log level: %d", lv)
	}

	event := &event{
		Level:     levelString[lv],
		Timestamp: l.now(),
		Message:   fmt.Sprintf(format, a...),
	}

	byte, _ := json.Marshal(event)
	l.logger.Println(string(byte))
	return nil
}

func (l *JsonLogger) Debug(format string, a ...interface{}) {
	l.output(LevelDebug, format, a...)
}

func (l *JsonLogger) Info(format string, a ...interface{}) {
	l.output(LevelInfo, format, a...)
}

func (l *JsonLogger) Warn(format string, a ...interface{}) {
	l.output(LevelWarn, format, a...)
}

func (l *JsonLogger) Error(format string, a ...interface{}) {
	l.output(LevelError, format, a...)
}

func Debug(format string, a ...interface{}) {
	std.Debug(format, a...)
}

func Info(format string, a ...interface{}) {
	std.Info(format, a...)
}

func Warn(format string, a ...interface{}) {
	std.Warn(format, a...)
}

func Error(format string, a ...interface{}) {
	std.Error(format, a...)
}
