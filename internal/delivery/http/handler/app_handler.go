package handler

import (
	"DRX_Test/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) RouteNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, dto.ErrorResponse{
		Message: "Route not found",
	})
}

func (h *AppHandler) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome!!"})
}
