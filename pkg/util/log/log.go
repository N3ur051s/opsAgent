package log

import (
	"io"
	"os"

	l "github.com/sirupsen/logrus"

	. "opsAgent/conf"
)

var (
	global = l.New()
)

func init() {
	l.SetFormatter(&l.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	l.SetLevel(GetLogLvl())
	global.SetFormatter(&l.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	file, err := os.OpenFile("/var/log/opsAgent/opsAgent.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		l.SetOutput(fileAndStdoutWriter)
		global.SetOutput(fileAndStdoutWriter)
	} else {
		global.Errorf("failed to log to file: [%s]", err)
	}
	global.SetLevel(GetLogLvl())
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
