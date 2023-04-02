package service

import (
	"errors"
	"icl-auth/pkg/domain/repository"
	"icl-auth/pkg/model"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetByCredentials(email string, password string) (*model.User, error)
	Create(u *model.User) (*model.User, error)
	UserById(id uint) (*model.User, error)
}

type userService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) User {
	return &userService{userRepo}
}

func (s *userService) GetByCredentials(email string, password string) (*model.User, error) {
	u, err := s.userRepo.ByEmail(email)
	if u == nil {
		log.Printf("error getting user %+v", err)
		return nil, errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Printf("passwords mismatch %+v", err)
		return nil, errors.New("Invalid credentials")
	}
	return u, nil

}

func (s *userService) Create(u *model.User) (*model.User, error) {
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return nil, err
	}
	u.Password = string(pwdBytes)

	return s.userRepo.Create(u)
}

func (s *userService) UserById(id uint) (*model.User, error) {
	return s.userRepo.UserById(id)
}
