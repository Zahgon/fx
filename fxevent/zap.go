package fxevent

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	Logger *zap.Logger

	logLevel   zapcore.Level
	errorLevel *zapcore.Level
}

var _ Logger = (*ZapLogger)(nil)

func (l *ZapLogger) UseErrorLevel(level zapcore.Level) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) UseLogLevel(level zapcore.Level) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) logEvent(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) logError(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *ZapLogger) LogEvent(event Event) { _ = "STUB: not implemented"; return }

func moduleField(name string) zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }

func maybeBool(name string, b bool) zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }
