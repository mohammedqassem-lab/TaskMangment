package route

import (
	handler "TaskMangment/Internal/handler"

	"github.com/gin-gonic/gin"
)

// Workspace Routes
func CreateWorkspaceRoutes(r *gin.RouterGroup, workspaceHandler *handler.WorkspaceHandler) {
	r.POST("/workspace/create", workspaceHandler.CreateWorkspace)
}
func GetAllWorkspaceRoutes(r *gin.RouterGroup, workspaceHandler *handler.WorkspaceHandler) {
	r.GET("/workspace", workspaceHandler.GetAllWorkspace)
}
func UpdateWorkspaceRoutes(r *gin.RouterGroup, workspaceHandler *handler.WorkspaceHandler) {
	r.PUT("/workspace/Update/:id", workspaceHandler.UpdateWorkspace)
}
func DeleteWorkspaceRoutes(r *gin.RouterGroup, workspaceHandler *handler.WorkspaceHandler) {
	r.DELETE("/workspace/Delete/:id", workspaceHandler.DeleteWorkspace)
}
