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
	log         *zap.SugaredLogger
	once        sync.Once
	closeFileFn func()
	atom        = zap.NewAtomicLevel()
)

func Init(dir, app, level string, reloadCh <-chan struct{}) {
	once.Do(func() {
		core, closeFile := newCore(level, jsonEncoder(), logFileName(dir, app))
		closeFileFn = closeFile
		log = zap.New(core, opts(app)...).Sugar()
		if reloadCh != nil {
			go reloadLogger(dir, app, level, reloadCh)
		}
	})
}
func reloadLogger(dir, app, level string, reloadCh <-chan struct{}) {
	for range reloadCh {
		preLog := log
		core, closeFile := newCore(level, jsonEncoder(), logFileName(dir, app))
		log = zap.New(core, opts(app)...).Sugar()
		_ = preLog.Sync()
		closeFileFn()
		closeFileFn = closeFile
	}
}
func opts(app string) []zap.Option {
	ops := make([]zap.Option, 0, 3)
	ops = append(ops, zap.AddCaller())
	ops = append(ops, zap.AddCallerSkip(1))
	ops = append(ops, zap.Fields(zap.String("app", app)))
	return ops
}
func logFileName(dir, app string) string {
	if strings.HasSuffix(dir, "/") {
		return fmt.Sprintf("%s%s.log", dir, app)
	}
	return fmt.Sprintf("%s/%s.log", dir, app)
}
func newCore(level string, encoder zapcore.Encoder, f string) (zapcore.Core, func()) {
	ws, closeFile, err := zap.Open(f)
	if err != nil {
		panic(err)
	}
	var wss []zapcore.WriteSyncer
	wss = append(wss, ws)
	wss = append(wss, zapcore.AddSync(os.Stdout))
	atom.SetLevel(logLevel(level))
	return zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(wss...),
		atom,
	), closeFile
}

func ChangeRuntimeLevel(level string) {
	atom.SetLevel(logLevel(level))
}

func Debug(args ...any) {
	log.Debug(args...)
}

func Debugf(template string, args ...any) {
	log.Debugf(template, args...)
}

func Info(args ...any) {
	log.Info(args...)
}

func Infof(template string, args ...any) {
	log.Infof(template, args...)
}

func Warn(args ...any) {
	log.Warn(args...)
}

func Warnf(template string, args ...any) {
	log.Warnf(template, args...)
}

func Error(args ...any) {
	log.Error(args...)
}

func Errorf(template string, args ...any) {
	log.Errorf(template, args...)
}

func Panic(args ...any) {
	log.Panic(args...)
}

func Panicf(template string, args ...any) {
	log.Panicf(template, args...)
}
func DPanic(args ...any) {
	log.DPanic(args...)
}
func DPanicf(template string, args ...any) {
	log.DPanicf(template, args...)
}

func Fatal(args ...any) {
	log.Fatal(args...)
}

func Fatalf(template string, args ...any) {
	log.Fatalf(template, args...)
}
