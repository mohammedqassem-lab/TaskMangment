package handler

import (
	model "TaskMangment/Internal/Model"
	service "TaskMangment/Internal/Service"
	"TaskMangment/Internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService service.ProjectService
}

func NewProjectHandler(projectService service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// CreateProject godoc
// @Summary Create a new project in the workspace
// @Description Create a new project in the workspace
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Param request body dto.AddProjectRequest true "AddProjectRequest Request"
// @Router /project/{id}/create [post]
func (h *ProjectHandler) Create(c *gin.Context) {
	var requst dto.AddProjectRequest
	workspaceID := c.Param("id")
	workspaceIDInt, err := strconv.ParseInt(workspaceID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid workspace ID",
		})
		return
	}
	if err := c.ShouldBindJSON(&requst); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "BadRequest",
			"error":   err.Error(),
		})
		return
	}
	UserId, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"massage": "you are not auth",
		})
		return
	}
	UserIdInt, err := strconv.ParseInt(UserId.(string), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"massage": "you are not auth",
		})
		return
	}
	model := model.Project{
		Name:        requst.Name,
		Description: requst.Description,
		WorkspaceId: workspaceIDInt,
		CreatedBy:   UserIdInt,
	}
	err = h.projectService.Create(c, &model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "an error occorded",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"massage": "Project created",
	})
}

// GetProjectById godoc
// @Summary Get a project by ID
// @Description Get a project by its ID
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Param projectid path int true "Project ID"
// @Router /project/{id} [get]
func (h *ProjectHandler) GetById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})
		return
	}
	projectIdstr := c.Param("projectid")
	projectId, err := strconv.ParseInt(projectIdstr, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "invalid project id",
			"error":   err.Error(),
		})
		return
	}
	workspaces, err := h.projectService.GetById(c.Request.Context(), projectId, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"workspaces": workspaces})
}

// GetProjects godoc
// @Summary get projects in the workspace
// @Description get projects in the workspace
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Router /Project/{id}/Get [get]
func (h *ProjectHandler) Get(c *gin.Context) {
	workspaceId := c.Param("id")
	id, err := strconv.ParseInt(workspaceId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})
		return
	}
	projects, err := h.projectService.Get(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}
	if projects == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

// UpdateProject godoc
// @Summary update the project in the workspace
// @Description update the project in the workspace
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Param request body dto.UpdateProject true "UpdateProject Request"
// @Router /Project/{id}/Update [put]
func (h *ProjectHandler) Update(c *gin.Context) {
	workspaceId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid workspace id",
			"error":   err.Error(),
		})
		return
	}
	var model dto.UpdateProject
	err = c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid data",
			"error":   err.Error(),
		})
		return
	}
	model.WorkspaceId = workspaceId
	err = h.projectService.Update(c, &model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "an error ocorded",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "the Project Updeted",
	})
}

// DeleteProject godoc
// @Summary Delete the projects in the workspace
// @Description Delete the projects in the workspace
// @Tags Project
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Param projectid path int true "project ID"
// @Router /Project/{id}/Delete/{projectid} [delete]
func (h *ProjectHandler) Delete(c *gin.Context) {
	workspaceId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid workspace id",
			"error":   err.Error(),
		})
		return
	}
	idstr := c.Param("projectid")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massege": "invalid id",
		})
		return
	}
	err = h.projectService.Delete(c, id, workspaceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "project deleted",
	})
}
