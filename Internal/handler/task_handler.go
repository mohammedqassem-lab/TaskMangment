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
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "InternalServerError",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "task Created",
	})
}
func (h *TaskHandler) Edit(c *gin.Context) {
	var requst dto.EditTask
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
	err = h.TaskService.Edit(c, &requst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "InternalServerError",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "Task Updeted",
	})
}
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"massage": "task Deleted",
	})
}
func (h *TaskHandler) GetAll(c *gin.Context) {
	var filter dto.TaskFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		c.JSON(400, nil)
	}
	Tasks, err := h.TaskService.GetAll(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Tasks)
}
