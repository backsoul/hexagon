// infrastructure/repositories/postgresql_user_repository.go

package repositories

import (
	"database/sql"
	"errors"

	"github.com/backsoul/hexagon/domain/entities"
)

type PostgreSQLUserRepository struct {
	db *sql.DB
}

func NewPostgreSQLUserRepository(db *sql.DB) *PostgreSQLUserRepository {
	return &PostgreSQLUserRepository{
		db: db,
	}
}

func (r *PostgreSQLUserRepository) Store(user *entities.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, username, email) VALUES ($1, $2, $3)", user.ID, user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLUserRepository) FindByID(id string) (*entities.User, error) {
	var user entities.User
	err := r.db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
