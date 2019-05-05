package main

import (
	pb "github.com/178inaba/grpc-microservices/proto/user"
	"github.com/178inaba/grpc-microservices/shared/inmemory"
	"github.com/pkg/errors"
)

type store interface {
	createUser(user *pb.User) (*pb.User, error)
	findUser(userID uint64) (*pb.User, error)
	findUserByEmail(email string) (*pb.User, error)
}

type storeOnMemory struct {
	users *inmemory.IndexMap
}

func newStoreOnMemory() *storeOnMemory {
	return &storeOnMemory{inmemory.NewIndexMap()}
}

func (s *storeOnMemory) createUser(user *pb.User) (*pb.User, error) {
	if _, err := s.findUserByEmail(user.Email); err == nil {
		return nil, errors.Errorf("already exists user, email: %s", user.Email)
	}

	newUser := *user
	idx := s.users.Index()
	newUser.Id = idx
	s.users.Set(idx, &newUser)

	return &newUser, nil
}

func (s *storeOnMemory) findUser(userID uint64) (*pb.User, error) {
	value, ok := s.users.Get(userID)
	if !ok {
		return nil, errors.Errorf("not found user, user id: %d", userID)
	}

	return value.(*pb.User), nil
}

func (s *storeOnMemory) findUserByEmail(email string) (*pb.User, error) {
	var user *pb.User
	s.users.Range(func(idx uint64, value interface{}) bool {
		u := value.(*pb.User)
		if u.GetEmail() == email {
			user = u
			return false
		}
		return true
	})

	if user == nil {
		return nil, errors.New("not found user")
	}
	return user, nil
}
