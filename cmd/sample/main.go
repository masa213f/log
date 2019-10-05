package main

import (
	"fmt"

	"github.com/masa213f/log"
)

func outputLogs() {
	log.Debug("debug message: %s = %d", "hoge", 1)
	log.Info("info message: %s = %d", "piyo", 2)
	log.Warn("warn message: %s = %d", "fuga", 3)
	log.Error("error message: %s = %d", "hoge", 4)
}

func main() {
	fmt.Printf("== Text format(default) ==\n")

	fmt.Printf("= Default(Info) =\n")
	outputLogs()

	fmt.Printf("\n= Debug =\n")
	log.SetLogLevel(log.LevelDebug)
	outputLogs()

	fmt.Printf("\n= Info =\n")
	log.SetLogLevel(log.LevelInfo)
	outputLogs()

	fmt.Printf("\n= Warning =\n")
	log.SetLogLevel(log.LevelWarning)
	outputLogs()

	fmt.Printf("\n= Error =\n")
	log.SetLogLevel(log.LevelError)
	outputLogs()

	fmt.Printf("\n== JSON format ==\n")
	log.SetLogFormat(log.FormatJSON)
	fmt.Printf("= Default(Info) =\n")
	log.SetLogLevel(log.LevelDefault)
	outputLogs()

	fmt.Printf("\n= Debug =\n")
	log.SetLogLevel(log.LevelDebug)
	outputLogs()

	fmt.Printf("\n= Info =\n")
	log.SetLogLevel(log.LevelInfo)
	outputLogs()

	fmt.Printf("\n= Warning =\n")
	log.SetLogLevel(log.LevelWarning)
	outputLogs()

	fmt.Printf("\n= Error =\n")
	log.SetLogLevel(log.LevelError)
	outputLogs()
}
