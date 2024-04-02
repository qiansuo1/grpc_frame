// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: todo/v2/todo.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TodoService_AddTask_FullMethodName     = "/todo.v2.TodoService/AddTask"
	TodoService_ListTasks_FullMethodName   = "/todo.v2.TodoService/ListTasks"
	TodoService_UpdateTasks_FullMethodName = "/todo.v2.TodoService/UpdateTasks"
	TodoService_DeleteTasks_FullMethodName = "/todo.v2.TodoService/DeleteTasks"
)

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*AddTaskResponse, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (TodoService_ListTasksClient, error)
	UpdateTasks(ctx context.Context, opts ...grpc.CallOption) (TodoService_UpdateTasksClient, error)
	DeleteTasks(ctx context.Context, opts ...grpc.CallOption) (TodoService_DeleteTasksClient, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*AddTaskResponse, error) {
	out := new(AddTaskResponse)
	err := c.cc.Invoke(ctx, TodoService_AddTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (TodoService_ListTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoService_ServiceDesc.Streams[0], TodoService_ListTasks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceListTasksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoService_ListTasksClient interface {
	Recv() (*ListTasksResponse, error)
	grpc.ClientStream
}

type todoServiceListTasksClient struct {
	grpc.ClientStream
}

func (x *todoServiceListTasksClient) Recv() (*ListTasksResponse, error) {
	m := new(ListTasksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoServiceClient) UpdateTasks(ctx context.Context, opts ...grpc.CallOption) (TodoService_UpdateTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoService_ServiceDesc.Streams[1], TodoService_UpdateTasks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceUpdateTasksClient{stream}
	return x, nil
}

type TodoService_UpdateTasksClient interface {
	Send(*UpdateTasksRequest) error
	CloseAndRecv() (*UpdateTasksResponse, error)
	grpc.ClientStream
}

type todoServiceUpdateTasksClient struct {
	grpc.ClientStream
}

func (x *todoServiceUpdateTasksClient) Send(m *UpdateTasksRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *todoServiceUpdateTasksClient) CloseAndRecv() (*UpdateTasksResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateTasksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoServiceClient) DeleteTasks(ctx context.Context, opts ...grpc.CallOption) (TodoService_DeleteTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoService_ServiceDesc.Streams[2], TodoService_DeleteTasks_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceDeleteTasksClient{stream}
	return x, nil
}

type TodoService_DeleteTasksClient interface {
	Send(*DeleteTasksRequest) error
	Recv() (*DeleteTasksResponse, error)
	grpc.ClientStream
}

type todoServiceDeleteTasksClient struct {
	grpc.ClientStream
}

func (x *todoServiceDeleteTasksClient) Send(m *DeleteTasksRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *todoServiceDeleteTasksClient) Recv() (*DeleteTasksResponse, error) {
	m := new(DeleteTasksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	AddTask(context.Context, *AddTaskRequest) (*AddTaskResponse, error)
	ListTasks(*ListTasksRequest, TodoService_ListTasksServer) error
	UpdateTasks(TodoService_UpdateTasksServer) error
	DeleteTasks(TodoService_DeleteTasksServer) error
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) AddTask(context.Context, *AddTaskRequest) (*AddTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (UnimplementedTodoServiceServer) ListTasks(*ListTasksRequest, TodoService_ListTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedTodoServiceServer) UpdateTasks(TodoService_UpdateTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateTasks not implemented")
}
func (UnimplementedTodoServiceServer) DeleteTasks(TodoService_DeleteTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteTasks not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoService_AddTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).AddTask(ctx, req.(*AddTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTasksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServiceServer).ListTasks(m, &todoServiceListTasksServer{stream})
}

type TodoService_ListTasksServer interface {
	Send(*ListTasksResponse) error
	grpc.ServerStream
}

type todoServiceListTasksServer struct {
	grpc.ServerStream
}

func (x *todoServiceListTasksServer) Send(m *ListTasksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TodoService_UpdateTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TodoServiceServer).UpdateTasks(&todoServiceUpdateTasksServer{stream})
}

type TodoService_UpdateTasksServer interface {
	SendAndClose(*UpdateTasksResponse) error
	Recv() (*UpdateTasksRequest, error)
	grpc.ServerStream
}

type todoServiceUpdateTasksServer struct {
	grpc.ServerStream
}

func (x *todoServiceUpdateTasksServer) SendAndClose(m *UpdateTasksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *todoServiceUpdateTasksServer) Recv() (*UpdateTasksRequest, error) {
	m := new(UpdateTasksRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TodoService_DeleteTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TodoServiceServer).DeleteTasks(&todoServiceDeleteTasksServer{stream})
}

type TodoService_DeleteTasksServer interface {
	Send(*DeleteTasksResponse) error
	Recv() (*DeleteTasksRequest, error)
	grpc.ServerStream
}

type todoServiceDeleteTasksServer struct {
	grpc.ServerStream
}

func (x *todoServiceDeleteTasksServer) Send(m *DeleteTasksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *todoServiceDeleteTasksServer) Recv() (*DeleteTasksRequest, error) {
	m := new(DeleteTasksRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.v2.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTask",
			Handler:    _TodoService_AddTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTasks",
			Handler:       _TodoService_ListTasks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateTasks",
			Handler:       _TodoService_UpdateTasks_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DeleteTasks",
			Handler:       _TodoService_DeleteTasks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "todo/v2/todo.proto",
}
