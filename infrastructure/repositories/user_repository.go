// infrastructure/repositories/user_repository.go

package repositories

import (
	"errors"

	"github.com/backsoul/hexagon/domain/entities"
)

// UserRepositoryImpl implementa UserRepository usando un mapa en memoria
type UserRepositoryImpl struct {
	users map[string]*entities.User
}

// NewUserRepository crea una nueva instancia de UserRepositoryImpl
func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: make(map[string]*entities.User),
	}
}

// Store almacena un usuario en el repositorio
func (r *UserRepositoryImpl) Store(user *entities.User) error {
	r.users[user.ID] = user
	return nil
}

// FindByID encuentra un usuario por su ID
func (r *UserRepositoryImpl) FindByID(id string) (*entities.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
