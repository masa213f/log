package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Format represents log format of a logger.
type Format int

const (
	FormatText = iota
	FormatJSON
)

// A Logger represents a logging object.
type Logger struct {
	mu        sync.Mutex
	format    Format
	threshold Level
	location  *time.Location
	out       io.Writer
}

// NewLogger creates a new Logger.
func NewLogger() *Logger {
	return &Logger{
		format:    FormatText,
		threshold: LevelDefault,
		location:  time.UTC,
		out:       os.Stdout,
	}
}

// SetLogLevel sets the log level.
func (logger *Logger) SetLogLevel(lv Level) error {
	if err := lv.validate(); err != nil {
		return err
	}
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.threshold = lv
	return nil
}

// SetLogFormat sets the log format of the logger.
func (logger *Logger) SetLogFormat(format Format) error {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.format = format
	return nil
}

func (logger *Logger) output(lv Level, message string, a ...interface{}) error {
	if err := lv.validate(); err != nil {
		return err
	}

	logger.mu.Lock()
	defer logger.mu.Unlock()

	if !lv.higher(logger.threshold) {
		return nil
	}

	event := newEvent(time.Now().In(logger.location), lv, fmt.Sprintf(message, a...))

	var str string
	switch logger.format {
	case FormatText:
		str = event.text()
	case FormatJSON:
		str = event.json()
	}
	fmt.Fprintln(logger.out, str)
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
