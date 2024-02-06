// interfaces/cli/main.go

package main

import (
	"log"

	"github.com/backsoul/hexagon/interfaces/cli"
	"github.com/joho/godotenv"
)

// LoadEnv carga las variables de entorno desde un archivo .env
func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func main() {
	// Cargar variables de entorno
	if err := LoadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cli.Run()
}
