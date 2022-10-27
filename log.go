package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

type Options struct {
	Path string
	DirType string
	Level logrus.Level
	MaxAge time.Duration
	MaxSize int64 // 单位：M
	RotationTime time.Duration
}

var defaultOpt = Options{
	Path: "storage/log",
	DirType: "/%Y/%m/%d/%H%M.log",
	Level: DebugLevel,
	MaxSize: 15,
	MaxAge: 365 * 24 * time.Hour,
	RotationTime: 24 * time.Hour,
}

func dealParams(options *Options) {
	if options.Path == "" {
		options.Path = defaultOpt.Path
	}
	if options.DirType == "" {
		options.DirType = defaultOpt.DirType
	}
	if options.Level == 0 {
		options.Level = defaultOpt.Level
	}
	if options.MaxSize == 0 {
		options.MaxSize = defaultOpt.MaxSize
	}
	if options.MaxAge == 0 {
		options.MaxAge = defaultOpt.MaxAge
	}
	if options.RotationTime == 0 {
		options.RotationTime = defaultOpt.RotationTime
	}
	return
}

func Init(options Options) {
	dealParams(&options)
	writer, err := rotatelogs.New(options.Path + options.DirType,
		rotatelogs.WithLinkName(options.Path),
		rotatelogs.WithMaxAge(options.MaxAge),
		rotatelogs.WithRotationSize(options.MaxSize),
		)
	if err != nil {
		panic("log init error:" + err.Error())
	}
	logrus.SetOutput(writer)
	logrus.SetLevel(options.Level)
	logrus.SetReportCaller(true)
	fmt.Println("log init success")
}

func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func SetJsonFormatter()  {
	SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimeFormat,
	})
}

func SetTextFormatter()  {
	SetFormatter(&logrus.TextFormatter{
		TimestampFormat: TimeFormat,
	})
}

func SetReportCaller(open bool) {
	logrus.SetReportCaller(open)
}

func SetFormatter(formatter logrus.Formatter) {
	logrus.SetFormatter(formatter)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logrus.WithFields(fields)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func HandleError(module string, err error) {
	logrus.WithFields(logrus.Fields{
		"Module": module,
	}).Error(err)
}