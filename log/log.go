package log

type Logger interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)

	WithError(err error) Logger
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}

// NOTE: Use RootLogger sparingly. Prefer using a derived logger based upon
// the RootLogger created at application start. If you aren't sure, don't use
// it and ask!

func RootLogger() Logger {
	return _logger
}