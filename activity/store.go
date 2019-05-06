package main

import (
	"errors"
	"sort"

	pb "github.com/178inaba/grpc-microservices/proto/activity"
	"github.com/178inaba/grpc-microservices/shared/inmemory"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

type store interface {
	createActivity(activity *pb.Activity) (*pb.Activity, error)
	findActivities(userID uint64) ([]*pb.Activity, error)
}

type storeOnMemory struct {
	activities *inmemory.IndexMap
}

func newStoreOnMemory() *storeOnMemory {
	return &storeOnMemory{inmemory.NewIndexMap()}
}

func (s *storeOnMemory) createActivity(activity *pb.Activity) (*pb.Activity, error) {
	if kindOf(activity.GetContent()) == kindUnknown {
		return nil, errors.New("unknown activity content")
	}

	newActivity := *activity
	idx := s.activities.Index()
	newActivity.Id = idx
	s.activities.Set(idx, &newActivity)

	return &newActivity, nil
}

func (s *storeOnMemory) findActivities(userID uint64) ([]*pb.Activity, error) {
	var activities []*pb.Activity
	s.activities.Range(func(idx uint64, value interface{}) bool {
		activity := value.(*pb.Activity)
		if activity.GetUserId() == userID {
			activities = append(activities, activity)
		}

		return true
	})

	sort.Slice(activities, func(i, j int) bool {
		return activities[i].GetCreatedAt().Seconds > activities[j].GetCreatedAt().Seconds
	})

	return activities, nil
}

type kind int32

const (
	kindUnknown kind = iota
	kindCreateTask
	kindUpdateTaskStatus
	kindCreateProject
)

func kindOf(any *any.Any) kind {
	if ptypes.Is(any, new(pb.CreateTaskContent)) {
		return kindCreateTask
	} else if ptypes.Is(any, new(pb.UpdateTaskStatusContent)) {
		return kindUpdateTaskStatus
	} else if ptypes.Is(any, new(pb.CreateProjectContent)) {
		return kindCreateProject
	}

	return kindUnknown
}
