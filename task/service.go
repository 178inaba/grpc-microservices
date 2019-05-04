package main

import (
	"context"
	"fmt"

	pbactivity "github.com/178inaba/grpc-microservices/proto/activity"
	pbproject "github.com/178inaba/grpc-microservices/proto/project"
	pb "github.com/178inaba/grpc-microservices/proto/task"
	"github.com/178inaba/grpc-microservices/shared/metadata"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskService struct {
	store          Store
	activityClient pbactivity.ActivityServiceClient
	projectClient  pbproject.ProjectServiceClient
}

func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	if req.GetName() == "" {
		// Generate error with gRPC status code.
		return nil, status.Error(codes.InvalidArgument, "empty task name")
	}

	// Get a project with the client stub of ProjectService.
	pj, err := s.projectClient.FindProject(
		ctx, &pbproject.FindProjectRequest{ProjectId: req.GetProjectId()})
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("find project: %v", err))
	}

	// Get UserID from metadata.
	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("get user id from context: %v", err))
	}

	// Get the current date and time with Timestamp type in protobuf.
	now := ptypes.TimestampNow()

	// Save the task.
	task, err := s.store.CreateTask(&pb.Task{
		Name:      req.GetName(),
		Status:    pb.Status_WAITING,
		UserId:    userID,
		ProjectId: pj.Project.GetId(),
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("create task: %v", err))
	}

	// Convert the contents of activity to Any type.
	content := &pbactivity.CreateTaskContent{
		TaskId:   task.GetId(),
		TaskName: task.GetName(),
	}
	any, err := ptypes.MarshalAny(content)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("marshal any: %v", err))
	}

	// Create an activity with client stub of ActivityService.
	if _, err := s.activityClient.CreateActivity(
		ctx, &pbactivity.CreateActivityRequest{Content: any}); err != nil {
		return nil, errors.Wrap(err, "create activity")
	}

	return &pb.CreateTaskResponse{Task: task}, nil
}

func (s *TaskService) FindTasks(ctx context.Context, _ *empty.Empty) (*pb.FindTasksResponse, error) {
	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("get user id from context: %v", err))
	}

	tasks, err := s.store.FindTasks(userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("find tasks: %v", err))
	}

	return &pb.FindTasksResponse{Tasks: tasks}, nil
}

func (s *TaskService) FindProjectTasks(ctx context.Context, req *pb.FindProjectTasksRequest) (*pb.FindProjectTasksResponse, error) {
	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("get user id from context: %v", err))
	}

	tasks, err := s.store.FindProjectTasks(req.GetProjectId(), userID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("find project tasks: %v", err))
	}

	return &pb.FindProjectTasksResponse{Tasks: tasks}, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty task name")
	} else if req.GetStatus() == pb.Status_UNKNOWN {
		return nil, status.Error(codes.InvalidArgument, "unknown task status")
	}

	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("get user id from context: %v", err))
	}

	task, err := s.store.FindTask(req.GetTaskId(), userID)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("find task: %v", err))
	}

	updatedTask, err := s.store.UpdateTask(&pb.Task{
		Id:        task.GetId(),
		Name:      req.GetName(),
		Status:    req.GetStatus(),
		ProjectId: task.GetProjectId(),
		UserId:    task.GetUserId(),
		CreatedAt: task.GetCreatedAt(),
		UpdatedAt: ptypes.TimestampNow(),
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("update task: %v", err))
	}

	if task.GetStatus() == updatedTask.GetStatus() {
		return &pb.UpdateTaskResponse{Task: updatedTask}, nil
	}

	any, err := ptypes.MarshalAny(&pbactivity.UpdateTaskStatusContent{
		TaskId:     updatedTask.GetId(),
		TaskName:   updatedTask.GetName(),
		TaskStatus: updatedTask.GetStatus(),
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("marshal any: %v", err))
	}

	if _, err := s.activityClient.CreateActivity(
		ctx, &pbactivity.CreateActivityRequest{Content: any}); err != nil {
		return nil, err
	}

	return &pb.UpdateTaskResponse{Task: updatedTask}, nil
}
