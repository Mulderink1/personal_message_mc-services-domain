package domain

import (
	"go.uber.org/zap"
)

type Message struct {
	Date   string `json:"date"`
	Action string `json:"action"`
	Email  string `json:"email"`
}

type Logger struct {
	Logger *zap.Logger
}

func (l Logger) Info(message string, fields ...zap.Field) {
	l.Logger.Info(message, fields...)
}

func (l Logger) Debug(message string, fields ...zap.Field) {
	l.Logger.Debug(message, fields...)
}

func (l Logger) Error(message string, fields ...zap.Field) {
	l.Logger.Error(message, fields...)
}
