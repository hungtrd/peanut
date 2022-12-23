package usecase_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"peanut/domain"
)

var _ = Describe("Todo", func() {

	var listTodo []domain.ListTodo
	var todo1 domain.Todo
	var itemTodo1 domain.ListTodo
	var itemTodo2 domain.ListTodo

	BeforeEach(func() {
		todo1 = domain.Todo{
			Title:   "title 1",
			Content: "Content 1",
			UserID:  1,
		}
		todo1.ID = uint(1)

		itemTodo1 = domain.ListTodo{
			Title:    "title 1",
			Content:  "Content 1",
			Username: "user 1",
		}
		itemTodo1.ID = 1

		itemTodo2 = domain.ListTodo{
			Title:    "title 2",
			Content:  "Content 2",
			Username: "user 1",
		}
		itemTodo2.ID = 2
	})

	Describe("API Get todo", func() {

		listTodo = []domain.ListTodo{
			itemTodo1,
			itemTodo2,
		}
		Context("- Get list todo of user", func() {
			It("should be return list", func() {
				// prepare
				todoRepo.EXPECT().ListTodo(1).Return(listTodo, nil)
				// do
				td, err := todoUc.ListTodo(1)
				// check
				Expect(err).To(BeNil())
				Expect(td).To(Equal(listTodo))
			})
		})
	})

	Describe("API Create todo", func() {
		Context("- Create todo of user", func() {
			It("should be success", func() {
				// prepare
				todoRepo.EXPECT().CreateTodo(todo1, todo1.UserID).Return(&todo1, nil)
				// do
				err := todoUc.CreateTodo(int(todo1.UserID), todo1)
				// check
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("API Update todo", func() {
		Context("- Update todo of user", func() {
			It("should be success", func() {
				// prepare
				todoRepo.EXPECT().UpdateTodo(int(todo1.UserID), todo1).Return(nil)
				// do
				err := todoUc.UpdateTodo(int(todo1.UserID), todo1)
				// check
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("API Delete todo", func() {
		Context("- Delete todo of user", func() {
			It("should be success", func() {
				// prepare
				todoRepo.EXPECT().DeleteTodo(int(todo1.UserID), int(todo1.ID)).Return(nil)
				// do
				err := todoUc.DeleteTodo(int(todo1.UserID), int(todo1.ID))
				// check
				Expect(err).To(BeNil())
			})
		})
	})
})
