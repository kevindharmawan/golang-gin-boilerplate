package user

import "boilerplate/internal/models"

type UserRepository interface {
	GetUserById(id int64) (*models.User, *models.AppError)
	GetUserByAuthId(authId string) (*models.User, *models.AppError)
	CreateUser(user *models.User) (*models.User, *models.AppError)
	UpdateUser(user *models.User) (*models.User, *models.AppError)
	DeleteUserById(id int64) *models.AppError
}

type UserCacheRepository interface {
	GetUserIdByAuthId(authId string) (int64, *models.AppError)
	SetUserIdByAuthId(authId string, userId int64) *models.AppError
}

type UserService interface {
	GetUserById(id int64) (*models.User, *models.AppError)
	GetUserByAuthId(authId string) (*models.User, *models.AppError)
	GetUserIdByAuthId(authId string) (int64, *models.AppError)
	CreateUser(user *models.User) (*models.User, *models.AppError)
	UpdateUser(user *models.User) (*models.User, *models.AppError)
	DeleteUserById(id int64) *models.AppError
}
