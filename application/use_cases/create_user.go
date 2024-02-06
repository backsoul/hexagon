package use_cases

import (
	"github.com/backsoul/hexagon/application/services"
)

// CreateUserInput representa la entrada para el caso de uso de crear usuario
type CreateUserInput struct {
	Username string
	Email    string
}

// CreateUserOutput representa la salida del caso de uso de crear usuario
type CreateUserOutput struct {
	UserID string
}

// CreateUserCase implementa el caso de uso de crear usuario
type CreateUserCase struct {
	UserService *services.UserService
}

// NewCreateUserCase crea una nueva instancia de CreateUserCase
func NewCreateUserCase(userService *services.UserService) *CreateUserCase {
	return &CreateUserCase{
		UserService: userService,
	}
}

// Execute ejecuta el caso de uso de crear usuario
func (uc *CreateUserCase) Execute(input *CreateUserInput) (*CreateUserOutput, error) {
	// Crear usuario usando el servicio
	user, err := uc.UserService.CreateUser(input.Username, input.Email)
	if err != nil {
		return nil, err
	}

	// Retornar la salida del caso de uso
	return &CreateUserOutput{UserID: user.ID}, nil
}
