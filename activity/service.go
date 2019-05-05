package main

import (
	"context"
	"fmt"

	pb "github.com/178inaba/grpc-microservices/proto/activity"
	"github.com/178inaba/grpc-microservices/shared/metadata"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type activityService struct {
	store store
}

func (s *activityService) CreateActivity(ctx context.Context, req *pb.CreateActivityRequest) (*empty.Empty, error) {
	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("get user id from context: %v", err))
	}

	if _, err := s.store.createActivity(&pb.Activity{
		Content:   req.GetContent(),
		UserId:    userID,
		CreatedAt: ptypes.TimestampNow(),
	}); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *activityService) FindActivities(ctx context.Context, _ *empty.Empty) (*pb.FindActivitiesResponse, error) {
	userID, err := metadata.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("get user id from context: %v", err))
	}

	activities, err := s.store.findActivities(userID)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("find activities: %v", err))
	}

	return &pb.FindActivitiesResponse{Activities: activities}, nil
}
