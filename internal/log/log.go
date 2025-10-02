package log

type Logger struct {
	Errorf func(string, ...any)
	Infof  func(string, ...any)
}

func (l Logger) LogErrorf(format string, args ...any) {
	if l.Errorf == nil {
		return
	}
	l.Errorf(format, args...)
}

func (l Logger) LogInfof(format string, args ...any) {
	if l.Infof == nil {
		return
	}
	l.Infof(format, args...)
}
