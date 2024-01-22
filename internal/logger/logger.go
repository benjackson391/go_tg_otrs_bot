package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

type Logger struct {
	debug  bool
	logger *slog.Logger
}

func New(debug bool) *Logger {
	filePath := "logs/log.json"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Ошибка при создании директории: %v\n", err)
		}
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Ошибка при создании файла: %v\n", err)
		}
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
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
