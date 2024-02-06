package cli

import (
	"database/sql"
	"os"
)

// ConnectToMySQL configures and establishes a connection to the MySQL database
func ConnectToMySQL() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := dbUser + ":" + dbPassword + "@tcp(localhost:3306)/" + dbName

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToPostgreSQL() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := "user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
