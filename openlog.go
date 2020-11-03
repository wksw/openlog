package openlog

var logger Logger

func SetLogger(l Logger) {
	logger = l
}
func GetLogger() Logger {
	return logger
}

func Debug(message string, opts ...Option) {
	opts = append(opts, WithDepth(2))
	logger.Debug(message, opts...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(message string, opts ...Option) {
	opts = append(opts, WithDepth(2))
	logger.Info(message, opts...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(message string, opts ...Option) {
	opts = append(opts, WithDepth(2))
	logger.Warn(message, opts...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(message string, opts ...Option) {
	opts = append(opts, WithDepth(2))
	logger.Error(message, opts...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(message string, opts ...Option) {
	opts = append(opts, WithDepth(2))
	logger.Fatal(message, opts...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func init() {
	logger = &golog{}
}
