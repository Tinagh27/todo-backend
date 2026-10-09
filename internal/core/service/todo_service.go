package service

import (
	"context"
	"strings"
	"time"

	"todo-backend/internal/core/domain"
	port "todo-backend/internal/core/port/outbound"
)

type TodoService struct {
	repo port.TodoRepository
}

func NewTodoService(repo port.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(ctx context.Context, title, description string) (*domain.Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, domain.ErrInvalidInput
	}

	now := time.Now().UTC()
	todo := &domain.Todo{
		Title:       title,
		Description: strings.TrimSpace(description),
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) GetByID(ctx context.Context, id int64) (*domain.Todo, error) {
	if id <= 0 {
		return nil, domain.ErrInvalidInput
	}
	return s.repo.GetByID(ctx, id)
}

func (s *TodoService) GetAll(ctx context.Context) ([]domain.Todo, error) {
	return s.repo.GetAll(ctx)
}

func (s *TodoService) Update(ctx context.Context, id int64, title, description string, completed bool) (*domain.Todo, error) {
	if id <= 0 {
		return nil, domain.ErrInvalidInput
	}

	title = strings.TrimSpace(title)
	if title == "" {
		return nil, domain.ErrInvalidInput
	}

	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	todo.Title = title
	todo.Description = strings.TrimSpace(description)
	todo.Completed = completed
	todo.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return domain.ErrInvalidInput
	}
	return s.repo.Delete(ctx, id)
}
