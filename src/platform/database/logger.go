package database

type Logger struct {
	callback func(format string, v ...interface{})
}

func (logger *Logger) SetCallback(callbak func(format string, v ...interface{})) {
	logger.callback = callbak
}

func (logger Logger) Printf(format string, v ...interface{}) {
	logger.callback(format, v)
}
