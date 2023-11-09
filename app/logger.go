package app

import "github.com/withmandala/go-log"

var Logger *log.Logger

func InitLogger(newLogger *log.Logger) {
	Logger = newLogger
}
