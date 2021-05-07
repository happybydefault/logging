package logging

import "strings"

type named struct {
	name   string
	logger Logger
}

func Named(name string, logger Logger) named {
	return named{
		name:   name,
		logger: logger,
	}
}

func (n named) Debug(v ...interface{}) {
	prefix := n.name + ":"
	if len(v) > 0 {
		prefix = prefix + " "
	}
	v = append([]interface{}{prefix}, v...)
	n.logger.Debug(v...)
}
func (n named) Info(v ...interface{}) {
	prefix := n.name + ":"
	if len(v) > 0 {
		prefix = prefix + " "
	}
	v = append([]interface{}{prefix}, v...)
	n.logger.Info(v...)
}
func (n named) Warn(v ...interface{}) {
	prefix := n.name + ":"
	if len(v) > 0 {
		prefix = prefix + " "
	}
	v = append([]interface{}{prefix}, v...)
	n.logger.Warn(v...)
}
func (n named) Error(v ...interface{}) {
	prefix := n.name + ":"
	if len(v) > 0 {
		prefix = prefix + " "
	}
	v = append([]interface{}{prefix}, v...)
	n.logger.Error(v...)
}

func (n named) Debugf(format string, v ...interface{}) {
	format = strings.Join([]string{"%s:", format}, " ")
	v = append([]interface{}{n.name}, v...)
	n.logger.Debugf(format, v...)
}
func (n named) Infof(format string, v ...interface{}) {
	format = strings.Join([]string{"%s:", format}, " ")
	v = append([]interface{}{n.name}, v...)
	n.logger.Infof(format, v...)
}
func (n named) Warnf(format string, v ...interface{}) {
	format = strings.Join([]string{"%s:", format}, " ")
	v = append([]interface{}{n.name}, v...)
	n.logger.Warnf(format, v...)
}
func (n named) Errorf(format string, v ...interface{}) {
	format = strings.Join([]string{"%s:", format}, " ")
	v = append([]interface{}{n.name}, v...)
	n.logger.Errorf(format, v...)
}

func (n named) Debugw(msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{n.name + ":", msg}, " ")
	n.logger.Debugw(msg, keysAndValues...)
}
func (n named) Infow(msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{n.name + ":", msg}, " ")
	n.logger.Infow(msg, keysAndValues...)
}
func (n named) Warnw(msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{n.name + ":", msg}, " ")
	n.logger.Warnw(msg, keysAndValues...)
}
func (n named) Errorw(msg string, keysAndValues ...interface{}) {
	msg = strings.Join([]string{n.name + ":", msg}, " ")
	n.logger.Errorw(msg, keysAndValues...)
}
