package handler

import (
	"net/http"
	"user/model"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.UserUsecase.CreateUser(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}
