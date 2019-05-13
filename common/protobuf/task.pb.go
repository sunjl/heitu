// Code generated by protoc-gen-go.
// source: task.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Task struct {
	Id         int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	HostName   string                     `protobuf:"bytes,2,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Ip         string                     `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	Name       string                     `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Content    string                     `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
	Status     string                     `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	CreateTime *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	UpdateTime *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Task) GetCreateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type Tasks struct {
	Items []*Task `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Tasks) Reset()                    { *m = Tasks{} }
func (m *Tasks) String() string            { return proto.CompactTextString(m) }
func (*Tasks) ProtoMessage()               {}
func (*Tasks) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Tasks) GetItems() []*Task {
	if m != nil {
		return m.Items
	}
	return nil
}

type TaskPageRequest struct {
	Task        *Task        `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
	PageRequest *PageRequest `protobuf:"bytes,2,opt,name=page_request,json=pageRequest" json:"page_request,omitempty"`
}

func (m *TaskPageRequest) Reset()                    { *m = TaskPageRequest{} }
func (m *TaskPageRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskPageRequest) ProtoMessage()               {}
func (*TaskPageRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *TaskPageRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *TaskPageRequest) GetPageRequest() *PageRequest {
	if m != nil {
		return m.PageRequest
	}
	return nil
}

type TaskPageResponse struct {
	Page  int32   `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	Size  int32   `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Count int32   `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Items []*Task `protobuf:"bytes,4,rep,name=items" json:"items,omitempty"`
}

