package handler

import (
	model "TaskMangment/Internal/Model"
	service "TaskMangment/Internal/Service"
	"TaskMangment/Internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
func (h *UserHandler) Register(c *gin.Context) {
	var requst dto.UserDto
	if err := c.ShouldBindJSON(&requst); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "BadRequest",
			"error":   err.Error(),
		})
		return
	}
	modelUser := model.User{
		Name:         requst.Name,
		Email:        requst.Email,
		Hashpassword: requst.Password,
	}
	err := h.userService.Register(c.Request.Context(), modelUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "InternalServerError",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
func (h *UserHandler) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	modelUser := model.User{
		Email:        loginDto.Email,
		Hashpassword: loginDto.Password,
	}
	token, err := h.userService.Login(c.Request.Context(), modelUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
