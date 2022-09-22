package log

import (
	l "github.com/sirupsen/logrus"
)

var (
	// 自定义全局log
	global = l.New()
)

func init() {
	// 全局log属性设置
	l.SetFormatter(&l.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	l.SetLevel(l.DebugLevel)
	l.SetReportCaller(true)

	// 自定义全局logs属性设置
	global.SetFormatter(&l.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	global.SetLevel(l.DebugLevel)
	global.SetReportCaller(true)
}

func SetLevel(v l.Level) {
	l.SetLevel(v)
	global.SetLevel(v)
}

func Debug(i ...interface{}) {
	global.Debug(i)
}

func Debugf(format string, values ...interface{}) {
	global.Debugf(format, values...)
}

func Info(i ...interface{}) {
	global.Info(i)
}

func Infof(format string, values ...interface{}) {
	global.Infof(format, values...)
}

func Warn(i ...interface{}) {
	global.Warn(i)
}

func Warnf(format string, values ...interface{}) {
	global.Warnf(format, values...)
}

func Error(i ...interface{}) {
	global.Error(i)
}

func Errorf(format string, values ...interface{}) {
	global.Errorf(format, values...)
}

func Fatal(i ...interface{}) {
	global.Fatal(i)
}

func Fatalf(format string, values ...interface{}) {
	global.Fatalf(format, values...)
}

func Panic(i ...interface{}) {
	global.Panic(i)
}

func Panicf(format string, args ...interface{}) {
	global.Panicf(format, args)
}
