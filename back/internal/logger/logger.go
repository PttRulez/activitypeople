package logger

import "go.uber.org/zap"

var l *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()

	l = logger.Sugar()
}

func Info(args ...interface{}) {
	l.Info(args...)
}

func Debug(args ...interface{}) {
	l.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	l.Debugf(template, args...)
}

func Error(args ...interface{}) {
	l.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	l.Errorf(template, args...)
}
