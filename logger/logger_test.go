package logger_test

import (
	"testing"

	"github.com/tuihub/librarian/logger"
)

func TestLogger(t *testing.T) {
	logger.Debug("test logger")
	logger.Debugf("%s", "test logger")
}
