package main

import (
	"sort"

	"github.com/pkg/errors"

	pbtask "github.com/178inaba/grpc-microservices/proto/task"
	"github.com/178inaba/grpc-microservices/shared/inmemory"
)

type Store interface {
	CreateTask(task *pbtask.Task) (*pbtask.Task, error)
	FindTask(taskID, userID uint64) (*pbtask.Task, error)
	FindTasks(userID uint64) ([]*pbtask.Task, error)
	FindProjectTasks(projectID, userID uint64) ([]*pbtask.Task, error)
	UpdateTask(task *pbtask.Task) (*pbtask.Task, error)
}

type StoreOnMemory struct {
	tasks *inmemory.IndexMap
}

func NewStoreOnMemory() *StoreOnMemory {
	return &StoreOnMemory{inmemory.NewIndexMap()}
}

func (s *StoreOnMemory) CreateTask(task *pbtask.Task) (*pbtask.Task, error) {
	newTask := *task
	idx := s.tasks.Index()
	newTask.Id = idx
	s.tasks.Set(idx, &newTask)

	return &newTask, nil
}

func (s *StoreOnMemory) FindTask(taskID, userID uint64) (*pbtask.Task, error) {
	value, ok := s.tasks.Get(taskID)
	if !ok {
		return nil, errors.New("not found task")
	}

	task := value.(*pbtask.Task)
	if task.UserId != userID {
		return nil, errors.New("not found task")
	}

	return task, nil
}

func (s *StoreOnMemory) FindTasks(userID uint64) ([]*pbtask.Task, error) {
	var tasks []*pbtask.Task
	s.tasks.Range(func(idx uint64, value interface{}) bool {
		task := value.(*pbtask.Task)
		if task.UserId == userID {
			tasks = append(tasks, task)
		}

		return true
	})

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].UpdatedAt.Seconds > tasks[j].UpdatedAt.Seconds
	})

	return tasks, nil
}

func (s *StoreOnMemory) FindProjectTasks(projectID, userID uint64) ([]*pbtask.Task, error) {
	var tasks []*pbtask.Task
	s.tasks.Range(func(idx uint64, value interface{}) bool {
		task := value.(*pbtask.Task)
		if task.ProjectId == projectID && task.UserId == userID {
			tasks = append(tasks, task)
		}

		return true
	})

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].UpdatedAt.Seconds > tasks[j].UpdatedAt.Seconds
	})

	return tasks, nil
}

func (s *StoreOnMemory) UpdateTask(task *pbtask.Task) (*pbtask.Task, error) {
	newTask := *task
	s.tasks.Set(newTask.Id, &newTask)

	return &newTask, nil
}
