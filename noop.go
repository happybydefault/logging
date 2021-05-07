package logging

// Noop is a no-op implementation of the logger interface. You can use it if you don't want to log at all.
type Noop struct{}

func (n Noop) Debug(...interface{}) {}
func (n Noop) Info(...interface{})  {}
func (n Noop) Warn(...interface{})  {}
func (n Noop) Error(...interface{}) {}

func (n Noop) Debugf(string, ...interface{}) {}
func (n Noop) Infof(string, ...interface{})  {}
func (n Noop) Warnf(string, ...interface{})  {}
func (n Noop) Errorf(string, ...interface{}) {}

func (n Noop) Debugw(string, ...interface{}) {}
func (n Noop) Infow(string, ...interface{})  {}
func (n Noop) Warnw(string, ...interface{})  {}
func (n Noop) Errorw(string, ...interface{}) {}
