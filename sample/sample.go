package main

import (
	"github.com/masa213f/log"
)

func main() {
	level := "default(info)"
	log.Debug("debug log (LogLevel = %s)", level)
	log.Info("info log (LogLevel = %s)", level)
	log.Warn("warn log (LogLevel = %s)", level)
	log.Error("error log (LogLevel = %s)", level)

	log.SetLogLevel(log.LevelDebug)
	level = "debug"
	log.Debug("debug log (LogLevel = %s)", level)
	log.Info("info log (LogLevel = %s)", level)
	log.Warn("warn log (LogLevel = %s)", level)
	log.Error("error log (LogLevel = %s)", level)

	log.SetLogLevel(log.LevelInfo)
	level = "info"
	log.Debug("debug log (LogLevel = %s)", level)
	log.Info("info log (LogLevel = %s)", level)
	log.Warn("warn log (LogLevel = %s)", level)
	log.Error("error log (LogLevel = %s)", level)

	log.SetLogLevel(log.LevelWarn)
	level = "warn"
	log.Debug("debug log (LogLevel = %s)", level)
	log.Info("info log (LogLevel = %s)", level)
	log.Warn("warn log (LogLevel = %s)", level)
	log.Error("error log (LogLevel = %s)", level)

	log.SetLogLevel(log.LevelError)
	level = "error"
	log.Debug("debug log (LogLevel = %s)", level)
	log.Info("info log (LogLevel = %s)", level)
	log.Warn("warn log (LogLevel = %s)", level)
	log.Error("error log (LogLevel = %s)", level)
}
