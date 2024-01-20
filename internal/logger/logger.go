package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	debug  bool
	logger *slog.Logger
}

func New(debug bool) *Logger {
	file, err := os.OpenFile("log.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))
	// logger.SetOptions(slf4go.WithTimestampFormat("2006-01-02 15:04:05"))

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	if l.debug {
		l.logger.Debug(msg, args...)
	}
}
