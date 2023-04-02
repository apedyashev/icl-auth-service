package repository

import (
	"icl-auth/pkg/model"
)

type User interface {
	ByEmail(email string) (*model.User, error)
	Create(u *model.User) (*model.User, error)
	UserById(id uint) (*model.User, error)
}
