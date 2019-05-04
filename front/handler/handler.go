package handler

import (
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/178inaba/grpc-microservices/front/session"
	pbactivity "github.com/178inaba/grpc-microservices/proto/activity"
	pbproject "github.com/178inaba/grpc-microservices/proto/project"
	pbtask "github.com/178inaba/grpc-microservices/proto/task"
	pbuser "github.com/178inaba/grpc-microservices/proto/user"
)

type FrontServer struct {
	ActivityClient pbactivity.ActivityServiceClient
	ProjectClient  pbproject.ProjectServiceClient
	TaskClient     pbtask.TaskServiceClient
	UserClient     pbuser.UserServiceClient
	SessionStore   session.Store
}

func (s *FrontServer) Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := s.UserClient.CreateUser(r.Context(), &pbuser.CreateUserRequest{
		Email:    r.Form.Get("email"),
		Password: []byte(r.Form.Get("password")),
	})
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	sessionID := session.ID()
	s.SessionStore.Set(sessionID, resp.GetUser().GetId())
	session.SetSessionIDToResponse(w, sessionID)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *FrontServer) CreateTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	projectIDStr := r.Form.Get("project_id")
	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if _, err := s.TaskClient.CreateTask(r.Context(), &pbtask.CreateTaskRequest{
		Name:      r.Form.Get("name"),
		ProjectId: projectID,
	}); err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	redirectURL := "/"
	if strings.Contains(r.Referer(), "/project/") {
		redirectURL = path.Join("/project", projectIDStr)
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
}
