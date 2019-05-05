package main

import (
	"context"
	"fmt"

	pbactivity "github.com/178inaba/grpc-microservices/proto/activity"
	pb "github.com/178inaba/grpc-microservices/proto/project"
	"github.com/178inaba/grpc-microservices/shared/metadata"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type projectService struct {
	store          store
	activityClient pbactivity.ActivityServiceClient
}

func (s *projectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "empty project name")
	}

	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("get user id from context: %v", err))
	}

	project, err := s.store.createProject(&pb.Project{
		Name:      req.GetName(),
		UserId:    userID,
		CreatedAt: ptypes.TimestampNow(),
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("create project: %v", err))
	}

	any, err := ptypes.MarshalAny(&pbactivity.CreateProjectContent{
		ProjectId:   project.GetId(),
		ProjectName: project.GetName(),
	})
	if err != nil {
		return nil, errors.Errorf("marshal any: %v", err)
	}

	if _, err := s.activityClient.CreateActivity(ctx, &pbactivity.CreateActivityRequest{
		Content: any,
	}); err != nil {
		return nil, errors.Errorf("create activity: %v", err)
	}

	return &pb.CreateProjectResponse{Project: project}, nil
}

func (s *projectService) FindProject(ctx context.Context, req *pb.FindProjectRequest) (*pb.FindProjectResponse, error) {
	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("get user id from context: %v", err))
	}

	project, err := s.store.findProject(req.GetProjectId(), userID)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("find project: %v", err))
	}

	return &pb.FindProjectResponse{Project: project}, nil
}

func (s *projectService) FindProjects(ctx context.Context, _ *empty.Empty) (*pb.FindProjectsResponse, error) {
	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("get user id from context: %v", err))
	}

	projects, err := s.store.findProjects(userID)
	if err != nil {
		return nil, errors.Errorf("find projects: %v", err)
	}

	return &pb.FindProjectsResponse{Projects: projects}, nil
}

func (s *projectService) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error) {
	if req.GetProjectName() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty project name")
	}

	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("get user id from context: %v", err))
	}

	project, err := s.store.findProject(req.GetProjectId(), userID)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("find project: %v", err))
	}

	project.Name = req.GetProjectName()
	if _, err := s.store.updateProject(project); err != nil {
		return nil, errors.Errorf("update project: %v", err)
	}

	return &pb.UpdateProjectResponse{Project: project}, nil
}
