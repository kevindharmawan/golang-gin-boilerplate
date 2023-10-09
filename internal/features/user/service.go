package user

import "boilerplate/internal/models"

type userServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(
	userRepository UserRepository,
) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (s *userServiceImpl) GetUserById(id int64) (*models.User, *models.AppError) {
	return s.userRepository.GetUserById(id)
}

func (s *userServiceImpl) GetUserByAuthId(authId string) (*models.User, *models.AppError) {
	return s.userRepository.GetUserByAuthId(authId)
}

func (s *userServiceImpl) GetUserIdByAuthId(authId string) (int64, *models.AppError) {

	user, err := s.userRepository.GetUserByAuthId(authId)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (s *userServiceImpl) CreateUser(user *models.User) (*models.User, *models.AppError) {
	return s.userRepository.CreateUser(user)
}

func (s *userServiceImpl) UpdateUser(user *models.User) (*models.User, *models.AppError) {
	return s.userRepository.UpdateUser(user)
}

func (s *userServiceImpl) DeleteUserById(id int64) *models.AppError {
	return s.userRepository.DeleteUserById(id)
}
