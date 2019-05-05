package handler

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/178inaba/grpc-microservices/front/session"
	"github.com/178inaba/grpc-microservices/front/support"
	"github.com/178inaba/grpc-microservices/front/template"
	pbactivity "github.com/178inaba/grpc-microservices/proto/activity"
	pbproject "github.com/178inaba/grpc-microservices/proto/project"
	pbtask "github.com/178inaba/grpc-microservices/proto/task"
	pbuser "github.com/178inaba/grpc-microservices/proto/user"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *FrontServer) Logout(w http.ResponseWriter, r *http.Request) {
	sessionID := session.GetSessionIDFromRequest(r)
	s.SessionStore.Delete(sessionID)
	session.DeleteSessionIDFromResponse(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (s *FrontServer) CreateProject(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if _, err := s.ProjectClient.CreateProject(r.Context(), &pbproject.CreateProjectRequest{
		Name: r.Form.Get("name"),
	}); err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *FrontServer) UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectIDStr := mux.Vars(r)["id"]
	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)
	if err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	r.ParseForm()
	if _, err := s.ProjectClient.UpdateProject(r.Context(), &pbproject.UpdateProjectRequest{
		ProjectId:   projectID,
		ProjectName: r.Form.Get("name"),
	}); err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, path.Join("/project", projectIDStr), http.StatusSeeOther)
}

func (s *FrontServer) CreateTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	projectIDStr := r.Form.Get("project_id")
	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if _, err := s.TaskClient.CreateTask(r.Context(), &pbtask.CreateTaskRequest{
		Name:      r.Form.Get("name"),
		ProjectId: projectID,
	}); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	redirectURL := "/"
	if strings.Contains(r.Referer(), "/project/") {
		redirectURL = path.Join("/project", projectIDStr)
	}

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (s *FrontServer) UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	r.ParseForm()
	status, err := strconv.ParseUint(r.Form.Get("status_id"), 10, 32)
	if err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	task, err := s.TaskClient.UpdateTask(r.Context(), &pbtask.UpdateTaskRequest{
		TaskId: taskID,
		Name:   r.Form.Get("name"),
		Status: pbtask.Status(status),
	})
	if err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	redirectURL := "/"
	if strings.Contains(r.Referer(), "/project/") {
		redirectURL = path.Join("/project", fmt.Sprint(task.Task.GetProjectId()))
	}

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (s *FrontServer) ViewHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var in empty.Empty
	activities, err := s.ActivityClient.FindActivities(ctx, &in)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	projects, err := s.ProjectClient.FindProjects(ctx, &in)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tasks, err := s.TaskClient.FindTasks(ctx, &in)
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var activityRows []*ActivityRow
	for _, activity := range activities.GetActivities() {
		activityRows = append(activityRows, &ActivityRow{activity})
	}

	idToPj := make(map[uint64]*pbproject.Project)
	for _, project := range projects.GetProjects() {
		idToPj[project.GetId()] = project
	}

	var taskRows []*TaskRow
	for _, task := range tasks.GetTasks() {
		project := idToPj[task.GetProjectId()]
		taskRows = append(taskRows, &TaskRow{task, project})
	}

	user := support.GetUserFromContext(ctx)
	if err := template.Render(w, "home.html", &HomeContent{
		PageName:     "Home",
		IsLoggedIn:   true,
		UserEmail:    user.Email,
		TaskStatuses: taskStatuses,
		ActivityRows: activityRows,
		Projects:     projects.GetProjects(),
		TaskRows:     taskRows,
	}); err != nil {
		http.Error(w,
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (s *FrontServer) ViewProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	projectID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	project, err := s.ProjectClient.FindProject(ctx, &pbproject.FindProjectRequest{
		ProjectId: projectID,
	})
	if err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	tasks, err := s.TaskClient.FindProjectTasks(ctx, &pbtask.FindProjectTasksRequest{
		ProjectId: projectID,
	})
	if err != nil {
		// TODO Error logging.
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var taskRows []*TaskRow
	for _, task := range tasks.GetTasks() {
		taskRows = append(taskRows, &TaskRow{task, project.GetProject()})
	}

	user := support.GetUserFromContext(ctx)
	template.Render(w, "project.html", &ProjectContent{
		PageName:     "Project",
		IsLoggedIn:   true,
		UserEmail:    user.GetEmail(),
		TaskStatuses: taskStatuses,
		Project:      project.GetProject(),
		TaskRows:     taskRows,
	})
}

type HomeContent struct {
	PageName     string
	IsLoggedIn   bool
	TaskStatuses []TaskStatus
	UserEmail    string
	ActivityRows []*ActivityRow
	Projects     []*pbproject.Project
	TaskRows     []*TaskRow
}

type ProjectContent struct {
	PageName     string
	IsLoggedIn   bool
	TaskStatuses []TaskStatus
	UserEmail    string
	Project      *pbproject.Project
	TaskRows     []*TaskRow
}

type ActivityRow struct {
	activity *pbactivity.Activity
}

func (r *ActivityRow) DateTime() string {
	t, err := ptypes.Timestamp(r.activity.CreatedAt)
	if err != nil {
		return ""
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return t.In(jst).Format("2006/1/2 15:04:05")
}

func (r *ActivityRow) Text() string {
	if msg := new(pbactivity.CreateTaskContent); ptypes.Is(r.activity.Content, msg) {
		proto.Unmarshal(r.activity.Content.Value, msg)
		return fmt.Sprintf("Create a task %q.", msg.TaskName)
	} else if msg := new(pbactivity.UpdateTaskStatusContent); ptypes.Is(r.activity.Content, msg) {
		proto.Unmarshal(r.activity.Content.Value, msg)
		return fmt.Sprintf("Changed the status of task %q to %q.", msg.TaskName, msg.TaskStatus)
	} else if msg := new(pbactivity.CreateProjectContent); ptypes.Is(r.activity.Content, msg) {
		proto.Unmarshal(r.activity.Content.Value, msg)
		return fmt.Sprintf("Create a project %q.", msg.ProjectName)
	}

	return ""
}

type TaskRow struct {
	task    *pbtask.Task
	project *pbproject.Project
}

func (r *TaskRow) ID() uint64 {
	return r.task.Id
}

func (r *TaskRow) Name() string {
	return r.task.Name
}

func (r *TaskRow) ProjectName() string {
	return r.project.Name
}

func (r *TaskRow) Status() int32 {
	return int32(r.task.Status)
}

func (r *TaskRow) StatusName() string {
	return r.task.Status.String()
}

type TaskStatus pbtask.Status

func (s TaskStatus) Status() int32 {
	return int32(s)
}

func (s *TaskStatus) StatusName() string {
	return pbtask.Status_name[s.Status()]
}

var taskStatuses = []TaskStatus{
	TaskStatus(pbtask.Status_WAITING),
	TaskStatus(pbtask.Status_WORKING),
	TaskStatus(pbtask.Status_COMPLETED),
}
