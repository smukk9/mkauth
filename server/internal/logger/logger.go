package logger // Change from 'main' to 'server'

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func LogInfo(msg string, keyVals ...any) {
	logger.Info(msg, keyVals...)
}

func LogError(msg string, keyVals ...any) {
	logger.Error(msg, keyVals...)
}

func LogDebug(msg string, keyVals ...any) {
	logger.Debug(msg, keyVals...)
}
