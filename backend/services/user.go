package services

import (
	"github.com/nambuitechx/go-social/models"
	"github.com/nambuitechx/go-social/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{ UserRepository: userRepository }
}

func (s *UserService) Health() string {
	return "User service is available"
}

func (s *UserService) GetAllUsers(limit *int, offset *int) ([]models.UserModel, error) {
	users, err := s.UserRepository.SelectUsers(limit, offset)
	return users, err
}

func (s *UserService) GetUserById(id *string) (*models.UserModel, error) {
	user, err := s.UserRepository.SelectUserById(id)
	return user, err
}

func (s *UserService) CreateUser(payload *models.CreateUserPayload) (*models.UserModel, error) {
	user, err := s.UserRepository.InsertUser(payload)
	return user, err
}

func (s *UserService) DeleteUserById(id *string) error {
	err := s.UserRepository.DeleteUserById(id)
	return err
}
