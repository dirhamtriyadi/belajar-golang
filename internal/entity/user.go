package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        int        `gorm:"primary_key"`                               // ID int
	Username  string     `gorm:"unique;not null" validate:"required"`       // Username string
	Email     string     `gorm:"unique;not null" validate:"required,email"` // Email string
	Password  string     `gorm:"not null" validate:"required"`              // Password string
	CreatedAt *time.Time // CreatedAt date
	UpdatedAt *time.Time // UpdatedAt date
}

type UserResponse struct {
	ID        int        `json:"id"`         // ID int
	Username  string     `json:"username"`   // Username string
	Email     string     `json:"email"`      // Email string
	CreatedAt *time.Time `json:"created_at"` // CreatedAt date
	UpdatedAt *time.Time `json:"updated_at"` // UpdatedAt date
}

// Validate will validate the user struct
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
