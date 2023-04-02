package repository

import (
	"icl-auth/pkg/domain/repository"
	"icl-auth/pkg/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.User {
	return &userRepository{db}
}

func (r *userRepository) ByEmail(email string) (*model.User, error) {
	var user model.User
	record := r.db.Where("email = ?", email).First(&user)
	if record.Error != nil {
		return nil, record.Error
	}

	return &user, nil
}

func (r *userRepository) Create(u *model.User) (*model.User, error) {
	record := r.db.Create(&u)
	if record.Error != nil {
		return nil, record.Error
	}
	return u, nil
}

func (r *userRepository) UserById(id uint) (*model.User, error) {
	var user model.User
	record := r.db.Where("id = ?", id).First(&user)
	if record.Error != nil {
		return nil, record.Error
	}

	return &user, nil
}
