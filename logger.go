package log

import (
	"fmt"
	deflog "log"
	"os"
	"time"
)

type Format int

const (
	FormatText = iota
	FormatJSON
)

type Logger struct {
	format    Format
	threshold Level
	location  *time.Location
	logger    *deflog.Logger
}

func NewLogger() *Logger {
	logger := &Logger{
		format:    FormatText,
		threshold: LevelDefault,
		location:  time.UTC,
		logger:    deflog.New(os.Stdout, "", 0),
	}
	return logger
}

func (logger *Logger) SetLogLevel(lv Level) error {
	if err := lv.validate(); err != nil {
		return err
	}
	logger.threshold = lv
	return nil
}

func (logger *Logger) SetLogFormat(format Format) error {
	logger.format = format
	return nil
}

func (logger *Logger) output(lv Level, message string, a ...interface{}) error {
	if !lv.higher(logger.threshold) {
		return nil
	}
	if err := lv.validate(); err != nil {
		return err
	}

	event := newEvent(time.Now().In(logger.location), lv, fmt.Sprintf(message, a...))

	var str string
	switch logger.format {
	case FormatText:
		str = event.text()
	case FormatJSON:
		str = event.json()
	}
	logger.logger.Println(str)
	return nil
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	logger.output(LevelDebug, format, a...)
}

func (logger *Logger) Info(format string, a ...interface{}) {
	logger.output(LevelInfo, format, a...)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
	logger.output(LevelWarning, format, a...)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	logger.output(LevelError, format, a...)
}
