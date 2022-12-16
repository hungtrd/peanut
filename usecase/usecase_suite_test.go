package usecase_test

import (
	"testing"

	"peanut/repository/mock"
	"peanut/usecase"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecase Suite")
}

var _ = BeforeSuite(func() {
	sqlDB, smock, _ := sqlmock.New()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	Expect(err).To(BeNil())
	Expect(smock).NotTo(BeNil())
	Expect(db).NotTo(BeNil())

	ctrl := gomock.NewController(GinkgoT())
	defer ctrl.Finish()
    userRepo := mock.NewMockUserRepo(ctrl)
	userUc := usecase.NewUserUsecase(userRepo)
})
