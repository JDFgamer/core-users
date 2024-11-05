package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
