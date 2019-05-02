package main

import (
	pbactivity "github.com/178inaba/grpc-microservices/proto/activity"
	pbproject "github.com/178inaba/grpc-microservices/proto/project"
)

type TaskService struct {
	store          Store
	activityClient pbactivity.ActivityServiceClient
	projectClient  pbproject.ProjectServiceClient
}
