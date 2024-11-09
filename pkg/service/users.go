package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	ctx := c.Request.Context()
	err := h.userService.GetAllUser(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "hola cami soy un texto")
}
