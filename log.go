package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  *zap.SugaredLogger
	once sync.Once
)

func Init(dir, serviceName, level string) {
	once.Do(func() {
		log = zap.New(newCore(level, jsonEncoder(), logFileName(dir, serviceName)), zap.AddCaller(), zap.AddCallerSkip(1), zap.Fields(zap.String("service", serviceName))).Sugar()
	})
}
func logFileName(dir, serviceName string) string {
	if strings.HasSuffix(dir, "/") {
		return fmt.Sprintf("%s%s.log", dir, serviceName)
	}
	return fmt.Sprintf("%s/%s.log", dir, serviceName)
}
func newCore(level string, encoder zapcore.Encoder, f string) zapcore.Core {
	logFile, err := os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile)),
		logLevel(level),
	)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}
