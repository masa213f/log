package log

// デフォルトのロガー
var std = NewLogger()

func SetLogLevel(lv Level) error {
	return std.SetLogLevel(lv)
}

func SetLogFormat(format Format) error {
	return std.SetLogFormat(format)
}

func Debug(format string, a ...interface{}) {
	std.Debug(format, a...)
}

func Info(format string, a ...interface{}) {
	std.Info(format, a...)
}

func Warn(format string, a ...interface{}) {
	std.Warn(format, a...)
}

func Error(format string, a ...interface{}) {
	std.Error(format, a...)
}
