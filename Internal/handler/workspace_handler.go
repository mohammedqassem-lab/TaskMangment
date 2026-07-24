package handler

import (
	model "TaskMangment/Internal/Model"
	service "TaskMangment/Internal/Service"
	"TaskMangment/Internal/dto"

	"github.com/gin-gonic/gin"
)

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
	modelWorkspace := model.Workspace{
		Name:        workspace.Name,
		Description: workspace.Description,
		OwnerID:     int64(workspace.OwnerId),
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
