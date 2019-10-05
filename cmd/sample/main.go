package main

import (
	"fmt"

	"github.com/masa213f/log"
)

func outputLogs(lv log.Level) {
	fmt.Printf("=== Set log level: %s ===\n", lv.String())
	log.SetLogLevel(lv)

	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")
}

func main() {
	outputLogs(log.LevelInfo)
	outputLogs(log.LevelDebug)
	outputLogs(log.LevelError)
	outputLogs(log.LevelWarning)
	outputLogs(log.LevelDefault)

	log.SetLogFormat(log.FormatJSON)
	outputLogs(log.LevelInfo)
	outputLogs(log.LevelDebug)
	outputLogs(log.LevelError)
	outputLogs(log.LevelWarning)
	outputLogs(log.LevelDefault)
}
