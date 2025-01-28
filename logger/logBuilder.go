package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

type logBuilder struct {
	core zapcore.Core
}

func NewLogBuilder() *logBuilder {
	core := zapcore.NewTee()
	return &logBuilder{
		core: core,
	}
}

func (l *logBuilder) AddFileEncoder(logPath string) (closeFunc func(), err error) {
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
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	l.core = zapcore.NewTee(l.core, zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.DebugLevel))
	return func() {
		file.Close()
	}, nil
}

func (l *logBuilder) AddConsoleEncoder() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	l.core = zapcore.NewTee(l.core, zapcore.NewCore(consoleEncoder, zapcore.Lock(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel))
}

func (l *logBuilder) Build() {
	logger = zap.New(l.core)
}
