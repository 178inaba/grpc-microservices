package interceptor

import (
	"context"

	"github.com/178inaba/grpc-microservices/front/support"
	"github.com/178inaba/grpc-microservices/shared/metadata"
	"google.golang.org/grpc"
)

func XTraceID(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	traceID := support.GetTraceIDFromContext(ctx)
	ctx = metadata.AddTraceIDToContext(ctx, traceID)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func XUserID(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	user := support.GetUserFromContext(ctx)
	ctx = metadata.AddUserIDToContext(ctx, user.GetId())
	return invoker(ctx, method, req, reply, cc, opts...)
}
