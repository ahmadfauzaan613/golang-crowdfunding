package handler

import (
	"crowdfunding/dto"
	"crowdfunding/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.ServiceUser
}

func HandlerUser(userService services.ServiceUser) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var input dto.UserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	user, err := h.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
