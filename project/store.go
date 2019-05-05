package main

import (
	"sort"

	pb "github.com/178inaba/grpc-microservices/proto/project"
	"github.com/178inaba/grpc-microservices/shared/inmemory"
	"github.com/pkg/errors"
)

type store interface {
	createProject(project *pb.Project) (*pb.Project, error)
	findProject(projectID, userID uint64) (*pb.Project, error)
	findProjects(userID uint64) ([]*pb.Project, error)
	updateProject(project *pb.Project) (*pb.Project, error)
}

type storeOnMemory struct {
	projects *inmemory.IndexMap
}

func newStoreOnMemory() *storeOnMemory {
	return &storeOnMemory{inmemory.NewIndexMap()}
}

func (s *storeOnMemory) createProject(project *pb.Project) (*pb.Project, error) {
	newProject := *project
	idx := s.projects.Index()
	newProject.Id = idx
	s.projects.Set(idx, &newProject)

	return &newProject, nil
}

func (s *storeOnMemory) findProject(projectID, userID uint64) (*pb.Project, error) {
	value, ok := s.projects.Get(projectID)
	if !ok {
		return nil, errors.New("not found project")
	}

	project := value.(*pb.Project)
	if project.GetUserId() != userID {
		return nil, errors.New("not found project")
	}

	return project, nil
}

func (s *storeOnMemory) findProjects(userID uint64) ([]*pb.Project, error) {
	var projects []*pb.Project
	s.projects.Range(func(idx uint64, value interface{}) bool {
		project := value.(*pb.Project)
		if project.GetUserId() == userID {
			projects = append(projects, project)
		}

		return true
	})

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].GetCreatedAt().Seconds > projects[j].GetCreatedAt().Seconds
	})

	return projects, nil
}

func (s *storeOnMemory) updateProject(project *pb.Project) (*pb.Project, error) {
	newProject := *project
	s.projects.Set(newProject.GetId(), &newProject)

	return &newProject, nil
}
