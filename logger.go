// Package logging defines a very simple logger interface conforming with logging best practices and Clean Architecture.
// It's specially useful as a standardized logger for microservices and libraries.
package logging

// Logger defines a logger with 4 levels: Debug, Info, Warn and Error. All of them have 3 types of functions:
// simple, according to a format and with key-value pairs.
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})

	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
}
