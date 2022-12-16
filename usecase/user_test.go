package usecase_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"peanut/domain"
)

var _ = Describe("User", func() {
	var u, createdU domain.User
	BeforeEach(func() {
		u = domain.User{
			ID:       1,
			Username: "hung",
			Email:    "hung@example.com",
		}
		createdU = domain.User{}
	})

	Describe("API Create", func() {
		Context("with existed username", func() {
			It("should be error", func() {
				// TODO: fill in your test in this case
			})
		})
		Context("with new user", func() {
			It("should be success", func() {
				// prepare
				userRepo.EXPECT().CreateUser(ctx, u).Return(&u, nil)
				// do
				err := userUc.CreateUser(ctx, u)
				// check
				Expect(err).To(BeNil())
			})
		})
		Context("with database error response", func() {
			It("should be err", func() {
				// prepare
				userRepo.EXPECT().CreateUser(ctx, u).
					Return(nil, fmt.Errorf("database error"))
				// do
				err := userUc.CreateUser(ctx, u)
				// check
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("API Get user", func() {
		Context("with existed id", func() {
			It("should be return user", func() {
				// Do something
				Expect(u).NotTo(Equal(createdU))
			})
		})
	})
})
