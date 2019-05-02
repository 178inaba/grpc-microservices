// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task/task.proto

package task

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	Status_UNKNOWN   Status = 0
	Status_WAITING   Status = 1
	Status_WORKING   Status = 2
	Status_COMPLETED Status = 3
)

var Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "WAITING",
	2: "WORKING",
	3: "COMPLETED",
}

var Status_value = map[string]int32{
	"UNKNOWN":   0,
	"WAITING":   1,
	"WORKING":   2,
	"COMPLETED": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{0}
}

type Task struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               Status               `protobuf:"varint,3,opt,name=status,proto3,enum=task.Status" json:"status,omitempty"`
	ProjectId            uint64               `protobuf:"varint,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	UserId               uint64               `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNKNOWN
}

func (m *Task) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *Task) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Task) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Task) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateTaskRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ProjectId            uint64   `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{1}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTaskRequest) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

type CreateTaskResponse struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{2}
}

func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type FindTasksResponse struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindTasksResponse) Reset()         { *m = FindTasksResponse{} }
func (m *FindTasksResponse) String() string { return proto.CompactTextString(m) }
func (*FindTasksResponse) ProtoMessage()    {}
func (*FindTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{3}
}

func (m *FindTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTasksResponse.Unmarshal(m, b)
}
func (m *FindTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTasksResponse.Marshal(b, m, deterministic)
}
func (m *FindTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTasksResponse.Merge(m, src)
}
func (m *FindTasksResponse) XXX_Size() int {
	return xxx_messageInfo_FindTasksResponse.Size(m)
}
func (m *FindTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindTasksResponse proto.InternalMessageInfo

func (m *FindTasksResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type FindProjectTasksRequest struct {
	ProjectId            uint64   `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindProjectTasksRequest) Reset()         { *m = FindProjectTasksRequest{} }
func (m *FindProjectTasksRequest) String() string { return proto.CompactTextString(m) }
func (*FindProjectTasksRequest) ProtoMessage()    {}
func (*FindProjectTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{4}
}

func (m *FindProjectTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindProjectTasksRequest.Unmarshal(m, b)
}
func (m *FindProjectTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindProjectTasksRequest.Marshal(b, m, deterministic)
}
func (m *FindProjectTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindProjectTasksRequest.Merge(m, src)
}
func (m *FindProjectTasksRequest) XXX_Size() int {
	return xxx_messageInfo_FindProjectTasksRequest.Size(m)
}
func (m *FindProjectTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindProjectTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindProjectTasksRequest proto.InternalMessageInfo

func (m *FindProjectTasksRequest) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

type FindProjectTasksResponse struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindProjectTasksResponse) Reset()         { *m = FindProjectTasksResponse{} }
func (m *FindProjectTasksResponse) String() string { return proto.CompactTextString(m) }
func (*FindProjectTasksResponse) ProtoMessage()    {}
func (*FindProjectTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{5}
}

