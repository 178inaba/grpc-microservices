package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pbproject "github.com/178inaba/grpc-microservices/proto/project"
	pb "github.com/178inaba/grpc-microservices/proto/user"
	"github.com/178inaba/grpc-microservices/shared/interceptor"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const port = 50051

func main() {
	// TODO See environment variable.
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	if err != nil {
		panic(fmt.Sprintf("zap new: %v.", err))
	}

	ctx := context.Background()

	projectConn, err := grpc.DialContext(ctx, os.Getenv("PROJECT_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Failed to dial project.", zap.Error(err))
	}
	projectClient := pbproject.NewProjectServiceClient(projectConn)

	srv := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.XTraceID(),
		interceptor.Logging(logger),
	)))

	pb.RegisterUserServiceServer(srv, &userService{
		store:         newStoreOnMemory(),
		projectClient: projectClient,
	})

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
