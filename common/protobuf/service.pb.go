// Code generated by protoc-gen-go.
// source: service.proto
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

type Service struct {
	Id          int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	GroupName   string                     `protobuf:"bytes,2,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	ProjectId   int32                      `protobuf:"varint,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	ProjectName string                     `protobuf:"bytes,4,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	Name        string                     `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	Protocol    string                     `protobuf:"bytes,6,opt,name=protocol" json:"protocol,omitempty"`
	HostName    string                     `protobuf:"bytes,7,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Ip          string                     `protobuf:"bytes,8,opt,name=ip" json:"ip,omitempty"`
	Port        int32                      `protobuf:"varint,9,opt,name=port" json:"port,omitempty"`
	CreateTime  *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	UpdateTime  *google_protobuf.Timestamp `protobuf:"bytes,11,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Service) GetCreateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Service) GetUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type Services struct {
	Items []*Service `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Services) Reset()                    { *m = Services{} }
func (m *Services) String() string            { return proto.CompactTextString(m) }
func (*Services) ProtoMessage()               {}
func (*Services) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Services) GetItems() []*Service {
	if m != nil {
		return m.Items
	}
	return nil
}

type ServicePageRequest struct {
	Service     *Service     `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	PageRequest *PageRequest `protobuf:"bytes,2,opt,name=page_request,json=pageRequest" json:"page_request,omitempty"`
}

func (m *ServicePageRequest) Reset()                    { *m = ServicePageRequest{} }
func (m *ServicePageRequest) String() string            { return proto.CompactTextString(m) }
func (*ServicePageRequest) ProtoMessage()               {}
func (*ServicePageRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ServicePageRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ServicePageRequest) GetPageRequest() *PageRequest {
	if m != nil {
		return m.PageRequest
	}
	return nil
}

type ServicePageResponse struct {
	Page  int32      `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	Size  int32      `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Count int32      `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Items []*Service `protobuf:"bytes,4,rep,name=items" json:"items,omitempty"`
}

