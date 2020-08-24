package log

import (
	"go.uber.org/zap"
)

var (
	sugaredLogger *zap.SugaredLogger
	logger        *zap.Logger
	cfg           zap.Config
	IsDebug       = true
)

func init() {
	InitLogger(true)
}

func InitLogger(isDebug bool) *zap.SugaredLogger {
	cfg = zap.NewProductionConfig()
	IsDebug = isDebug
	if isDebug {
		cfg = zap.NewDevelopmentConfig()
	}
	logger, _ = cfg.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.DPanicLevel))
	defer func() { _ = logger.Sync() }()
	sugaredLogger = logger.Named("main").Sugar()
	return sugaredLogger
}

func Debug(args ...interface{}) { sugaredLogger.Debug(args...) }

func Debugf(template string, args ...interface{}) { sugaredLogger.Debugf(template, args...) }

func Debugw(msg string, keysAndValues ...interface{}) { sugaredLogger.Debugw(msg, keysAndValues...) }

func Info(args ...interface{}) { sugaredLogger.Info(args...) }

func Infof(template string, args ...interface{}) { sugaredLogger.Infof(template, args...) }

func InfoW(msg string, keysAndValues ...interface{}) { sugaredLogger.Infow(msg, keysAndValues...) }

func Warn(args ...interface{}) { sugaredLogger.Warn(args...) }

func Warnf(template string, args ...interface{}) { sugaredLogger.Warnf(template, args...) }

func Warnw(msg string, keysAndValues ...interface{}) { sugaredLogger.Warnw(msg, keysAndValues...) }

func Error(args ...interface{}) { sugaredLogger.Error(args...) }

func Errorf(template string, args ...interface{}) { sugaredLogger.Errorf(template, args...) }

func Errorw(msg string, keysAndValues ...interface{}) { sugaredLogger.Errorw(msg, keysAndValues...) }

func DPanic(args ...interface{}) { sugaredLogger.DPanic(args...) }

func DPanicf(template string, args ...interface{}) { sugaredLogger.DPanicf(template, args...) }

func DPanicw(msg string, keysAndValues ...interface{}) { sugaredLogger.DPanicw(msg, keysAndValues...) }

func Panic(args ...interface{}) { sugaredLogger.Panic(args...) }

func Panicf(template string, args ...interface{}) { sugaredLogger.Panicf(template, args...) }

func Panicw(msg string, keysAndValues ...interface{}) { sugaredLogger.Panicw(msg, keysAndValues...) }

func Fatal(args ...interface{}) { sugaredLogger.Fatal(args...) }

func Fatalf(template string, args ...interface{}) { sugaredLogger.Fatalf(template, args...) }

func Fatalw(msg string, keysAndValues ...interface{}) { sugaredLogger.Fatalw(msg, keysAndValues...) }

//flush
func Sync() {
	logger.Sync()
}

func CheckFatal(err error) {
	if err != nil {
		sugaredLogger.Fatalf("error: %s", err)
	}
}
func CheckPanic(err error) {
	if err != nil {
		sugaredLogger.Panicf("error: %s", err)
	}
}
