package usecase_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"peanut/domain"
	"peanut/pkg/hash"
	"time"
)

var _ = Describe("User", func() {

	var user domain.User
	var users []domain.User
	var req domain.RequestLogin
	var register domain.User

	BeforeEach(func() {
		register = domain.User{
			Email:    "abc@gmail.com",
			Password: "12345678",
		}
		register.ID = 2
		register.CreatedAt = time.Now()
		register.UpdatedAt = time.Now()

		req = domain.RequestLogin{
			Password: "Les Miserables",
			Email:    "abc@gmail.com",
		}

		p := hash.GenerateFromPassword("Les Miserables")
		user = domain.User{
			Email:    "abc@gmail.com",
			Password: p,
		}
		user.ID = 1
	})

	Describe("API Login", func() {
		Context(" - Login", func() {
			It("should be success", func() {
				// prepare
				userRepo.EXPECT().GetUserByEmail(req.Email).Return(user, nil)
				// do
				_, err := userUc.Login(ctx, req)
				// check
				Expect(err).To(BeNil())
			})

			It("email not found", func() {
				// prepare
				userRepo.EXPECT().GetUserByEmail(req.Email).Return(domain.User{}, nil)
				// do
				_, err := userUc.Login(ctx, req)
				// check
				Expect(err).NotTo(BeNil())
			})

			It("password wrong", func() {
				// prepare
				req.Password = "12345678"
				userRepo.EXPECT().GetUserByEmail(req.Email).Return(user, nil)
				// do
				_, err := userUc.Login(ctx, req)
				// check
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("API Register", func() {
		Context(" - Register", func() {
			It("should be success", func() {
				// prepare
				userRepo.EXPECT().CreateUser(gomock.Any()).Return(&register, nil)
				// do
				err := userUc.CreateUser(ctx, register)
				// check
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("API Get user", func() {
		Context("- Get user", func() {
			It("should be return user", func() {
				// prepare
				userRepo.EXPECT().GetUser(int(user.ID)).Return(&user, nil)
				// do
				u, err := userUc.GetUser(ctx, int(user.ID))
				// check
				Expect(err).To(BeNil())
				Expect(*u).To(Equal(user))
			})
		})

		Context("- Get list user", func() {
			users = []domain.User{
				user,
				register,
			}
			It("should be success", func() {
				// prepare
				userRepo.EXPECT().GetUsers().Return(users, nil)
				// do
				u, err := userUc.GetUsers(ctx)
				// check
				Expect(err).To(BeNil())
				Expect(u).To(Equal(users))
			})
		})
	})
})
