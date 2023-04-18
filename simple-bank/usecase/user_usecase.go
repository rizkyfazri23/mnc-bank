package usecase

import (
	"github.com/rizkyfazri23/mnc/model/entity"
	"github.com/rizkyfazri23/mnc/repository"
)

type UserUsecase interface {
	GetOne(id int) (entity.User, error)
	Add(newUser *entity.User) (entity.User, error)
	Edit(user *entity.User, userId int) (entity.User, error)
	LoginCheck(username string, password string) (string, error)
	Logout(userId int) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func (u *userUsecase) GetOne(id int) (entity.User, error) {
	return u.userRepo.FindOne(id)
}

func (u *userUsecase) Add(newUser *entity.User) (entity.User, error) {
	return u.userRepo.Create(newUser)
}

func (u *userUsecase) Edit(user *entity.User, userId int) (entity.User, error) {
	return u.userRepo.Update(user, userId)
}

func (u *userUsecase) LoginCheck(username string, password string) (string, error) {
	return u.userRepo.LoginCheck(username, password)
}

func (u *userUsecase) Logout(userId int) (string, error) {
	return u.userRepo.Logout(userId)
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
