// routes.go
package handler

import (
	"core-users/pkg/dependencies"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas de la aplicación
func SetupRoutes(r *gin.Engine, userService dependencies.Service) {
	handler := NewHandler(userService)

	// Rutas relacionadas con los usuarios
	r.GET("/users", handler.GetAllUsers) // Pasa las funciones sin paréntesis
}
