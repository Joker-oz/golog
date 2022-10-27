package log

import "testing"

func TestInit(t *testing.T) {
	opt := Options{}
	Init(opt)
}

func TestInfo(t *testing.T) {
	opt := Options{
		Level: ErrorLevel,
	}
	Init(opt)
	SetLevel(InfoLevel)

	Info("hello world", "-- info")
}

func TestSetJsonFormatter(t *testing.T) {
	opt := Options{}
	Init(opt)
	SetJsonFormatter()
	SetReportCaller(true)
	Info("hello world", "-- info")
}

func TestSetReportCaller(t *testing.T) {
	opt := Options{}
	Init(opt)
	SetTextFormatter()
	SetReportCaller(true)
	Info("hello world", "-- info")
}