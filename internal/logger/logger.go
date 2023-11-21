package logger

import (
	"log"
	"os"
)

var debugLogger = log.New(os.Stdout, "[DEBUG]:\t", log.Ldate|log.Ltime)
var infoLogger = log.New(os.Stdout, "[INFO]:\t", log.Ldate|log.Ltime)
var warningLogger = log.New(os.Stdout, "[WARNING]:\t", log.Ldate|log.Ltime)
var errorLogger = log.New(os.Stdout, "[ERROR]:\t", log.Ldate|log.Ltime)

var debugEnabled = false

func Debug(message string, params ...interface{}) {
	if debugEnabled {
		debugLogger.Printf(message, params...)
	}
}

func Info(message string, params ...interface{}) {
	infoLogger.Printf(message, params...)
}

func Warning(message string, params ...interface{}) {
	warningLogger.Printf(message, params...)
}

func Error(message string, params ...interface{}) {
	errorLogger.Printf(message, params...)
}

func EnableDebug() {
	debugEnabled = true
}
