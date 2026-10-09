package memory

import (
	"context"
	"sync"

	"todo-backend/internal/core/domain"
)

type TodoRepository struct {
	mu     sync.RWMutex
	todos  map[int64]*domain.Todo
	nextID int64
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos:  make(map[int64]*domain.Todo),
		nextID: 1,
	}
}

func (r *TodoRepository) Create(_ context.Context, todo *domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo.ID = r.nextID
	r.nextID++

	cloned := *todo
	r.todos[todo.ID] = &cloned
	return nil
}

func (r *TodoRepository) GetByID(_ context.Context, id int64) (*domain.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todo, ok := r.todos[id]
	if !ok {
		return nil, domain.ErrNotFound
	}

	cloned := *todo
	return &cloned, nil
}

func (r *TodoRepository) GetAll(_ context.Context) ([]domain.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		result = append(result, *todo)
	}
	return result, nil
}

func (r *TodoRepository) Update(_ context.Context, todo *domain.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.todos[todo.ID]; !ok {
		return domain.ErrNotFound
	}

	cloned := *todo
	r.todos[todo.ID] = &cloned
	return nil
}

func (r *TodoRepository) Delete(_ context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.todos[id]; !ok {
		return domain.ErrNotFound
	}

	delete(r.todos, id)
	return nil
}