func (m *FindProjectTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindProjectTasksResponse.Unmarshal(m, b)
}
func (m *FindProjectTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindProjectTasksResponse.Marshal(b, m, deterministic)
}
func (m *FindProjectTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindProjectTasksResponse.Merge(m, src)
}
func (m *FindProjectTasksResponse) XXX_Size() int {
	return xxx_messageInfo_FindProjectTasksResponse.Size(m)
}
func (m *FindProjectTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindProjectTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindProjectTasksResponse proto.InternalMessageInfo

func (m *FindProjectTasksResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type UpdateTaskRequest struct {
	TaskId               uint64   `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=task.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{6}
}

func (m *UpdateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskRequest.Unmarshal(m, b)
}
func (m *UpdateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskRequest.Merge(m, src)
}
func (m *UpdateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskRequest.Size(m)
}
func (m *UpdateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskRequest proto.InternalMessageInfo

func (m *UpdateTaskRequest) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *UpdateTaskRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTaskRequest) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNKNOWN
}

type UpdateTaskResponse struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskResponse) Reset()         { *m = UpdateTaskResponse{} }
func (m *UpdateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskResponse) ProtoMessage()    {}
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{7}
}

func (m *UpdateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskResponse.Unmarshal(m, b)
}
func (m *UpdateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskResponse.Merge(m, src)
}
func (m *UpdateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskResponse.Size(m)
}
func (m *UpdateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskResponse proto.InternalMessageInfo

func (m *UpdateTaskResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterEnum("task.Status", Status_name, Status_value)
	proto.RegisterType((*Task)(nil), "task.Task")
	proto.RegisterType((*CreateTaskRequest)(nil), "task.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "task.CreateTaskResponse")
	proto.RegisterType((*FindTasksResponse)(nil), "task.FindTasksResponse")
	proto.RegisterType((*FindProjectTasksRequest)(nil), "task.FindProjectTasksRequest")
	proto.RegisterType((*FindProjectTasksResponse)(nil), "task.FindProjectTasksResponse")
	proto.RegisterType((*UpdateTaskRequest)(nil), "task.UpdateTaskRequest")
	proto.RegisterType((*UpdateTaskResponse)(nil), "task.UpdateTaskResponse")
}

func init() { proto.RegisterFile("task/task.proto", fileDescriptor_8e8f2b86464a95fe) }

var fileDescriptor_8e8f2b86464a95fe = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0xd6, 0x71, 0x1d, 0x79, 0xf2, 0x51, 0x92, 0x3d, 0x10, 0xcb, 0xa8, 0xc5, 0xb2, 0x38,
	0x44, 0x48, 0xd8, 0x6a, 0x0a, 0xa2, 0x48, 0x15, 0x55, 0x28, 0x29, 0xb2, 0x0a, 0x49, 0x71, 0x53,
	0x55, 0xe2, 0x52, 0x39, 0xf6, 0x36, 0x98, 0xe2, 0xd8, 0x78, 0xd7, 0x48, 0xfc, 0x1b, 0x7e, 0x29,
	0x42, 0xbb, 0x6b, 0x27, 0xa9, 0x0d, 0xa2, 0x12, 0x97, 0x68, 0x67, 0xde, 0xcc, 0x9b, 0x99, 0x37,
	0x13, 0xc3, 0x7d, 0x16, 0xd0, 0x1b, 0x97, 0xff, 0x38, 0x59, 0x9e, 0xb2, 0x14, 0xab, 0xfc, 0x6d,
	0x3e, 0x5a, 0xa4, 0xe9, 0xe2, 0x0b, 0x71, 0x85, 0x6f, 0x5e, 0x5c, 0xbb, 0x2c, 0x4e, 0x08, 0x65,
	0x41, 0x92, 0xc9, 0x30, 0xf3, 0x61, 0x3d, 0x80, 0x24, 0x19, 0xfb, 0x2e, 0x41, 0xfb, 0x27, 0x02,
	0x75, 0x16, 0xd0, 0x1b, 0xbc, 0x0d, 0x4a, 0x1c, 0x19, 0xc8, 0x42, 0x03, 0xd5, 0x57, 0xe2, 0x08,
	0x63, 0x50, 0x97, 0x41, 0x42, 0x0c, 0xc5, 0x42, 0x03, 0xdd, 0x17, 0x6f, 0xfc, 0x18, 0x34, 0xca,
	0x02, 0x56, 0x50, 0xa3, 0x65, 0xa1, 0xc1, 0xf6, 0xf0, 0x7f, 0x47, 0x74, 0x73, 0x2e, 0x7c, 0x7e,
	0x89, 0xe1, 0x1d, 0x80, 0x2c, 0x4f, 0x3f, 0x93, 0x90, 0x5d, 0xc5, 0x91, 0xa1, 0x0a, 0x46, 0xbd,
	0xf4, 0x78, 0x11, 0xee, 0x43, 0xbb, 0xa0, 0x24, 0xe7, 0xd8, 0x96, 0xc0, 0x34, 0x6e, 0x7a, 0x11,
	0x7e, 0x09, 0x10, 0xe6, 0x24, 0x60, 0x24, 0xba, 0x0a, 0x98, 0xa1, 0x59, 0x68, 0xd0, 0x19, 0x9a,
	0x8e, 0x6c, 0xde, 0xa9, 0x9a, 0x77, 0x66, 0xd5, 0x74, 0xbe, 0x5e, 0x46, 0x8f, 0x18, 0x4f, 0x2d,
	0xb2, 0xa8, 0x4a, 0x6d, 0xff, 0x3d, 0xb5, 0x8c, 0x1e, 0x31, 0xfb, 0x04, 0x7a, 0xc7, 0x82, 0x87,
	0xab, 0xe0, 0x93, 0xaf, 0x05, 0xa1, 0x6c, 0x35, 0x3c, 0xda, 0x18, 0xfe, 0xf6, 0x58, 0x4a, 0x6d,
	0x2c, 0xfb, 0x19, 0xe0, 0x4d, 0x1e, 0x9a, 0xa5, 0x4b, 0x4a, 0xf0, 0x2e, 0x88, 0x25, 0x09, 0xa2,
	0xce, 0x10, 0xa4, 0x5e, 0x22, 0x42, 0xf8, 0xed, 0xe7, 0xd0, 0x3b, 0x89, 0x97, 0x11, 0xf7, 0xd0,
	0x55, 0x92, 0x05, 0x5b, 0x1c, 0xa4, 0x06, 0xb2, 0x5a, 0xb5, 0x2c, 0x09, 0xd8, 0x07, 0xd0, 0xe7,
	0x69, 0x67, 0xb2, 0x7a, 0x99, 0x2d, 0x5b, 0xbf, 0xdd, 0x26, 0xaa, 0xb7, 0x79, 0x08, 0x46, 0x33,
	0xf3, 0xce, 0x75, 0xaf, 0xa1, 0x77, 0x21, 0x94, 0xdb, 0x14, 0xab, 0x0f, 0x6d, 0x8e, 0xae, 0xcb,
	0x69, 0xdc, 0xf4, 0xfe, 0xe1, 0x84, 0xb8, 0x98, 0x9b, 0x75, 0xee, 0x26, 0xe6, 0x93, 0x57, 0xa0,
	0x49, 0x1e, 0xdc, 0x81, 0xf6, 0xc5, 0xe4, 0x74, 0x32, 0xbd, 0x9c, 0x74, 0xff, 0xe3, 0xc6, 0xe5,
	0xc8, 0x9b, 0x79, 0x93, 0xb7, 0x5d, 0x24, 0x8c, 0xa9, 0x7f, 0xca, 0x0d, 0x05, 0xdf, 0x03, 0xfd,
	0x78, 0xfa, 0xfe, 0xec, 0xdd, 0x78, 0x36, 0x7e, 0xd3, 0x6d, 0x0d, 0x7f, 0x28, 0xd0, 0xe1, 0x74,
	0xe7, 0x24, 0xff, 0x16, 0x87, 0x04, 0x1f, 0x01, 0xac, 0x57, 0x8a, 0xfb, 0xb2, 0x5e, 0xe3, 0x58,
	0x4c, 0xa3, 0x09, 0x94, 0x0d, 0x1f, 0x82, 0xbe, 0xda, 0x2e, 0x7e, 0xd0, 0xb8, 0xc7, 0x31, 0xff,
	0x1f, 0x9a, 0x25, 0x6f, 0xf3, 0x0c, 0x3e, 0x40, 0xb7, 0xbe, 0x2a, 0xbc, 0xb3, 0x0e, 0xfe, 0xcd,
	0xf2, 0xcd, 0xdd, 0x3f, 0xc1, 0x25, 0xe5, 0x11, 0xc0, 0x5a, 0xd7, 0x6a, 0xa2, 0xc6, 0x46, 0xab,
	0x89, 0x9a, 0x2b, 0x78, 0xbd, 0xff, 0x71, 0x6f, 0x11, 0xb3, 0x4f, 0xc5, 0xdc, 0x09, 0xd3, 0xc4,
	0xdd, 0x7b, 0x71, 0x10, 0x2f, 0x83, 0x79, 0xe0, 0x2e, 0xf2, 0x2c, 0x7c, 0x9a, 0xc4, 0x61, 0x9e,
	0x52, 0xa9, 0x1d, 0x95, 0x5f, 0x1b, 0xf1, 0xb5, 0x9a, 0x6b, 0xe2, 0xbd, 0xff, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0xcc, 0x20, 0xc3, 0xc1, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	FindTasks(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FindTasksResponse, error)
	FindProjectTasks(ctx context.Context, in *FindProjectTasksRequest, opts ...grpc.CallOption) (*FindProjectTasksResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) FindTasks(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FindTasksResponse, error) {
	out := new(FindTasksResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/FindTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) FindProjectTasks(ctx context.Context, in *FindProjectTasksRequest, opts ...grpc.CallOption) (*FindProjectTasksResponse, error) {
	out := new(FindProjectTasksResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/FindProjectTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	FindTasks(context.Context, *empty.Empty) (*FindTasksResponse, error)
	FindProjectTasks(context.Context, *FindProjectTasksRequest) (*FindProjectTasksResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
}

// UnimplementedTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (*UnimplementedTaskServiceServer) FindTasks(ctx context.Context, req *empty.Empty) (*FindTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTasks not implemented")
}
func (*UnimplementedTaskServiceServer) FindProjectTasks(ctx context.Context, req *FindProjectTasksRequest) (*FindProjectTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProjectTasks not implemented")
}
func (*UnimplementedTaskServiceServer) UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_FindTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).FindTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/FindTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).FindTasks(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_FindProjectTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProjectTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).FindProjectTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/FindProjectTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).FindProjectTasks(ctx, req.(*FindProjectTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
		{
			MethodName: "FindTasks",
			Handler:    _TaskService_FindTasks_Handler,
		},
		{
			MethodName: "FindProjectTasks",
			Handler:    _TaskService_FindProjectTasks_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskService_UpdateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task/task.proto",
}
