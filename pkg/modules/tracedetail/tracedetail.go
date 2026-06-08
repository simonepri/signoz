package tracedetail

import (
	"context"
	"net/http"

	"github.com/SigNoz/signoz/pkg/types/spantypes"
)

// Handler exposes HTTP handlers for trace detail APIs.
type Handler interface {
	GetWaterfallV4(http.ResponseWriter, *http.Request)
	GetTraceAggregations(http.ResponseWriter, *http.Request)
}

// Module defines the business logic for trace detail operations.
type Module interface {
	GetWaterfallV4(ctx context.Context, traceID string, selectedSpanID string, uncollapsedSpans []string) (*spantypes.GettableWaterfallTrace, error)
	GetTraceAggregations(ctx context.Context, traceID string, req *spantypes.PostableTraceAggregations) (*spantypes.GettableTraceAggregations, error)
}
