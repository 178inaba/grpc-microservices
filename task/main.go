package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pbactivity "github.com/178inaba/grpc-microservices/proto/activity"
	pbproject "github.com/178inaba/grpc-microservices/proto/project"
	pb "github.com/178inaba/grpc-microservices/proto/task"
	"github.com/178inaba/grpc-microservices/shared/interceptor"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const port = 50051

func main() {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	if err != nil {
		panic(fmt.Sprintf("zap new: %v.", err))
	}

	// Client stub generation.
	activityConn, err := grpc.Dial(os.Getenv("ACTIVITY_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Failed to dial activity", zap.Error(err))
	}

	projectConn, err := grpc.Dial(os.Getenv("PROJECT_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Failed to dial project", zap.Error(err))
	}

	// Add Interceptor.
	chain := grpc_middleware.ChainUnaryServer(
		interceptor.XTraceID(),
		interceptor.Logging(logger),
		interceptor.XUserID(),
	)

	srv := grpc.NewServer(grpc.UnaryInterceptor(chain))

	// Service registration.
	pb.RegisterTaskServiceServer(srv, &TaskService{
		store:          NewStoreOnMemory(),
		activityClient: pbactivity.NewActivityServiceClient(activityConn),
		projectClient:  pbproject.NewProjectServiceClient(projectConn),
	})

	// Waiting for gRPC connection.
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			logger.Fatal("Failed to create listener.", zap.Error(err))
		}

		logger.Info("Start server.", zap.Int("port", port))
		if err := srv.Serve(lis); err != nil {
			logger.Info("Exit server.", zap.Error(err))
		}
	}()

	// Graceful stop
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM)
	<-sigint
	logger.Info("TODO")
	stopped := make(chan struct{})
	go func() {
		srv.GracefulStop()
		close(stopped)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	select {
	case <-ctx.Done():
		srv.Stop()
	case <-stopped:
		cancel()
	}

	logger.Info("TODO")
}
