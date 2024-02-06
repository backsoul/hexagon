package services

import (
	"github.com/backsoul/hexagon/domain/entities"
	"github.com/backsoul/hexagon/infrastructure/repositories"
)

// UserService proporciona funcionalidades más complejas relacionadas con la gestión de usuarios
type UserService struct {
	UserRepository repositories.UserRepository
}

// NewUserService crea una nueva instancia de UserService
func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

// CreateUser crea un nuevo usuario
func (s *UserService) CreateUser(username, email string) (*entities.User, error) {
	user := &entities.User{
		Username: username,
		Email:    email,
	}
	err := s.UserRepository.Store(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
