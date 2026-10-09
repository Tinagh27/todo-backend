package port

import (
	"context"

	"todo-backend/internal/core/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo) error
	GetByID(ctx context.Context, id int64) (*domain.Todo, error)
	GetAll(ctx context.Context) ([]domain.Todo, error)
	Update(ctx context.Context, todo *domain.Todo) error
	Delete(ctx context.Context, id int64) error
}
