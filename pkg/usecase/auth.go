package usecase

import (
	"icl-auth/pkg/domain/service"
	"icl-auth/pkg/model"
)

type Auth interface {
	Login(l *LogingDTO) (*model.User, error)
	Register(user *model.User) (*model.User, error)
	UserById(id uint) (*model.User, error)
}

type authUsecase struct {
	userService service.User
}

func NewAuth(userService service.User) Auth {
	return &authUsecase{userService}
}

type LogingDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *authUsecase) Login(l *LogingDTO) (*model.User, error) {
	return u.userService.GetByCredentials(l.Email, l.Password)
}

type RegisterDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *authUsecase) Register(user *model.User) (*model.User, error) {
	return u.userService.Create(user)
}

func (u *authUsecase) UserById(id uint) (*model.User, error) {
	return u.userService.UserById(id)
}
