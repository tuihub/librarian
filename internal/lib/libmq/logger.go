package libmq

import (
	"fmt"

	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/go-kratos/kratos/v2/log"
)

const defaultCallerValue = "watermill"

type mqLogger struct {
	fields watermill.LogFields
}

func newMQLogger() *mqLogger {
	return new(mqLogger)
}

func (l *mqLogger) Error(msg string, err error, fields watermill.LogFields) {
	fields = fields.Add(l.fields)
	fields = fields.Add(watermill.LogFields{logger.DefaultCallerKey: defaultCallerValue})
	fields = fields.Add(watermill.LogFields{log.DefaultMessageKey: fmt.Sprintf("%s err: %s", msg, err)})
	log.Log(log.LevelError, l.toKeyValues(fields)...)
}
func (l *mqLogger) Info(msg string, fields watermill.LogFields) {
	fields = fields.Add(l.fields)
	fields = fields.Add(watermill.LogFields{logger.DefaultCallerKey: defaultCallerValue})
	fields = fields.Add(watermill.LogFields{log.DefaultMessageKey: msg})
	log.Log(log.LevelInfo, l.toKeyValues(fields)...)
}
func (l *mqLogger) Debug(msg string, fields watermill.LogFields) {
	fields = fields.Add(l.fields)
	fields = fields.Add(watermill.LogFields{logger.DefaultCallerKey: defaultCallerValue})
	fields = fields.Add(watermill.LogFields{log.DefaultMessageKey: msg})
	log.Log(log.LevelDebug, l.toKeyValues(fields)...)
}
func (l *mqLogger) Trace(msg string, fields watermill.LogFields) {
	fields = fields.Add(l.fields)
	fields = fields.Add(watermill.LogFields{logger.DefaultCallerKey: defaultCallerValue})
	fields = fields.Add(watermill.LogFields{log.DefaultMessageKey: msg})
	log.Log(log.LevelDebug, l.toKeyValues(fields)...)
}
func (l *mqLogger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &mqLogger{fields: l.fields.Add(fields)}
}

func (l *mqLogger) toKeyValues(fields watermill.LogFields) []interface{} {
	res := make([]interface{}, len(fields)*2) //nolint:gomnd //double size is correct
	i := 0
	for k, v := range fields {
		res[i] = k
		res[i+1] = v
		i += 2
	}
	return res
}
