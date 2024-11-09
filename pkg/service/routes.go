// routes.go
package handler

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas de la aplicación
func (h *Handler) SetupRoutes(r *gin.Engine) {
	// Rutas relacionadas con los usuarios
	r.GET("/users", h.GetAllUsers) // Pasa las funciones sin paréntesis
}
