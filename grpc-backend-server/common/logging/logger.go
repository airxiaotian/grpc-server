package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"sync"
)

var (
	logLevel    zapcore.Level
	appVersion  string
	appRevision string
	localLogger *zap.SugaredLogger
	once        sync.Once
)

const (
	versionKey  = "version"
	revisionKey = "revision"
)

type Config struct {
	LogLevel    string
	AppVersion  string
	AppRevision string
}

func Configure(config Config) {
	once.Do(func() {
		switch strings.ToUpper(config.LogLevel) {
		case "DEBUG":
			logLevel = zapcore.DebugLevel
		case "INFO":
			logLevel = zapcore.InfoLevel
		case "WARN":
			logLevel = zapcore.WarnLevel
		case "ERROR":
			logLevel = zapcore.ErrorLevel
		case "FATAL":
			logLevel = zapcore.FatalLevel
		default:
			logLevel = zapcore.InfoLevel
		}

		if len(config.AppVersion) > 0 {
			appVersion = config.AppVersion
		}
		if len(config.AppRevision) > 0 {
			appRevision = config.AppRevision
		}

		zapConfig := defaultZapConfig()
		logger, _ := zapConfig.Build() // TODO handle error
		fields := zap.Fields(defaultZapFields()...)
		localLogger = logger.WithOptions(fields).Sugar()
	})
}

func NewLogger() *zap.SugaredLogger {
	if localLogger == nil {
		defaultConfig := Config{
			LogLevel:    "INFO",
			AppVersion:  "NotSetting",
			AppRevision: "NotSetting",
		}
		Configure(defaultConfig)
	}
	return localLogger
}

func With(logger *zap.SugaredLogger, key, value string) *zap.SugaredLogger {
	desugared := logger.Desugar()
	desugared = desugared.WithOptions(zap.Fields(zap.String(key, value)))
	return desugared.Sugar()
}

func defaultZapConfig() zap.Config {
	return zap.Config{
		Level:    zap.NewAtomicLevelAt(logLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func defaultZapFields() []zap.Field {
	return []zap.Field{
		zap.String(versionKey, appVersion),
		zap.String(revisionKey, appRevision),
	}
}
