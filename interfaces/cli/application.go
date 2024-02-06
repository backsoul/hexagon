// interfaces/cli/application.go

package cli

import (
	"database/sql"
	"fmt"

	"github.com/backsoul/hexagon/application/services"
	"github.com/backsoul/hexagon/application/use_cases"
	"github.com/backsoul/hexagon/infrastructure/repositories"
)

// Run inicializa y ejecuta la aplicación
func Run() {
	// instance repository

	dbConnection, err := ConnectToMySQL()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repositories.NewMySQLUserRepository(dbConnection)

	// Configurar el servicio de usuarios usando el repositorio
	userService := services.NewUserService(userRepository)

	// Configurar el caso de uso de crear usuario usando el servicio
	createUserCase := use_cases.NewCreateUserCase(userService)

	// Ejemplo de uso del caso de uso de crear usuario
	createUserInput := &use_cases.CreateUserInput{
		Username: "john_doe",
		Email:    "john@example.com",
	}
	output, err := createUserCase.Execute(createUserInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User created with ID:", output.UserID)
}

// setupRepositoryAndConnection configura y establece la conexión con la base de datos
func setupRepositoryAndConnection(dbType string) (repositories.UserRepository, *sql.DB, error) {
	var userRepository repositories.UserRepository
	var dbConnection *sql.DB

	switch dbType {
	case "mysql":
		dbConnection, err := ConnectToMySQL()
		if err != nil {
			return nil, nil, err
		}
		userRepository = repositories.NewMySQLUserRepository(dbConnection)
	case "postgresql":
		dbConnection, err := ConnectToPostgreSQL()
		if err != nil {
			return nil, nil, err
		}
		userRepository = repositories.NewPostgreSQLUserRepository(dbConnection)
	default:
		return nil, nil, fmt.Errorf("invalid database type specified in .env file: %s", dbType)
	}

	return userRepository, dbConnection, nil
}
