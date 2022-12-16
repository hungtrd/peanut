package repository

import (
	"fmt"
	"gorm.io/gorm"
	"peanut/domain"
)

type TodoRepo interface {
	ListTodo(userID int) ([]domain.ListTodo, error)
	CreateTodo(td domain.Todo, userID uint) (todo *domain.Todo, err error)
	UpdateTodo(todoID int, td domain.Todo) error
	DeleteTodo(userID int, todoID int) error
}

type todoRepo struct {
	DB *gorm.DB
}

func NewTodoRepo(db *gorm.DB) TodoRepo {
	return &todoRepo{DB: db}
}

func (r *todoRepo) CreateTodo(td domain.Todo, userID uint) (todo *domain.Todo, err error) {
	todo = &domain.Todo{
		Title:   td.Title,
		UserID:  userID,
		Content: td.Content,
	}

	result := r.DB.Create(todo)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Todo.CreateTodo] failed: %w", result.Error)
		return nil, err
	}

	return
}

func (r *todoRepo) ListTodo(userID int) ([]domain.ListTodo, error) {
	var todos []domain.ListTodo
	result := r.DB.
		Table("todos").
		Select("todos.id, users.username as username, todos.title, todos.content, todos.created_at").
		Joins("JOIN users ON todos.user_id = users.id AND users.id = ?", userID).
		Scan(&todos)
	return todos, result.Error
}

func (r *todoRepo) UpdateTodo(todoID int, td domain.Todo) error {
	result := r.DB.Model(domain.Todo{}).Where("id = ?", todoID).Updates(td)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *todoRepo) DeleteTodo(userID int, todoID int) error {
	result := r.DB.Where("id = ? AND user_id = ?", todoID, userID).Delete(&domain.Todo{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
