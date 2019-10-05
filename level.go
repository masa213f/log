package log

import (
	"fmt"
	"strings"
)

type Level int

const (
	LevelDefault = Level(0)
	LevelDebug   = Level(1)
	LevelInfo    = Level(2)
	LevelWarning = Level(3)
	LevelError   = Level(4)
)

func StringToLevel(s string) Level {
	switch strings.ToLower(s) {
	case "d", "debug":
		return LevelDebug
	case "i", "info":
		return LevelInfo
	case "w", "warn", "warning":
		return LevelWarning
	case "e", "err", "error":
		return LevelError
	case "", "def", "default":
		return LevelDefault
	default:
		return LevelDefault
	}
}

var levelString = []string{
	"default",
	"debug",
	"info",
	"warning",
	"error",
}

func (lv Level) validate() error {
	if lv < LevelDefault || LevelError < lv {
		return fmt.Errorf("invalid log level: %d", lv)
	}
	return nil
}

func (lv Level) String() string {
	if lv.validate() != nil {
		return ""
	}
	return levelString[lv]
}

func (lv Level) higher(threshold Level) bool {
	if threshold == LevelDefault {
		return lv >= LevelInfo
	}
	return lv >= threshold
}
