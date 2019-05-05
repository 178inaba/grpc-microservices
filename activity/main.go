package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/178inaba/grpc-microservices/proto/activity"
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

	logger.Info("Start activity service.", zap.String("environment", environment))

	srv := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.XTraceID(),
		interceptor.Logging(logger),
		interceptor.XUserID(),
	)))

	pb.RegisterActivityServiceServer(srv, &activityService{
		store: newStoreOnMemory(),
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
