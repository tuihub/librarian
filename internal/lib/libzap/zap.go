package libzap

import (
	"io"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics.
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

type RotateOptions struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type LevelEnablerFunc func(level Level) bool

type TeeOption struct {
	Filename         string
	Rotate           RotateOptions
	LevelEnablerFunc LevelEnablerFunc
}

func NewTeeWithRotate(teeOptions []TeeOption, zapOptions ...zap.Option) *zap.Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.MessageKey = ""
	cfg.EncoderConfig.TimeKey = ""

	for _, opt := range teeOptions {
		opt := opt

		lv := zap.LevelEnablerFunc(func(level Level) bool {
			return opt.LevelEnablerFunc(level)
		})

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   opt.Filename,
			MaxSize:    opt.Rotate.MaxSize,
			MaxAge:     opt.Rotate.MaxAge,
			MaxBackups: opt.Rotate.MaxBackups,
			LocalTime:  false,
			Compress:   opt.Rotate.Compress,
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(w),
			lv,
		)
		cores = append(cores, core)
	}

	return zap.New(zapcore.NewTee(cores...), zapOptions...)
}

// New create a new logs (not support log rotating).
func New(writer io.Writer, level Level, opts ...zap.Option) *zap.Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		level,
	)

	return zap.New(core, opts...)
}

func NewDefaultLogger(basePath string) *zap.Logger {
	var tops = []TeeOption{
		{
			Filename: path.Join(basePath, "log", "access.log"),
			Rotate: RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3, //nolint:gomnd // default settings
				Compress:   true,
			},
			LevelEnablerFunc: func(lvl Level) bool {
				return lvl > DebugLevel
			},
		},
		{
			Filename: path.Join(basePath, "log", "error.log"),
			Rotate: RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3, //nolint:gomnd // default settings
				Compress:   true,
			},
			LevelEnablerFunc: func(lvl Level) bool {
				return lvl > InfoLevel
			},
		},
	}

	return NewTeeWithRotate(tops)
}
