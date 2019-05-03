package metadata

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

const (
	metadataKeyUserID  = "x-user-id"
	metadataKeyTraceID = "x-trace-id"
)

func AddUserIDToContext(ctx context.Context, userID uint64) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataKeyUserID, strconv.FormatUint(userID, 10))
}

var ErrNotFoundUserID = errors.New("not found user id")

func GetUserIDFromContext(ctx context.Context) (uint64, error) {
	userID, err := SafeGetUserIDFromContext(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "safe get user id from context")
	}

	return userID, nil
}

func SafeGetUserIDFromContext(ctx context.Context) (uint64, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, ErrNotFoundUserID
	}

	values := md.Get(metadataKeyUserID)
	if len(values) < 1 {
		return 0, ErrNotFoundUserID
	}

	userID, err := strconv.ParseUint(values[0], 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "parse uint: %q", values[0])
	}

	return userID, nil
}

func GetTraceIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(metadataKeyTraceID)
	if len(values) < 1 {
		return ""
	}

	return values[0]
}

func AddTraceIDToContext(ctx context.Context, traceID string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataKeyTraceID, traceID)
}
