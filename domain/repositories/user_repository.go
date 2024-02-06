// domain/repositories/user_repository.go

package repositories

import "github.com/backsoul/hexagon/domain/entities"

// UserRepository define las operaciones sobre usuarios
type UserRepository interface {
	Store(user *entities.User) error
	FindByID(id string) (*entities.User, error)
}
