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
	r.PUT("/workspace/:id/Update", workspaceHandler.UpdateWorkspace)
}
func DeleteWorkspaceRoutes(r *gin.RouterGroup, workspaceHandler *handler.WorkspaceHandler) {
	r.DELETE("/workspace/:id/Delete", workspaceHandler.DeleteWorkspace)
}
