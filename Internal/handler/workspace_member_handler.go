package handler

import (
	service "TaskMangment/Internal/Service"
	"TaskMangment/Internal/dto"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkspaceMemberHandler struct {
	workspaceMemberService service.WorkspaceMemberService
}

func NewWorkspaceMemberHandler(workspaceMemberService service.WorkspaceMemberService) *WorkspaceMemberHandler {
	return &WorkspaceMemberHandler{
		workspaceMemberService: workspaceMemberService,
	}
}
func (h *WorkspaceMemberHandler) InviteMember(c *gin.Context) {
	var memper dto.AddMemberDto
	workspaceID := c.Param("id")
	workspaceIDInt, err := strconv.ParseInt(workspaceID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid workspace ID",
		})
		return
	}
	if err := c.ShouldBindJSON(&memper); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	if workspaceID == "" {
		c.JSON(400, gin.H{
			"message": "Workspace ID is required",
		})
		return
	}
	err = h.workspaceMemberService.InviteMember(c.Request.Context(), workspaceIDInt, memper.UserID, memper.Role)
	if err != nil {
		c.JSON(404, gin.H{
			"error":   err.Error(),
			"message": "Workspace not found",
		})
		return
	}
	c.JSON(200, gin.H{"message": "Member invited successfully"})
}
func (h *WorkspaceMemberHandler) GetWorkspaceMembers(c *gin.Context) {
	workspaceID := c.Param("id")
	workspaceIDInt, err := strconv.ParseInt(workspaceID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid workspace ID",
		})
		return
	}
	members, err := h.workspaceMemberService.GetWorkspaceMembers(c.Request.Context(), workspaceIDInt)
	if err != nil {
		c.JSON(404, gin.H{
			"error":   err.Error(),
			"message": "Workspace not found",
		})
		return
	}
	c.JSON(200, gin.H{"members": members})
}
func (h *WorkspaceMemberHandler) UpdateMemberRole(c *gin.Context) {
	fmt.Println("UpdateMemberRole called")
	var req dto.UpdateMemberRoleDto
	var workspaceIDStr = c.Param("id")
	workspaceID, err := strconv.ParseInt(workspaceIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid workspace ID",
		})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	err = h.workspaceMemberService.UpdateMemberRole(c.Request.Context(), workspaceID, req.UserID, req.Role)
	if err != nil {
		c.JSON(404, gin.H{
			"error":   err.Error(),
			"message": "Workspace not found",
		})
		return
	}
	c.JSON(200, gin.H{"message": "Member role updated successfully"})
}
func (h *WorkspaceMemberHandler) DeleteMember(c *gin.Context) {
	workspaceIDStr := c.Param("id")
	workspaceID, err := strconv.ParseInt(workspaceIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid workspace ID",
		})
		return
	}
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	err = h.workspaceMemberService.DeleteMember(c.Request.Context(), workspaceID, userID)
	if err != nil {
		c.JSON(404, gin.H{
			"error":   err.Error(),
			"message": "Workspace not found",
		})
		return
	}
	c.JSON(200, gin.H{"message": "Member deleted successfully"})
}
