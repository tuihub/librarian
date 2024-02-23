package libcron

import "github.com/tuihub/librarian/internal/lib/logger"

type cronLogger struct{}

func newCronLogger() *cronLogger {
	return new(cronLogger)
}

func (c cronLogger) Debug(msg string, args ...any) {
	logger.Debugf(msg, args...)
}

func (c cronLogger) Error(msg string, args ...any) {
	logger.Errorf(msg, args...)
}

func (c cronLogger) Info(msg string, args ...any) {
	logger.Infof(msg, args...)
}

func (c cronLogger) Warn(msg string, args ...any) {
	logger.Warnf(msg, args...)
}
