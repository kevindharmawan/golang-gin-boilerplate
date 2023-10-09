package auth

import "boilerplate/internal/models"

type AuthRepository interface {
	Validate(token string) (string, *models.AppError)
}

type AuthService interface {
	Validate(token string) (string, *models.AppError)
}
