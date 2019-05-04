package support

import (
	"context"

	pbuser "github.com/178inaba/grpc-microservices/proto/user"
)

type (
	contextKeyTraceID struct{}
	contextKeyUser    struct{}
)

func GetTraceIDFromContext(ctx context.Context) string {
	id := ctx.Value(contextKeyTraceID{})
	traceID, ok := id.(string)
	if !ok {
		return ""
	}

	return traceID
}

func AddTraceIDToContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, contextKeyTraceID{}, traceID)
}

func GetUserFromContext(ctx context.Context) *pbuser.User {
	u := ctx.Value(contextKeyUser{})
	user, ok := u.(*pbuser.User)
	if !ok {
		return nil
	}

	return user
}

func AddUserToContext(ctx context.Context, user *pbuser.User) context.Context {
	return context.WithValue(ctx, contextKeyUser{}, user)
}
