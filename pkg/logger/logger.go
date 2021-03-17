package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger algo logger component
type Logger struct {
	*zap.SugaredLogger
	level      string
	filename   string
	maxsize    int
	maxage     int
	maxbackups int
	compress   bool
}

// Option options
type Option func(*Logger)

// New new a logger
func newLogger(opts ...Option) *Logger {
	var zaplogger *zap.Logger
	var level zapcore.Level
	l := &Logger{
		filename:   "", // if filename is empty, logger use os.Stdout output
		level:      "info",
		maxsize:    10,
		maxage:     30,
		maxbackups: 5,
		compress:   false,
	}
	for _, o := range opts {
		o(l)
	}
	//info is the default level
	if err := level.Set(strings.TrimSpace(l.level)); err != nil {
		level = zapcore.InfoLevel
	}
	core := zapcore.NewCore(
		l.getEncoder(),
		l.getWriter(),
		level,
	)
	zaplogger = zap.New(core, zap.AddCaller())
	l.SugaredLogger = zaplogger.Sugar()
	return l
}

// init custom encoder
func (l *Logger) getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// init custom writer
func (l *Logger) getWriter() zapcore.WriteSyncer {
	if l.filename == "" {
		return zapcore.AddSync(os.Stdout)
	}
	lbj := &lumberjack.Logger{
		Filename:   l.filename,
		MaxSize:    l.maxsize,
		MaxBackups: l.maxbackups,
		MaxAge:     l.maxage,
		Compress:   l.compress,
	}
	return zapcore.AddSync(lbj)
}