func (m *TaskPageResponse) Reset()                    { *m = TaskPageResponse{} }
func (m *TaskPageResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskPageResponse) ProtoMessage()               {}
func (*TaskPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *TaskPageResponse) GetItems() []*Task {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "protobuf.Task")
	proto.RegisterType((*Tasks)(nil), "protobuf.Tasks")
	proto.RegisterType((*TaskPageRequest)(nil), "protobuf.TaskPageRequest")
	proto.RegisterType((*TaskPageResponse)(nil), "protobuf.TaskPageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for TaskService service

type TaskServiceClient interface {
	Greate(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	Get(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	Count(ctx context.Context, in *Task, opts ...grpc.CallOption) (*google_protobuf1.Int32Value, error)
	ListAll(ctx context.Context, in *TaskPageRequest, opts ...grpc.CallOption) (*Tasks, error)
	List(ctx context.Context, in *TaskPageRequest, opts ...grpc.CallOption) (*TaskPageResponse, error)
	Update(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	Delete(ctx context.Context, in *Task, opts ...grpc.CallOption) (*google_protobuf1.Int32Value, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Greate(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/protobuf.TaskService/Greate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Get(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/protobuf.TaskService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Count(ctx context.Context, in *Task, opts ...grpc.CallOption) (*google_protobuf1.Int32Value, error) {
	out := new(google_protobuf1.Int32Value)
	err := grpc.Invoke(ctx, "/protobuf.TaskService/Count", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListAll(ctx context.Context, in *TaskPageRequest, opts ...grpc.CallOption) (*Tasks, error) {
	out := new(Tasks)
	err := grpc.Invoke(ctx, "/protobuf.TaskService/ListAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) List(ctx context.Context, in *TaskPageRequest, opts ...grpc.CallOption) (*TaskPageResponse, error) {
	out := new(TaskPageResponse)
	err := grpc.Invoke(ctx, "/protobuf.TaskService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Update(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/protobuf.TaskService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) Delete(ctx context.Context, in *Task, opts ...grpc.CallOption) (*google_protobuf1.Int32Value, error) {
	out := new(google_protobuf1.Int32Value)
	err := grpc.Invoke(ctx, "/protobuf.TaskService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceServer interface {
	Greate(context.Context, *Task) (*Task, error)
	Get(context.Context, *Task) (*Task, error)
	Count(context.Context, *Task) (*google_protobuf1.Int32Value, error)
	ListAll(context.Context, *TaskPageRequest) (*Tasks, error)
	List(context.Context, *TaskPageRequest) (*TaskPageResponse, error)
	Update(context.Context, *Task) (*Task, error)
	Delete(context.Context, *Task) (*google_protobuf1.Int32Value, error)
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_Greate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Greate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.TaskService/Greate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Greate(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.TaskService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Get(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.TaskService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Count(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.TaskService/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListAll(ctx, req.(*TaskPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.TaskService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).List(ctx, req.(*TaskPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.TaskService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Update(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.TaskService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Delete(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greate",
			Handler:    _TaskService_Greate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TaskService_Get_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _TaskService_Count_Handler,
		},
		{
			MethodName: "ListAll",
			Handler:    _TaskService_ListAll_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TaskService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TaskService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TaskService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor5 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0x13, 0xdb, 0x49, 0xc6, 0x55, 0x8b, 0x56, 0x80, 0x8c, 0x7b, 0x41, 0x16, 0x48, 0xbd,
	0xe0, 0x88, 0x94, 0x4a, 0xa0, 0x8a, 0x03, 0x05, 0xa9, 0x42, 0x42, 0xa8, 0x32, 0x85, 0x6b, 0xe4,
	0xb8, 0x53, 0x67, 0xc1, 0xf6, 0x2e, 0xde, 0x75, 0x11, 0x5c, 0xf9, 0x5c, 0x7e, 0x82, 0xdd, 0x71,
	0x52, 0xa7, 0x69, 0x95, 0x9c, 0x32, 0xef, 0xcd, 0x7b, 0x9b, 0x9d, 0x37, 0x5e, 0x00, 0x9d, 0xaa,
	0x1f, 0xb1, 0xac, 0x85, 0x16, 0x6c, 0x48, 0x3f, 0xb3, 0xe6, 0x2a, 0x3c, 0xc9, 0xb9, 0x9e, 0x37,
	0xb3, 0x38, 0x13, 0xe5, 0x38, 0x17, 0x45, 0x5a, 0xe5, 0xe3, 0x65, 0x6f, 0x2c, 0xf5, 0x6f, 0x89,
	0x6a, 0xac, 0x79, 0x89, 0x4a, 0xa7, 0xa5, 0xec, 0xaa, 0xf6, 0x98, 0xf0, 0xcd, 0x76, 0xf3, 0xaf,
	0x3a, 0x95, 0x12, 0xeb, 0xae, 0x58, 0x58, 0x77, 0x8d, 0xa7, 0x14, 0x55, 0x8b, 0xa2, 0xbf, 0x3d,
	0x70, 0x2e, 0xcc, 0xf5, 0xd8, 0x1e, 0xf4, 0xf8, 0x65, 0xb0, 0xf3, 0x74, 0xe7, 0xd0, 0x4d, 0x4c,
	0xc5, 0x0e, 0x60, 0x34, 0x17, 0x4a, 0x4f, 0xab, 0xb4, 0xc4, 0xa0, 0x67, 0xe8, 0x51, 0x32, 0xb4,
	0xc4, 0x67, 0x83, 0x49, 0x2c, 0x83, 0x3e, 0xb1, 0xa6, 0x62, 0x0c, 0x1c, 0xd2, 0x39, 0xc4, 0x50,
	0xcd, 0x02, 0x18, 0x64, 0xa2, 0xd2, 0x58, 0xe9, 0xc0, 0x25, 0x7a, 0x09, 0xd9, 0x63, 0xf0, 0xcc,
	0x2c, 0xba, 0x51, 0x81, 0x47, 0x8d, 0x05, 0x62, 0x27, 0xe0, 0x67, 0x35, 0xa6, 0x1a, 0xa7, 0x76,
	0xdc, 0x60, 0x60, 0x9a, 0xfe, 0x24, 0x8c, 0x73, 0x21, 0xf2, 0x02, 0xe3, 0xe5, 0x7c, 0xf1, 0xc5,
	0x32, 0x8b, 0x04, 0x5a, 0xb9, 0x25, 0xac, 0xb9, 0x91, 0x97, 0x37, 0xe6, 0xe1, 0x76, 0x73, 0x2b,
	0xb7, 0x44, 0xf4, 0x02, 0x5c, 0x1b, 0x82, 0x62, 0xcf, 0xc0, 0xe5, 0x1a, 0x4b, 0x65, 0x82, 0xe8,
	0x1b, 0xff, 0xde, 0x8a, 0xd1, 0xf4, 0x93, 0xb6, 0x19, 0x09, 0xd8, 0xb7, 0xf0, 0x3c, 0xcd, 0x31,
	0xc1, 0x9f, 0x8d, 0x39, 0x90, 0x45, 0xe0, 0xd8, 0x2d, 0x53, 0x80, 0x77, 0x7d, 0xd4, 0x63, 0xaf,
	0x61, 0x57, 0x1a, 0xcb, 0xb4, 0x6e, 0x3d, 0x94, 0xaa, 0x3f, 0x79, 0xd4, 0x69, 0x57, 0x0e, 0x4c,
	0x7c, 0xd9, 0x81, 0xe8, 0x1a, 0x1e, 0x74, 0x7f, 0xa8, 0xa4, 0xa8, 0x14, 0xda, 0xcc, 0xad, 0x64,
	0xb1, 0x32, 0xaa, 0x2d, 0xa7, 0xf8, 0x9f, 0x76, 0x5f, 0x86, 0xb3, 0x35, 0x7b, 0x08, 0x6e, 0x26,
	0x1a, 0xb3, 0x85, 0x3e, 0x91, 0x2d, 0xe8, 0x06, 0x75, 0x36, 0x0c, 0x3a, 0xf9, 0xd7, 0x03, 0xdf,
	0xe2, 0x2f, 0x58, 0x5f, 0xf3, 0x0c, 0xd9, 0x21, 0x78, 0x67, 0x14, 0x39, 0x5b, 0x33, 0x84, 0x6b,
	0x98, 0x3d, 0x87, 0xfe, 0x19, 0xea, 0xad, 0xb2, 0x57, 0xe0, 0xbe, 0xa7, 0xfb, 0xac, 0x0b, 0x0f,
	0xee, 0x6c, 0xee, 0x63, 0xa5, 0x8f, 0x26, 0xdf, 0xd2, 0xa2, 0x41, 0x76, 0x0c, 0x83, 0x4f, 0x5c,
	0xe9, 0x77, 0x45, 0xc1, 0x9e, 0xdc, 0xf6, 0xad, 0x24, 0x18, 0xee, 0xdf, 0x6e, 0x29, 0xf6, 0x16,
	0x1c, 0x6b, 0xdb, 0xe4, 0x09, 0xef, 0x6b, 0x2d, 0x02, 0x37, 0xc3, 0x7f, 0xa5, 0x4f, 0x66, 0xeb,
	0x54, 0xc7, 0xe0, 0x7d, 0xc0, 0x02, 0xef, 0x51, 0x6e, 0x1a, 0xeb, 0xf4, 0x25, 0x44, 0xe6, 0x6d,
	0xc6, 0xdf, 0xb9, 0x9a, 0x37, 0x57, 0x69, 0x15, 0xcf, 0x91, 0xeb, 0x26, 0xce, 0x0a, 0x6e, 0x1e,
	0xcd, 0x8d, 0xe1, 0x74, 0x44, 0x17, 0xb3, 0xe8, 0x7c, 0x67, 0xe6, 0x11, 0x7d, 0xf4, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0x0f, 0xd1, 0x5a, 0x63, 0x04, 0x00, 0x00,
}
