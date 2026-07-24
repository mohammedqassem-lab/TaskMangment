package handler

import (
	model "TaskMangment/Internal/Model"
	service "TaskMangment/Internal/Service"
	"TaskMangment/Internal/dto"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(c *gin.Context) (int64, error) {
	user, exists := c.Get("user")
	if !exists {
		return 0, fmt.Errorf("user not found in context")
	}
	userStr, ok := user.(string)
	if !ok {
		return 0, fmt.Errorf("user ID is not a string")
	}
	userId, err := strconv.ParseInt(userStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID")
	}
	return userId, nil
}

type WorkspaceHandler struct {
	workspaceService service.WorkspaceService
}

func NewWorkspaceHandler(workspaceService service.WorkspaceService) *WorkspaceHandler {
	return &WorkspaceHandler{
		workspaceService: workspaceService,
	}
}
func (h *WorkspaceHandler) CreateWorkspace(c *gin.Context) {
	var workspace dto.CreateWorkspaceRequest
	if err := c.ShouldBindJSON(&workspace); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	fmt.Println("User from context:", user)
	userStr, ok := user.(string)
	if !ok {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	userId, err := strconv.ParseInt(userStr, 10, 46)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	modelWorkspace := model.Workspace{
		Name:        workspace.Name,
		Description: workspace.Description,
		OwnerID:     userId,
	}
	if err := h.workspaceService.Create(c.Request.Context(), &modelWorkspace); err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Workspace created successfully",
	})
}
func (h *WorkspaceHandler) InviteMember(c *gin.Context) {
	var memper dto.MemperDto
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
	err = h.workspaceService.InviteMember(c.Request.Context(), workspaceIDInt, memper.UserID, memper.Role)
	if err != nil {
		c.JSON(404, gin.H{
			"error":   err.Error(),
			"message": "Workspace not found",
		})
		return
	}
	c.JSON(200, gin.H{"message": "Member invited successfully"})
}
