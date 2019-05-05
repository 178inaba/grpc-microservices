package main

import (
	"context"
	"fmt"

	pbproject "github.com/178inaba/grpc-microservices/proto/project"
	pb "github.com/178inaba/grpc-microservices/proto/user"
	"github.com/178inaba/grpc-microservices/shared/metadata"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultProjectName = "Default"

type userService struct {
	store         store
	projectClient pbproject.ProjectServiceClient
}

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.Email == "" || len(req.Password) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "empty email or password")
	}

	passwordHash, err := bcrypt.GenerateFromPassword(req.Password, bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("generate from password: %v", err))
	}

	user, err := s.store.createUser(&pb.User{
		Email:        req.Email,
		PasswordHash: passwordHash,
		CreatedAt:    ptypes.TimestampNow(),
	})
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("create user: %v", err))
	}

	ctx = metadata.AddUserIDToContext(ctx, user.GetId())
	if _, err := s.projectClient.CreateProject(ctx, &pbproject.CreateProjectRequest{
		Name: defaultProjectName,
	}); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{User: user}, nil
}

func (s *userService) FindUser(ctx context.Context, req *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	user, err := s.store.findUser(req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("find user: %v", err))
	}

	return &pb.FindUserResponse{User: user}, nil
}

func (s *userService) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyUserResponse, error) {
	user, err := s.store.findUserByEmail(req.GetEmail())
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("find user by email: %v", err))
	}

	if err := bcrypt.CompareHashAndPassword(user.GetPasswordHash(), req.GetPassword()); err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("compare hash and password: %v", err))
	}

	return &pb.VerifyUserResponse{User: user}, nil
}
