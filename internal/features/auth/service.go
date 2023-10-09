package auth

import (
	"boilerplate/internal/models"
	"fmt"
)

// TODO: Add verification feature.
type authServiceImpl struct {
	authRepository AuthRepository
}

func NewAuthService(
	authRepository AuthRepository,
) AuthService {
	return &authServiceImpl{
		authRepository: authRepository,
	}
}

func (s *authServiceImpl) Validate(token string) (string, *models.AppError) {
	firebaseAuthId, err := s.authRepository.Validate(token)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("firebase_%v", firebaseAuthId), nil
}
