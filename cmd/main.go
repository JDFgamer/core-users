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
	domains := dependencies.InitService(repository)

	// Crear Gin y las rutas
	r := gin.Default()

	handler := handler.NewHandler(domains.UserService) // Inicializa el handler con el servicio
	handler.SetupRoutes(r)                             // Configura las rutas usando el handler

	// Definir el puerto (puede ser tomado de una variable de entorno o de config)
	port := ":8080"
	log.Printf("Server is running on http://localhost%s", port)

	// Iniciar el servidor
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
