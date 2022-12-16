package usecase

import (
	"gorm.io/gorm"
	"peanut/domain"
	"peanut/repository"
)

type TodoUsecase interface {
	CreateTodo(userID int, td domain.Todo) (err error)
	ListTodo(userID int) (todos []domain.ListTodo, err error)
	UpdateTodo(userID int, td domain.Todo) (err error)
	DeleteTodo(userID int, todoID int) (err error)
}

type todoUsecase struct {
	TodoRepo repository.TodoRepo
}

func NewTodoUsecase(db *gorm.DB) TodoUsecase {
	return &todoUsecase{
		TodoRepo: repository.NewTodoRepo(db),
	}
}

func (uc *todoUsecase) CreateTodo(userID int, td domain.Todo) (err error) {
	_, err = uc.TodoRepo.CreateTodo(td, uint(userID))
	if err != nil {
		return err
	}
	return
}

func (uc *todoUsecase) ListTodo(userID int) (todos []domain.ListTodo, err error) {
	todos, err = uc.TodoRepo.ListTodo(userID)
	if err != nil {
		return nil, err
	}
	return
}

func (uc *todoUsecase) UpdateTodo(userID int, td domain.Todo) (err error) {
	err = uc.TodoRepo.UpdateTodo(userID, td)
	if err != nil {
		return err
	}
	return
}

func (uc *todoUsecase) DeleteTodo(userID int, todoID int) (err error) {
	err = uc.TodoRepo.DeleteTodo(userID, todoID)
	if err != nil {
		return err
	}
	return
}
