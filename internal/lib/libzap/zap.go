package libzap

import (
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Level = zapcore.Level

const (
	DebugLevel  Level = zap.DebugLevel  // -1
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics.
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
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
		lv := zap.LevelEnablerFunc((func(opt TeeOption) func(level Level) bool {
			return opt.LevelEnablerFunc
		})(opt))

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

func NewStdout(accessLogLevel Level) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		accessLogLevel,
	)
	return zap.New(core)
}

func New(basePath string, accessLogLevel Level) *zap.Logger {
	var tops = []TeeOption{
		{
			Filename: path.Join(basePath, "log", "error.log"),
			Rotate: RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3, //nolint:gomnd // default settings
				Compress:   true,
			},
			LevelEnablerFunc: func(lvl Level) bool {
				return lvl >= ErrorLevel
			},
		},
	}
	if accessLogLevel < ErrorLevel {
		tops = append(tops,
			TeeOption{
				Filename: path.Join(basePath, "log", "access.log"),
				Rotate: RotateOptions{
					MaxSize:    1,
					MaxAge:     1,
					MaxBackups: 3, //nolint:gomnd // default settings
					Compress:   true,
				},
				LevelEnablerFunc: func(lvl Level) bool {
					return lvl >= accessLogLevel
				},
			},
		)
	}

	return NewTeeWithRotate(tops)
}
