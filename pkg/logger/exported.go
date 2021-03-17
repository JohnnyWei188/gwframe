package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	stdLogger *zap.SugaredLogger
	once      sync.Once
)

// Instance logger 实例
func Instance() *zap.SugaredLogger {
	if stdLogger == nil {
		panic("logger is nil, you should Init it in main")
	}
	return stdLogger
}

// New logger 实例
func New(opts ...Option) *zap.SugaredLogger {
	once.Do(func() {
		var log *Logger
		log = newLogger(opts...)
		stdLogger = log.SugaredLogger
	})
	return stdLogger
}
