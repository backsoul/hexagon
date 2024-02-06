// interfaces/cli/main.go

package main

import (
	"fmt"

	"github.com/backsoul/hexagon/application/services"
	"github.com/backsoul/hexagon/application/use_cases"
	"github.com/backsoul/hexagon/infrastructure/repositories"
)

func main() {
	// Crear el repositorio de usuarios
	userRepository := repositories.NewUserRepository()

	// Crear el servicio de usuarios usando el repositorio
	userService := services.NewUserService(userRepository)

	// Crear el caso de uso de crear usuario usando el servicio de usuarios
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
