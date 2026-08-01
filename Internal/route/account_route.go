package route

import (
	handler "TaskMangment/Internal/handler"

	"github.com/gin-gonic/gin"
)

// Account Routes
func RegisterUserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	r.POST("/register", userHandler.Register)
}
func LoginUserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	r.POST("/login", userHandler.Login)
}
func RefreshTokenRoute(r *gin.Engine, userHandler *handler.UserHandler) {
	r.POST("/refreshToken", userHandler.RefreshToken)
}
