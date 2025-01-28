package main

import (
	"github.com/Merlicne/merl/env"
	"github.com/Merlicne/merl/logger"
	"os"
)

func main() {
	os.Setenv("TZ", "UTC")
	err := env.NewEnvReader(".env")
	if err != nil {
		panic(err)
	}
	logPath := env.GetStringValue("LOGS.PATH")
	println(logPath)
	closeLog, err := logger.InitLogger(logPath)
	if err != nil {
		panic(err)
	}
	defer closeLog()

	logger.Info("Hello, World!")
	logger.Error(env.GetStringValue("TZ"))
}
