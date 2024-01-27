package xlog

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type Formatter struct {
	formatter logrus.JSONFormatter
}

func NewFormatter() *Formatter {
	return &Formatter{
		formatter: logrus.JSONFormatter{},
	}
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	span := trace.SpanFromContext(entry.Context)
	entry.Data["trace_id"] = span.SpanContext().TraceID().String()
	entry.Data["span_id"] = span.SpanContext().SpanID().String()
	// entry.Data["context"] = span.SpanContext()
	return f.formatter.Format(entry)
}
