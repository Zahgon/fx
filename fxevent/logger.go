package fxevent

type Logger interface {
	LogEvent(Event)
}

var NopLogger = nopLogger{}

type nopLogger struct{}

var _ Logger = nopLogger{}

func (nopLogger) LogEvent(Event) { _ = "STUB: not implemented"; return }

func (nopLogger) String() string { _ = "STUB: not implemented"; return "" }
