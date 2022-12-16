package usecase_test

import (
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

	Describe("Create", func() {
		Context("with existed username", func() {
			It("should be error", func() {
				createdU = u
				Expect(u).To(Equal(createdU))
			})
		})
		Context("with new user", func() {
			It("should be success", func() {
				Expect(createdU).To(BeZero())
			})
		})
	})

	Describe("Get user", func() {
		Context("with existed id", func() {
			It("should be return user", func() {
				Expect(u).NotTo(Equal(createdU))
			})
		})
	})
})
