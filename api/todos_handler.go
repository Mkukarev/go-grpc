package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	pb "github.com/Mkukarev/go-grpc/proto"
	"github.com/Mkukarev/go-grpc/store"
)

var (
	BadRequest = errors.New("Bad Request")
)

// TODO: переделать
func (s *Server) GetAllTodos(ctx context.Context, _ *pb.Empty) (*pb.GetAllTodosResponse, error) {
	todos := s.store.GetAllTodos(ctx)
	if todos == nil {
		return nil, nil
	}

	var list []*pb.Todo

	for _, v := range todos {
		list = append(list, &pb.Todo{
			Id:      v.Id.String(),
			Title:   v.Title,
			Content: v.Content,
		})
	}

	data := &pb.GetAllTodosResponse{
		Todos: list,
	}

	return data, nil
}

func (s *Server) GetTodo(ctx context.Context, param *pb.Id) (*pb.Todo, error) {
	id, err := uuid.Parse(param.Id)
	if err != nil {
		return &pb.Todo{}, BadRequest
	}

	todo, err := s.store.GetTodo(ctx, id)
	if err != nil {
		return &pb.Todo{}, BadRequest
	}

	t := &pb.Todo{
		Id:      param.Id,
		Title:   todo.Title,
		Content: todo.Content,
	}

	return t, nil
}

func (s *Server) CreateTodo(ctx context.Context, t *pb.CreateTodoMessage) (*pb.Todo, error) {
	id, _ := uuid.NewUUID()

	// TODO: разобраться с ошибками между сервисами
	if t.Title == "" {
		return nil, BadRequest
	}

	todo := store.Todo{
		Id:      id,
		Title:   t.Title,
		Content: t.Content,
	}

	s.store.CreateTodo(ctx, todo)

	responseTodo := &pb.Todo{
		Id:      id.String(),
		Title:   t.Title,
		Content: t.Content,
	}

	return responseTodo, nil
}

func (s *Server) UpdateTodo(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	id, err := uuid.Parse(todo.Id)
	if err != nil {
		return nil, BadRequest
	}

	if todo.Title == "" {
		return nil, BadRequest
	}

	t := store.Todo{
		Id:      id,
		Title:   todo.Title,
		Content: todo.Content,
	}

	e := s.store.UpdateTodo(ctx, t)
	if e != nil {
		return nil, BadRequest
	}

	responseTodo := &pb.Todo{
		Id:      id.String(),
		Title:   todo.Title,
		Content: todo.Content,
	}

	return responseTodo, nil
}

func (s *Server) DeleteTodo(ctx context.Context, param *pb.Id) (*pb.Empty, error) {
	id, err := uuid.Parse(param.Id)
	if err != nil {
		fmt.Println("error parse uuid:", err.Error())
		return &pb.Empty{}, BadRequest
	}

	e := s.store.DeleteTodo(ctx, id)
	if e != nil {
		fmt.Println("db error", e.Error())
		return &pb.Empty{}, BadRequest
	}

	return &pb.Empty{}, nil
}
