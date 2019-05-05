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
	var logger *zap.Logger
	var err error
	environment := os.Getenv("GRPC_MICROSERVICES_ENVIRONMENT")
	if environment == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(fmt.Sprintf("zap new: %v, environment: %s.", err, environment))
	}
	defer logger.Sync()

	logger.Info("Start task service.", zap.String("environment", environment))

	ctx := context.Background()

	// Client stub generation.
	activityConn, err := grpc.DialContext(ctx, os.Getenv("ACTIVITY_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Failed to dial activity.", zap.Error(err))
	}

	projectConn, err := grpc.DialContext(ctx, os.Getenv("PROJECT_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Failed to dial project.", zap.Error(err))
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
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT)
	<-quit

	logger.Info("Received a signal of graceful shutdown.")

	stopped := make(chan struct{})
	go func() {
		srv.GracefulStop()
		close(stopped)
	}()

	select {
	case <-time.After(time.Minute):
		srv.Stop()
	case <-stopped:
	}

	logger.Info("Completed graceful shutdown.")
}
