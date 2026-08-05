package handler

import (
	service "TaskMangment/Internal/Service"
	"TaskMangment/Internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService service.TaskService
}

func NewTaskHandler(TaskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: TaskService,
	}
}

// CreateTask godoc
// @Summary Create a new task in the workspace
// @Description Create a new task in the workspace
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Param request body dto.AddTask true "AddTask Request"
// @Router /Task/{id}/Create [post]
func (h *TaskHandler) Create(c *gin.Context) {
	var requst dto.AddTask
	workspaceID := c.Param("id")
	workspaceIDInt, err := strconv.ParseInt(workspaceID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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
	requst.WorkSpaceId = workspaceIDInt
	requst.CreateUserId = UserIdInt
	err = h.TaskService.Create(c, &requst)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "task Created",
	})
}

// EditTask godoc
// @Summary Edit the task in the workspace
// @Description Edit the task in the workspace
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Param request body dto.EditTask true "EditTask Request"
// @Router /Task/{id}/Edit [put]
func (h *TaskHandler) Edit(c *gin.Context) {
	var requst dto.EditTask
	workspaceId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid workspace id",
			"error":   err.Error(),
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
	requst.UserId = UserIdInt
	requst.WorkSpaceId = workspaceId
	err = h.TaskService.Edit(c, &requst)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "Task Updeted",
	})
}

// DeleteTask godoc
// @Summary Delete the task in the workspace
// @Description Delete the task in the workspace
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Param task_id path int true "Task ID"
// @Router /Task/{id}/Delete/{task_id} [delete]
func (h *TaskHandler) Delete(c *gin.Context) {
	idstr := c.Param("task_id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "invalid task id",
		})
		return
	}
	err = h.TaskService.Delete(c, id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "task Deleted",
	})
}

// GetTasks godoc
// @Summary Get all the task in the workspace
// @Description Get all the task in the workspace
// @Tags Task
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "WorkSpace ID"
// @Param Status query string false "Task Status"
// @Param ProjectId query int false "ProjectId Number"
// @Param Priorty query string false "Task Priorty"
// @Param AssigneeId query int false "AssigneeId"
// @Param Serch query string false "Serch"
// @Param SortBy query string false "SortBy"
// @Param Order query string false "Order"
// @Param Limit query int false "Limit Number"
// @Param Offset query int false "Offset Number"
// @Router /Task/{id}/GetAll [get]
func (h *TaskHandler) GetAll(c *gin.Context) {
	var filter dto.TaskFilter
	workspaceId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid workspace id",
			"error":   err.Error(),
		})
		return
	}
	filter.WorkSpaceId = workspaceId
	err = c.ShouldBindQuery(&filter)
	if err != nil {
		c.JSON(400, nil)
	}
	Tasks, err := h.TaskService.GetAll(c, filter)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, Tasks)
}
