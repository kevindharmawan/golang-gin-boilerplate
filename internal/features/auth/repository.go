package auth

import (
	"boilerplate/internal/models"
	"context"

	"firebase.google.com/go/auth"
)

type authRepositoryImpl struct {
	firebaseAuthClient *auth.Client
}

func NewAuthRepository(firebaseAuthClient *auth.Client) AuthRepository {
	return &authRepositoryImpl{
		firebaseAuthClient: firebaseAuthClient,
	}
}

func (r *authRepositoryImpl) Validate(token string) (string, *models.AppError) {
	firebaseUser, err := r.firebaseAuthClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		return "", models.NewUnauthorizedError(true)
	}

	return firebaseUser.UID, nil
}
