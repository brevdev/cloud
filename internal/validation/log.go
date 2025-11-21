package validation

import (
	"context"
	"fmt"
	"strings"

	v1 "github.com/brevdev/cloud/v1"
)

// implementation of the "v1.Logger" interface
type ValidationLogger struct{}

var _ v1.Logger = &ValidationLogger{}

func (l *ValidationLogger) Debug(_ context.Context, msg string, fields ...v1.Field) {
	log("DEBUG", msg, fields...)
}

func (l *ValidationLogger) Info(_ context.Context, msg string, fields ...v1.Field) {
	log("INFO", msg, fields...)
}

func (l *ValidationLogger) Warn(_ context.Context, msg string, fields ...v1.Field) {
	log("WARN", msg, fields...)
}

func (l *ValidationLogger) Error(_ context.Context, err error, fields ...v1.Field) {
	log("ERROR", err.Error(), fields...)
}

func log(level string, msg string, fields ...v1.Field) {
	fmt.Printf("%s: %s\n", level, msg)

	if len(fields) > 0 {
		fieldStrings := []string{}
		for _, field := range fields {
			fieldStrings = append(fieldStrings, fmt.Sprintf("%s: %s", field.Key, field.Value))
		}

		fmt.Printf("{%s}\n", strings.Join(fieldStrings, ", "))
	}
}
