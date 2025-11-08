package logger

type Ilogger interface {
	Debug(msg string, args ...any)
	DebugMsgf(msg string, args ...interface{})
	Info(msg string, args ...any)
	InfoMsgf(msg string, args ...interface{})
	Warn(msg string, args ...any)
	WarnMsgf(msg string, args ...interface{})
	Error(msg string, args ...any) error
	ErrorMsgf(msg string, args ...interface{}) error
}
