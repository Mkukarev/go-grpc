package store

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Todo struct {
	Id      uuid.UUID
	Title   string
	Content string
}

type todoStore struct {
	databaseUrl string
	db          *pgx.Conn
}

func NewTodoStore(databaseUrl string) *todoStore {
	return &todoStore{
		databaseUrl: databaseUrl,
	}
}

func (store *todoStore) connect(ctx context.Context) {
	conn, err := pgx.Connect(ctx, store.databaseUrl)

	if err != nil {
		fmt.Println("Unable to connect to database: ", err.Error())
	}

	store.db = conn
}

func (store *todoStore) close(ctx context.Context) {
	store.db.Close(ctx)
}

func (store *todoStore) CreateTodo(ctx context.Context, todo Todo) {
	store.connect(ctx)
	defer store.close(ctx)

	query := `INSERT INTO todos (id, title, content) VALUES (@id, @title, @content)`

	args := pgx.NamedArgs{
		"id":      todo.Id,
		"title":   todo.Title,
		"content": todo.Content,
	}

	_, err := store.db.Exec(ctx, query, args)
	if err != nil {
		fmt.Println("unable to insert row:", err.Error())
		return
	}
}

func (store *todoStore) GetAllTodos(ctx context.Context) []Todo {
	store.connect(ctx)
	defer store.close(ctx)

	query := `SELECT id, title, content
						From todos`

	rows, err := store.db.Query(ctx, query)
	if err != nil {
		fmt.Println(123, err.Error())
		return nil
	}

	var todos []Todo

	for rows.Next() {
		var id, title, content string
		rows.Scan(&id, &title, &content)
		d, _ := uuid.Parse(id)

		todos = append(todos, Todo{Id: d, Title: title, Content: content})
	}

	return todos
}

func (store *todoStore) GetTodo(ctx context.Context, id uuid.UUID) (Todo, error) {
	store.connect(ctx)
	defer store.close(ctx)

	query := `Select id, title, content
						From todos
						WHERE id =$1`

	var title, content string

	err := store.db.QueryRow(ctx, query, id).Scan(&id, &title, &content)

	if err != nil {
		return Todo{}, err
	}

	todo := Todo{
		Id:      id,
		Title:   title,
		Content: content,
	}

	return todo, nil
}

// TODO: изучить частичное изменение
func (store *todoStore) UpdateTodo(ctx context.Context, todo Todo) error {
	store.connect(ctx)
	defer store.close(ctx)

	query := `UPDATE todos
						SET title = @title, content = @content
						WHERE id = @id`

	args := pgx.NamedArgs{
		"id":      todo.Id,
		"title":   todo.Title,
		"content": todo.Content,
	}

	_, err := store.db.Exec(ctx, query, args)
	if err != nil {
		fmt.Println("delete row error:", err.Error())
		return err
	}

	return nil
}

func (store *todoStore) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	store.connect(ctx)
	defer store.close(ctx)

	query := `DELETE 
						FROM todos
						WHERE id = @id`

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := store.db.Exec(ctx, query, args)
	if err != nil {
		fmt.Println("delete row error:", err.Error())
		return err
	}

	return nil
}

type Interface interface {
	CreateTodo(ctx context.Context, todo Todo)
	GetAllTodos(ctx context.Context) []Todo
	GetTodo(ctx context.Context, id uuid.UUID) (Todo, error)
	UpdateTodo(ctx context.Context, todo Todo) error
	DeleteTodo(ctx context.Context, id uuid.UUID) error
}
