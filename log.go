package fx

import (
	"go.uber.org/fx/fxevent"
)

type logBuffer struct {
	events []fxevent.Event
	logger fxevent.Logger
}

func (l *logBuffer) LogEvent(event fxevent.Event) { _ = "STUB: not implemented"; return }

func (l *logBuffer) Connect(logger fxevent.Logger) { _ = "STUB: not implemented"; return }
