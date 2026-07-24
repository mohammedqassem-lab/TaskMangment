package route

import (
	handler "TaskMangment/Internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, userHandler *handler.UserHandler) {
	r.POST("/register", userHandler.Register)
}
func LoginUserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	r.POST("/login", userHandler.Login)
}
func CreateWorkspaceRoutes(r *gin.RouterGroup, workspaceHandler *handler.WorkspaceHandler) {
	r.POST("/workspace/create", workspaceHandler.CreateWorkspace)
}
