package user

import (
	"boilerplate/internal/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

type UserGorm struct {
	ID        int64     `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	AuthID    string    `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

func (UserGorm) TableName() string {
	return "users"
}

func (r *userRepositoryImpl) makeUserGormFrom(user *models.User) *UserGorm {
	return &UserGorm{
		ID:     user.ID,
		Name:   user.Name,
		AuthID: user.AuthID,
	}
}

func (r *userRepositoryImpl) makeUserFrom(userGorm *UserGorm) *models.User {
	return &models.User{
		ID:        userGorm.ID,
		Name:      userGorm.Name,
		AuthID:    userGorm.AuthID,
		CreatedAt: userGorm.CreatedAt,
		UpdatedAt: userGorm.UpdatedAt,
	}
}

func (r *userRepositoryImpl) GetUserById(id int64) (*models.User, *models.AppError) {
	var userGorm UserGorm
	if err := r.db.First(&userGorm, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.NewNotFoundError("User")
		}

		return nil, models.NewInternalServerError()
	}

	return r.makeUserFrom(&userGorm), nil
}

func (r *userRepositoryImpl) GetUserByAuthId(authId string) (*models.User, *models.AppError) {
	var userGorm UserGorm
	if err := r.db.First(&userGorm, "auth_id = ?", authId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.NewNotFoundError("User")
		}

		return nil, models.NewInternalServerError()
	}

	return r.makeUserFrom(&userGorm), nil
}

func (r *userRepositoryImpl) CreateUser(user *models.User) (*models.User, *models.AppError) {
	userGorm := r.makeUserGormFrom(user)

	if err := r.db.Create(&userGorm).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, models.NewConflictError("User")
		}

		return nil, models.NewInternalServerError()
	}

	return r.makeUserFrom(userGorm), nil
}

func (r *userRepositoryImpl) UpdateUser(user *models.User) (*models.User, *models.AppError) {
	userGorm := r.makeUserGormFrom(user)

	if err := r.db.Model(&userGorm).Updates(userGorm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.NewNotFoundError("User")
		}

		return nil, models.NewInternalServerError()
	}

	return r.makeUserFrom(userGorm), nil
}

func (r *userRepositoryImpl) DeleteUserById(id int64) *models.AppError {
	if err := r.db.Delete(&UserGorm{}, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.NewNotFoundError("User")
		}

		return models.NewInternalServerError()
	}

	return nil
}
