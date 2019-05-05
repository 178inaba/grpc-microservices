package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/178inaba/grpc-microservices/front/handler"
	"github.com/178inaba/grpc-microservices/front/interceptor"
	"github.com/178inaba/grpc-microservices/front/middleware"
	"github.com/178inaba/grpc-microservices/front/session"
	pbactivity "github.com/178inaba/grpc-microservices/proto/activity"
	pbproject "github.com/178inaba/grpc-microservices/proto/project"
	pbtask "github.com/178inaba/grpc-microservices/proto/task"
	pbuser "github.com/178inaba/grpc-microservices/proto/user"
	"github.com/gorilla/mux"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

const port = 8080

func main() {
	ctx := context.Background()

	activityConn, err := getGRPCConn(ctx, os.Getenv("ACTIVITY_SERVICE_ADDR"), interceptor.XTraceID, interceptor.XUserID)
	if err != nil {
		log.Fatalf("Failed to get connection of activity: %v.", err)
	}
	activityClient := pbactivity.NewActivityServiceClient(activityConn)

	projectConn, err := getGRPCConn(ctx, os.Getenv("PROJECT_SERVICE_ADDR"), interceptor.XTraceID, interceptor.XUserID)
	if err != nil {
		log.Fatalf("Failed to get connection of project: %v.", err)
	}
	projectClient := pbproject.NewProjectServiceClient(projectConn)

	taskConn, err := getGRPCConn(ctx, os.Getenv("TASK_SERVICE_ADDR"), interceptor.XTraceID, interceptor.XUserID)
	if err != nil {
		log.Fatalf("Failed to get connection of task: %v.", err)
	}
	taskClient := pbtask.NewTaskServiceClient(taskConn)

	userConn, err := getGRPCConn(ctx, os.Getenv("USER_SERVICE_ADDR"), interceptor.XTraceID)
	if err != nil {
		log.Fatalf("Failed to get connection of user: %v.", err)
	}
	userClient := pbuser.NewUserServiceClient(userConn)

	sessionStore := session.NewStoreOnMemory()
	frontSrv := &handler.FrontServer{
		ActivityClient: activityClient,
		ProjectClient:  projectClient,
		TaskClient:     taskClient,
		UserClient:     userClient,
		SessionStore:   sessionStore,
	}

	r := mux.NewRouter()

	// Add handler common processing using middleware.
	r.Use(middleware.Tracing)
	r.Use(middleware.Logging)
	auth := middleware.NewAuthentication(userClient, sessionStore)

	// Mapping of endpoints and methods.
	// (Add middleware for certification check to endpoints that require certification.)
	// TODO Implements method.
	r.Path("/").Methods(http.MethodGet).HandlerFunc(auth(frontSrv.ViewHome))
	r.Path("/logout").Methods(http.MethodPost).HandlerFunc(auth(frontSrv.Logout))
	r.Path("/project").Methods(http.MethodPost).HandlerFunc(auth(frontSrv.CreateProject))
	r.Path("/project/{id}").Methods(http.MethodGet).HandlerFunc(auth(frontSrv.ViewProject))
	r.Path("/project/{id}").Methods(http.MethodPost).HandlerFunc(auth(frontSrv.UpdateProject))
	r.Path("/task").Methods(http.MethodPost).HandlerFunc(auth(frontSrv.CreateTask))
	r.Path("/task/{id}").Methods(http.MethodPost).HandlerFunc(auth(frontSrv.UpdateTask))
	r.Path("/signup").Methods(http.MethodGet).HandlerFunc(frontSrv.ViewSignup)
	r.Path("/signup").Methods(http.MethodPost).HandlerFunc(frontSrv.Signup)
	//r.Path("/login").Methods(http.MethodGet).HandlerFunc(frontSrv.ViewLogin)
	//r.Path("/login").Methods(http.MethodPost).HandlerFunc(frontSrv.Login)

	static := http.StripPrefix("/static", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(static)

	log.Printf("Start server on port: %d.", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Printf("Exit server: %v.", err)
	}
}

func getGRPCConn(ctx context.Context, target string, interceptors ...grpc.UnaryClientInterceptor) (*grpc.ClientConn, error) {
	chain := grpc_middleware.ChainUnaryClient(interceptors...)
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure(), grpc.WithUnaryInterceptor(chain))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
