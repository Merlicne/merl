package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(logPath string) (closeFunc func(),err error) {

	if logPath == "" {
		return nil, errors.New("ERROR : logPath is required")
	}

	now := time.Now()
	logdir := filepath.Join(logPath, fmt.Sprintf("%s", now.Format("2006-01-02")))
	logfile := filepath.Join(logPath, fmt.Sprintf("%s", now.Format("2006-01-02")), fmt.Sprintf("%s.log", now.Format("15-04-05")))

	err = os.MkdirAll(logdir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.DebugLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	close := func() {
		file.Close()
	}
	
	return close, nil

}




func Info(msg string) {
	if(msg == "") {
		return
	}
	logger.Info(msg)
}

func Error(msg string) {
	if(msg == "") {
		return
	}
	logger.Error(msg)
}

func Debug(msg string) {
	if(msg == "") {
		return
	}
	logger.Debug(msg)
}

func Warn(msg string) {
	if(msg == "") {
		return
	}
	logger.Warn(msg)
}