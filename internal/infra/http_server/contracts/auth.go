package contracts

import "github.com/pttrulez/activitypeople/internal/domain"

type RegisterUserRequest struct {
	Email           string      `json:"email" validate:"required,email"`
	Name            string      `json:"name" validate:"required"`
	Password        string      `json:"password" validate:"required"`
	ConfirmPassword string      `json:"confirmPassword" validate:"required,eqfield=Password"`
	Role            domain.Role `json:"role"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
