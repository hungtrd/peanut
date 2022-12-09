package usecase

type UserUsecase interface {
	GetUsers()
	GetUser(id int)
	CreateUser()
}

type userUsecase struct {
	UserRepo repository.UserRepo
}
