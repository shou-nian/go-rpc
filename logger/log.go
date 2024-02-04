package logger

import (
	"log/slog"
	"time"
)

type msgType interface {
	any
}

type Log interface {
	Info(msg msgType)
	Error(msg msgType)
}

type Logger struct{}

func (l *Logger) Info(msg msgType) {
	infoTime := time.Now()
	slog.Info(infoTime.Local().String(), msg)
}

func (l *Logger) Error(msg msgType) {
	errorTime := time.Now()
	slog.Info(errorTime.Local().String(), msg)
}
