package errors

import (
	"context"

	c "github.com/piyushkumar96/app-error/constants"
)

type TraceMeta struct {
	Trace              []string
	Error              []string
	IdentifierMappings map[string]interface{}
}

func AddTraceLog(ctx context.Context, errorMsg string) *TraceMeta {
	if ctx == nil {
		return nil
	}

	trace := ctx.Value(c.TraceMetaKey)
	traceMeta, ok := trace.(*TraceMeta)
	if !ok {
		return nil
	}

	traceMeta.Error = append(traceMeta.Error, errorMsg)
	return traceMeta
}
