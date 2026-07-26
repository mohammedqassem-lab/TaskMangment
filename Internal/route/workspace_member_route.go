package route

import (
	handler "TaskMangment/Internal/handler"

	"github.com/gin-gonic/gin"
)

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
