package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/grafodb/grafodb/internal/util/random"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewPreFlightRPCInterceptor(logger *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, next grpc.UnaryHandler) (interface{}, error) {
		meta, _ := metadata.FromIncomingContext(ctx)

		// Add request id
		var requestID string
		if header, ok := meta[string(KeyRequestID)]; ok && len(header) == 1 {
			requestID = header[0]
		}
		if requestID == "" {
			requestID = fmt.Sprintf("%s.%s", random.HexString(10), random.HexString(6))
		}

		// Add request logger
		startTime := time.Now()
		requestLogger := logger.With().Str("reqid", requestID).Logger()
		method := info.FullMethod

		defer func() {
			latency := time.Since(startTime)
			requestLogger.Debug().
				Str("method", method).
				Str("latency", latency.String()).
				Msg("RPC request")
		}()

		ctx = context.WithValue(ctx, KeyRequestID, requestID)
		ctx = context.WithValue(ctx, KeyRequestLogger, &requestLogger)

		return next(ctx, req)
	}
}
