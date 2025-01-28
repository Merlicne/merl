package main

import (
	"os"

	"github.com/Merlicne/merl/env"
	"github.com/Merlicne/merl/logger"
)

func main() {
	err := env.NewEnvReader(".env")
	if err != nil {
		panic(err)
	}

	if os.Getenv("TZ") == "" {
		os.Setenv("TZ", env.GetStringValue("TZ"))
	}


	logPath := env.GetStringValue("LOGS.PATH")
	println(logPath)

	logBuilder := logger.NewLogBuilder()
	closeLog, err := logBuilder.AddFileEncoder(logPath)
	logBuilder.AddConsoleEncoder()
	logBuilder.Build()
	if err != nil {
		panic(err)
	}
	defer closeLog()

	logger.Info("Hello, World!")
	logger.Error(env.GetStringValue("TZ"))
}
