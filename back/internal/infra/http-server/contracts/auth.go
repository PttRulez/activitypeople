package contracts

import "github.com/pttrulez/activitypeople/internal/domain"

type RegisterUserRequest struct {
	LoginUserRequest
	Name            string      `json:"name" validate:"required"`
	ConfirmPassword string      `json:"confirmPassword" validate:"required,eqfield=Password"`
	Role            domain.Role `json:"role"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
