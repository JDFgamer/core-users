package main

import (
	"log"

	"core-users/pkg/configs"
	"core-users/pkg/dependencies"
	handler "core-users/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuraciones
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	external := dependencies.InitExternal(*cfg)

	// Crear repository
	repository := dependencies.InitRepository(external)

	// Crear servicios
	service := dependencies.InitService(repository)

	// Crear Handlers
	userHandler := handler.NewHandler(service)

	// Crear Gin y las rutas
	r := gin.Default()

	// Configurar rutas
	handler.SetupRoutes(r, userHandler)

	// Definir el puerto (puede ser tomado de una variable de entorno o de config)
	port := ":8080"
	log.Printf("Server is running on http://localhost%s", port)

	// Iniciar el servidor
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
