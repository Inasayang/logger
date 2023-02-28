package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"sync"
)

var (
	log  *zap.SugaredLogger
	once sync.Once
)

func Init(dir, serviceName, level string) {
	once.Do(func() {
		jsonEncoder := jsonEncoder()
		core := zapcore.NewTee(getCores(level, jsonEncoder, dir, serviceName)...)
		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Fields(zap.String("service", serviceName)))
		log = logger.Sugar()
	})
}

func getLogFullPath(dir, serviceName, level string) string {
	if strings.HasSuffix(dir, "/") {
		return fmt.Sprintf("%s%s.%s", dir, serviceName, level)
	}
	return fmt.Sprintf("%s/%s.%s", dir, serviceName, level)
}
func getCores(level string, encoder zapcore.Encoder, d, s string) []zapcore.Core {
	cores := make([]zapcore.Core, 0, 6)
	l := logLevel(level)
	if l <= zapcore.DebugLevel {
		debugLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.DebugLevel
		})
		f := getLogFullPath(d, s, "debug")
		w := zapFileWriter(f)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), debugLevel))
	}
	if l <= zapcore.InfoLevel {
		infoLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.InfoLevel
		})
		f := getLogFullPath(d, s, "info")
		w := zapFileWriter(f)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), infoLevel))
	}
	if l <= zapcore.WarnLevel {
		warnLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.WarnLevel
		})
		f := getLogFullPath(d, s, "warn")
		w := zapFileWriter(f)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), warnLevel))
	}
	if l <= zapcore.ErrorLevel {
		errorLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.ErrorLevel
		})
		f := getLogFullPath(d, s, "error")
		w := zapFileWriter(f)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), errorLevel))
	}
	if l <= zapcore.PanicLevel {
		panicLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.PanicLevel
		})
		f := getLogFullPath(d, s, "panic")
		w := zapFileWriter(f)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), panicLevel))
	}
	if l <= zapcore.FatalLevel {
		fatalLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l >= zapcore.FatalLevel
		})
		f := getLogFullPath(d, s, "fatal")
		w := zapFileWriter(f)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), fatalLevel))
	}
	return cores
}

//func GetLogger() *zap.SugaredLogger {
//	return _log
//}

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

//func DPanic(args ...interface{}) {
//	_log.DPanic(args...)
//}
//
//func DPanicf(template string, args ...interface{}) {
//	_log.DPanicf(template, args...)
//}

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
