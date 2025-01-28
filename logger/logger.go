package logger

func Info(msg string) {
	if msg == "" {
		return
	}
	logger.Info(msg)
}

func Error(msg string) {
	if msg == "" {
		return
	}
	logger.Error(msg)
}

func Debug(msg string) {
	if msg == "" {
		return
	}
	logger.Debug(msg)
}

func Warn(msg string) {
	if msg == "" {
		return
	}
	logger.Warn(msg)
}
