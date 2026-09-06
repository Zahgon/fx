//go:build go1.21

package fxevent

import (
	"context"
	"log/slog"
)

var _ Logger = (*SlogLogger)(nil)

type SlogLogger struct {
	Logger *slog.Logger

	ctx        context.Context
	logLevel   slog.Level
	errorLevel *slog.Level
}

func (l *SlogLogger) UseContext(ctx context.Context) { _ = "STUB: not implemented"; return }

func (l *SlogLogger) UseLogLevel(level slog.Level) { _ = "STUB: not implemented"; return }

func (l *SlogLogger) UseErrorLevel(level slog.Level) { _ = "STUB: not implemented"; return }

func (l *SlogLogger) filter(fields []any) []any { _ = "STUB: not implemented"; return nil }

func (l *SlogLogger) logEvent(msg string, fields ...any) { _ = "STUB: not implemented"; return }

func (l *SlogLogger) logError(msg string, fields ...any) { _ = "STUB: not implemented"; return }

func (l *SlogLogger) LogEvent(event Event) { _ = "STUB: not implemented"; return }

type slogFieldSkip struct{}

func slogMaybeModuleField(name string) slog.Attr { _ = "STUB: not implemented"; return *new(slog.Attr) }

func slogMaybeBool(name string, b bool) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}

func slogErr(err error) slog.Attr { _ = "STUB: not implemented"; return *new(slog.Attr) }

func slogStrings(key string, str []string) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}
