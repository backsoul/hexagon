// infrastructure/adapters/http_handler.go

package adapters

import (
	"encoding/json"
	"net/http"

	"github.com/backsoul/hexagon/application/use_cases"
)

// HTTPHandler representa un manejador HTTP para interactuar con la aplicación de gestión de usuarios
type HTTPHandler struct {
	CreateUserCase *use_cases.CreateUserCase
}

// NewHTTPHandler crea una nueva instancia de HTTPHandler
func NewHTTPHandler(createUserCase *use_cases.CreateUserCase) *HTTPHandler {
	return &HTTPHandler{
		CreateUserCase: createUserCase,
	}
}

// CreateUserHandler maneja las solicitudes HTTP para crear un nuevo usuario
func (h *HTTPHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud JSON
	var input use_cases.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Error al decodificar la solicitud JSON", http.StatusBadRequest)
		return
	}

	// Ejecutar el caso de uso de crear usuario
	output, err := h.CreateUserCase.Execute(&input)
	if err != nil {
		http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
		return
	}

	// Escribir la respuesta JSON
	response, err := json.Marshal(output)
	if err != nil {
		http.Error(w, "Error al serializar la respuesta JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
