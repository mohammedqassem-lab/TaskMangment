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

// Workspace Member Routes
func InviteMemberRoutes(r *gin.RouterGroup, workspaceMemberHandler *handler.WorkspaceMemberHandler) {
	r.POST("/workspace/:id/invite", workspaceMemberHandler.InviteMember)
}
func GetWorkspaceMembersRoutes(r *gin.RouterGroup, workspaceMemberHandler *handler.WorkspaceMemberHandler) {
	r.GET("/workspace/:id/members", workspaceMemberHandler.GetWorkspaceMembers)
}
func UpdateMemberRoleRoutes(r *gin.RouterGroup, workspaceMemberHandler *handler.WorkspaceMemberHandler) {
	r.PUT("/workspace/:id/UpdateMemberRole", workspaceMemberHandler.UpdateMemberRole)
}
func DeleteMemberRoutes(r *gin.RouterGroup, workspaceMemberHandler *handler.WorkspaceMemberHandler) {
	r.DELETE("/workspace/:id/DeleteMember/:user_id", workspaceMemberHandler.DeleteMember)
}
