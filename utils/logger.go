package utils

import (
	"fmt"
	"os"
)

type LogLevel string

const (
	INFO    LogLevel = "INFO"
	WARNING LogLevel = "WARNING"
	ERROR   LogLevel = "ERROR"
)

type Logger struct {
	env string
}

func NewLogger() *Logger {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	return &Logger{env: env}
}

func (l *Logger) Log(level LogLevel, message string, err error) {
	formattedMessage := fmt.Sprintf("[%s] %s: %s", level, message, err)
	if l.env == "local" {
		fmt.Println(formattedMessage)
	} else {
		// this is use on third party - "sentry" ...
	}
}

func (l *Logger) LogError(message string, err error) {
	l.Log(ERROR, message, err)
}

func (l *Logger) LogWarning(message string, err error) {
	l.Log(WARNING, message, err)
}

func (l *Logger) LogInfo(message string) {
	l.Log(INFO, message, nil)
}
