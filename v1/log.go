package v1

import "context"

type Field struct {
	Key   string
	Value any
}

func LogField(key string, value any) Field {
	return Field{Key: key, Value: value}
}

type Logger interface {
	Debug(ctx context.Context, msg string, fields ...Field)
	Info(ctx context.Context, msg string, fields ...Field)
	Warn(ctx context.Context, msg string, args ...Field)
	Error(ctx context.Context, err error, fields ...Field)
}

type NoopLogger struct{}

func (l *NoopLogger) Debug(_ context.Context, _ string, _ ...Field) {}
func (l *NoopLogger) Info(_ context.Context, _ string, _ ...Field)  {}
func (l *NoopLogger) Warn(_ context.Context, _ string, _ ...Field)  {}
func (l *NoopLogger) Error(_ context.Context, _ error, _ ...Field)  {}

var _ Logger = &NoopLogger{}
