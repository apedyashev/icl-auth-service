package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Username  string    `json:"username" gorm:"unique" validate:"required"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email"`
	Password  string    `json:"password"  validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
