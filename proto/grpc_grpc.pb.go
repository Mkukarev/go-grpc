// protoc --go-grpc_out=. --go_out=. --go_opt=paths=source_relative *.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: grpc.proto

package proto

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
	TodoCRUD_CreateTodo_FullMethodName  = "/todo.TodoCRUD/CreateTodo"
	TodoCRUD_GetTodo_FullMethodName     = "/todo.TodoCRUD/GetTodo"
	TodoCRUD_GetAllTodos_FullMethodName = "/todo.TodoCRUD/GetAllTodos"
	TodoCRUD_UpdateTodo_FullMethodName  = "/todo.TodoCRUD/UpdateTodo"
	TodoCRUD_DeleteTodo_FullMethodName  = "/todo.TodoCRUD/DeleteTodo"
)

// TodoCRUDClient is the client API for TodoCRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoCRUDClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoMessage, opts ...grpc.CallOption) (*Todo, error)
	GetTodo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Todo, error)
	GetAllTodos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllTodosResponse, error)
	UpdateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	DeleteTodo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type todoCRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoCRUDClient(cc grpc.ClientConnInterface) TodoCRUDClient {
	return &todoCRUDClient{cc}
}

func (c *todoCRUDClient) CreateTodo(ctx context.Context, in *CreateTodoMessage, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, TodoCRUD_CreateTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoCRUDClient) GetTodo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, TodoCRUD_GetTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoCRUDClient) GetAllTodos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllTodosResponse, error) {
	out := new(GetAllTodosResponse)
	err := c.cc.Invoke(ctx, TodoCRUD_GetAllTodos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoCRUDClient) UpdateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, TodoCRUD_UpdateTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoCRUDClient) DeleteTodo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, TodoCRUD_DeleteTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoCRUDServer is the server API for TodoCRUD service.
// All implementations must embed UnimplementedTodoCRUDServer
// for forward compatibility
type TodoCRUDServer interface {
	CreateTodo(context.Context, *CreateTodoMessage) (*Todo, error)
	GetTodo(context.Context, *Id) (*Todo, error)
	GetAllTodos(context.Context, *Empty) (*GetAllTodosResponse, error)
	UpdateTodo(context.Context, *Todo) (*Todo, error)
	DeleteTodo(context.Context, *Id) (*Empty, error)
	mustEmbedUnimplementedTodoCRUDServer()
}

// UnimplementedTodoCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedTodoCRUDServer struct {
}

func (UnimplementedTodoCRUDServer) CreateTodo(context.Context, *CreateTodoMessage) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoCRUDServer) GetTodo(context.Context, *Id) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (UnimplementedTodoCRUDServer) GetAllTodos(context.Context, *Empty) (*GetAllTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodos not implemented")
}
func (UnimplementedTodoCRUDServer) UpdateTodo(context.Context, *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedTodoCRUDServer) DeleteTodo(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (UnimplementedTodoCRUDServer) mustEmbedUnimplementedTodoCRUDServer() {}

// UnsafeTodoCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoCRUDServer will
// result in compilation errors.
type UnsafeTodoCRUDServer interface {
	mustEmbedUnimplementedTodoCRUDServer()
}

func RegisterTodoCRUDServer(s grpc.ServiceRegistrar, srv TodoCRUDServer) {
	s.RegisterService(&TodoCRUD_ServiceDesc, srv)
}

func _TodoCRUD_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoCRUDServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoCRUD_CreateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoCRUDServer).CreateTodo(ctx, req.(*CreateTodoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoCRUD_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoCRUDServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoCRUD_GetTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoCRUDServer).GetTodo(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoCRUD_GetAllTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoCRUDServer).GetAllTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoCRUD_GetAllTodos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoCRUDServer).GetAllTodos(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoCRUD_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoCRUDServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoCRUD_UpdateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoCRUDServer).UpdateTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoCRUD_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoCRUDServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoCRUD_DeleteTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoCRUDServer).DeleteTodo(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoCRUD_ServiceDesc is the grpc.ServiceDesc for TodoCRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoCRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoCRUD",
	HandlerType: (*TodoCRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoCRUD_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoCRUD_GetTodo_Handler,
		},
		{
			MethodName: "GetAllTodos",
			Handler:    _TodoCRUD_GetAllTodos_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoCRUD_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoCRUD_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