func (m *ServicePageResponse) Reset()                    { *m = ServicePageResponse{} }
func (m *ServicePageResponse) String() string            { return proto.CompactTextString(m) }
func (*ServicePageResponse) ProtoMessage()               {}
func (*ServicePageResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ServicePageResponse) GetItems() []*Service {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "protobuf.Service")
	proto.RegisterType((*Services)(nil), "protobuf.Services")
	proto.RegisterType((*ServicePageRequest)(nil), "protobuf.ServicePageRequest")
	proto.RegisterType((*ServicePageResponse)(nil), "protobuf.ServicePageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for ServiceService service

type ServiceServiceClient interface {
	Create(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error)
	Get(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error)
	Count(ctx context.Context, in *Service, opts ...grpc.CallOption) (*google_protobuf1.Int32Value, error)
	ListAll(ctx context.Context, in *ServicePageRequest, opts ...grpc.CallOption) (*Services, error)
	List(ctx context.Context, in *ServicePageRequest, opts ...grpc.CallOption) (*ServicePageResponse, error)
	Update(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error)
	Delete(ctx context.Context, in *Service, opts ...grpc.CallOption) (*google_protobuf1.Int32Value, error)
}

type serviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewServiceServiceClient(cc *grpc.ClientConn) ServiceServiceClient {
	return &serviceServiceClient{cc}
}

func (c *serviceServiceClient) Create(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/protobuf.ServiceService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Get(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/protobuf.ServiceService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Count(ctx context.Context, in *Service, opts ...grpc.CallOption) (*google_protobuf1.Int32Value, error) {
	out := new(google_protobuf1.Int32Value)
	err := grpc.Invoke(ctx, "/protobuf.ServiceService/Count", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) ListAll(ctx context.Context, in *ServicePageRequest, opts ...grpc.CallOption) (*Services, error) {
	out := new(Services)
	err := grpc.Invoke(ctx, "/protobuf.ServiceService/ListAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) List(ctx context.Context, in *ServicePageRequest, opts ...grpc.CallOption) (*ServicePageResponse, error) {
	out := new(ServicePageResponse)
	err := grpc.Invoke(ctx, "/protobuf.ServiceService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Update(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := grpc.Invoke(ctx, "/protobuf.ServiceService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceServiceClient) Delete(ctx context.Context, in *Service, opts ...grpc.CallOption) (*google_protobuf1.Int32Value, error) {
	out := new(google_protobuf1.Int32Value)
	err := grpc.Invoke(ctx, "/protobuf.ServiceService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceService service

type ServiceServiceServer interface {
	Create(context.Context, *Service) (*Service, error)
	Get(context.Context, *Service) (*Service, error)
	Count(context.Context, *Service) (*google_protobuf1.Int32Value, error)
	ListAll(context.Context, *ServicePageRequest) (*Services, error)
	List(context.Context, *ServicePageRequest) (*ServicePageResponse, error)
	Update(context.Context, *Service) (*Service, error)
	Delete(context.Context, *Service) (*google_protobuf1.Int32Value, error)
}

func RegisterServiceServiceServer(s *grpc.Server, srv ServiceServiceServer) {
	s.RegisterService(&_ServiceService_serviceDesc, srv)
}

func _ServiceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ServiceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Create(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ServiceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Get(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ServiceService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Count(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicePageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ServiceService/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).ListAll(ctx, req.(*ServicePageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicePageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ServiceService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).List(ctx, req.(*ServicePageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ServiceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Update(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ServiceService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServiceServer).Delete(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.ServiceService",
	HandlerType: (*ServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ServiceService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServiceService_Get_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _ServiceService_Count_Handler,
		},
		{
			MethodName: "ListAll",
			Handler:    _ServiceService_ListAll_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ServiceService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServiceService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ServiceService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor4 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xdb, 0x8a, 0xdb, 0x3c,
	0x10, 0xc7, 0xc9, 0x39, 0x99, 0xe4, 0x5b, 0xf8, 0xd4, 0x16, 0x8c, 0xb7, 0x0b, 0x6d, 0x6e, 0x5a,
	0x28, 0x38, 0x90, 0x14, 0xba, 0x65, 0xaf, 0xba, 0x29, 0x94, 0x85, 0x52, 0x16, 0xf7, 0x70, 0x1b,
	0x1c, 0x47, 0xeb, 0x68, 0xb1, 0x2d, 0xd5, 0x92, 0x5b, 0x7a, 0x78, 0x8a, 0xbe, 0x63, 0xdf, 0xa3,
	0xd2, 0x48, 0x8e, 0xb7, 0x38, 0xdd, 0xec, 0x55, 0x46, 0xff, 0xff, 0xfc, 0x26, 0xd2, 0xcc, 0x60,
	0xf8, 0x4f, 0xd2, 0xe2, 0x0b, 0x8b, 0x69, 0x20, 0x0a, 0xae, 0x38, 0x19, 0xe2, 0xcf, 0xba, 0xbc,
	0xf2, 0xcf, 0x12, 0xa6, 0xb6, 0xe5, 0x3a, 0x88, 0x79, 0x36, 0x4b, 0x78, 0x1a, 0xe5, 0xc9, 0xac,
	0xf2, 0x66, 0x42, 0x7d, 0x13, 0x54, 0xce, 0x14, 0xcb, 0xa8, 0x54, 0x51, 0x26, 0xea, 0xc8, 0x96,
	0xf1, 0x5f, 0x1e, 0x86, 0xbf, 0x16, 0x91, 0x10, 0xb4, 0xa8, 0x03, 0x87, 0x4e, 0x34, 0x93, 0xf1,
	0xdc, 0x9e, 0xa6, 0xbf, 0xdb, 0x30, 0x78, 0x6f, 0x6f, 0x48, 0x8e, 0xa0, 0xcd, 0x36, 0x5e, 0xeb,
	0x51, 0xeb, 0x69, 0x2f, 0xd4, 0x11, 0x39, 0x01, 0x48, 0x0a, 0x5e, 0x8a, 0x55, 0x1e, 0x65, 0xd4,
	0x6b, 0x6b, 0x7d, 0x14, 0x8e, 0x50, 0x79, 0xa7, 0x05, 0x63, 0xeb, 0x1a, 0xd7, 0x34, 0x56, 0x2b,
	0x8d, 0x75, 0x10, 0x1b, 0x39, 0xe5, 0x62, 0x43, 0x1e, 0xc3, 0xa4, 0xb2, 0x91, 0xef, 0x22, 0x3f,
	0x76, 0x1a, 0x56, 0x20, 0xd0, 0x45, 0xab, 0x87, 0x16, 0xc6, 0xc4, 0x07, 0xdb, 0xa2, 0x98, 0xa7,
	0x5e, 0x1f, 0xf5, 0xdd, 0x99, 0x1c, 0xc3, 0x68, 0xcb, 0xa5, 0xab, 0x37, 0xb0, 0xa6, 0x11, 0xb0,
	0x98, 0xb9, 0xbd, 0xf0, 0x86, 0xa8, 0xea, 0xc8, 0x14, 0x17, 0xbc, 0x50, 0xde, 0x08, 0x2f, 0x86,
	0x31, 0x39, 0x83, 0x71, 0x5c, 0xd0, 0x48, 0xd1, 0x95, 0x69, 0xa8, 0x07, 0xda, 0x1a, 0xcf, 0xfd,
	0x20, 0xe1, 0x3c, 0x49, 0xdd, 0x84, 0x74, 0x07, 0x83, 0x0f, 0x55, 0xb7, 0x43, 0xb0, 0xe9, 0x46,
	0x30, 0x70, 0x29, 0x36, 0x3b, 0x78, 0x7c, 0x18, 0xb6, 0xe9, 0x46, 0x98, 0x2e, 0x60, 0xe8, 0xda,
	0x2c, 0xc9, 0x13, 0xe8, 0x31, 0x45, 0x33, 0xa9, 0x5b, 0xdd, 0xd1, 0x25, 0xfe, 0xaf, 0x59, 0x97,
	0x12, 0x5a, 0x7f, 0xfa, 0x03, 0x88, 0x53, 0x2e, 0xa3, 0x84, 0x86, 0xf4, 0x73, 0xa9, 0x2b, 0x93,
	0x67, 0x30, 0x70, 0x3b, 0x85, 0xb3, 0xda, 0x5b, 0xa0, 0xca, 0x20, 0xa7, 0x7a, 0x0a, 0x9a, 0x5d,
	0x15, 0x16, 0xc6, 0x29, 0x8e, 0xe7, 0x0f, 0x6a, 0xe2, 0x46, 0x65, 0x3d, 0x9c, 0xfa, 0x30, 0xfd,
	0x09, 0xf7, 0xfe, 0xfa, 0x73, 0x29, 0x78, 0x2e, 0x71, 0x66, 0x26, 0xcb, 0xad, 0x09, 0xc6, 0x46,
	0x93, 0xec, 0xbb, 0x5d, 0x11, 0xad, 0x99, 0x98, 0xdc, 0x87, 0x5e, 0xcc, 0xcb, 0x5c, 0xb9, 0xc5,
	0xb0, 0x87, 0xfa, 0xe9, 0xdd, 0xdb, 0x9f, 0x3e, 0xff, 0xd5, 0x81, 0x23, 0x27, 0x55, 0xeb, 0x19,
	0x40, 0x7f, 0x89, 0xd3, 0x20, 0x4d, 0xcc, 0x6f, 0x4a, 0xba, 0x4f, 0x9d, 0x37, 0x54, 0xdd, 0x31,
	0xf9, 0x05, 0xf4, 0x96, 0x78, 0xc3, 0x3d, 0xe9, 0xc7, 0x8d, 0x19, 0x5f, 0xe4, 0x6a, 0x31, 0xff,
	0x14, 0xa5, 0xa5, 0xd9, 0x8a, 0xc1, 0x5b, 0x26, 0xd5, 0xab, 0x34, 0x25, 0x0f, 0x1b, 0xe8, 0x8d,
	0xe6, 0xfa, 0xa4, 0xe1, 0x4a, 0xb2, 0x84, 0xae, 0x81, 0x0f, 0x90, 0x27, 0xff, 0x70, 0xdd, 0x44,
	0x74, 0x5f, 0x3e, 0xe2, 0xa2, 0xdd, 0xf1, 0xa9, 0xa7, 0xd0, 0x7f, 0x4d, 0x53, 0xba, 0x3f, 0xff,
	0xb6, 0xb7, 0x9e, 0x3f, 0x87, 0xa9, 0xfe, 0x78, 0x04, 0xd7, 0x4c, 0x6e, 0xcb, 0xab, 0x28, 0x0f,
	0xb6, 0x94, 0xa9, 0x32, 0x88, 0x53, 0x46, 0x73, 0xb5, 0x03, 0xce, 0x27, 0xd5, 0x25, 0x8d, 0x70,
	0xd9, 0x5a, 0xf7, 0xd1, 0x59, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xcd, 0xd9, 0x52, 0x0a,
	0x05, 0x00, 0x00,
}