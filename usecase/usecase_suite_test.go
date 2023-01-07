package usecase_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"peanut/config"
	"peanut/repository/mock"
	"peanut/usecase"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctx context.Context
var db *gorm.DB
var userRepo *mock.MockUserRepo
var todoRepo *mock.MockTodoRepo
var contentRepo *mock.MockContentRepo
var userUc usecase.UserUsecase
var todoUc usecase.TodoUsecase
var contentUc usecase.ContentUsecase

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecase Suite")
}

var _ = BeforeSuite(func() {
	config.Setup()

	sqlDB, smock, _ := sqlmock.New()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	Expect(err).To(BeNil())
	Expect(smock).NotTo(BeNil())
	Expect(db).NotTo(BeNil())

	ctrl := gomock.NewController(GinkgoT())
	defer ctrl.Finish()
	userRepo = mock.NewMockUserRepo(ctrl)
	userUc = usecase.NewUserUsecase(userRepo)

	todoRepo = mock.NewMockTodoRepo(ctrl)
	todoUc = usecase.NewTodoUsecase(todoRepo)

	contentRepo = mock.NewMockContentRepo(ctrl)
	contentUc = usecase.NewContentUsecase(contentRepo)
})
