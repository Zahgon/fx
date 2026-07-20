package fxevent

import (
	"io"
)

type ConsoleLogger struct {
	W io.Writer
}

var _ Logger = (*ConsoleLogger)(nil)

func (l *ConsoleLogger) logf(msg string, args ...any) { _ = "STUB: not implemented"; return }

func (l *ConsoleLogger) LogEvent(event Event) { _ = "STUB: not implemented"; return }
