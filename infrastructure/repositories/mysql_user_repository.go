// infrastructure/repositories/mysql_user_repository.go

package repositories

import (
	"database/sql"
	"errors"

	"github.com/backsoul/hexagon/domain/entities"
)

// MySQLUserRepository implementa UserRepository utilizando MySQL
type MySQLUserRepository struct {
	db *sql.DB
}

// NewMySQLUserRepository crea una nueva instancia de MySQLUserRepository
func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{
		db: db,
	}
}

// Store almacena un usuario en la base de datos MySQL
func (r *MySQLUserRepository) Store(user *entities.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, username, email) VALUES (?, ?, ?)", user.ID, user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil
}

// FindByID encuentra un usuario por su ID en la base de datos MySQL
func (r *MySQLUserRepository) FindByID(id string) (*entities.User, error) {
	var user entities.User
	err := r.db.QueryRow("SELECT id, username, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
