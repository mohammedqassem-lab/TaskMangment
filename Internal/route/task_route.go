package route

import (
	"TaskMangment/Internal/handler"

	"github.com/gin-gonic/gin"
)

func CreateTaskRoute(r *gin.RouterGroup, taskHandler *handler.TaskHandler) {
	r.POST("/Task/:id/Create", taskHandler.Create)
}
func EditTaskRoute(r *gin.RouterGroup, taskHandler *handler.TaskHandler) {
	r.PUT("Task/:id/Edit", taskHandler.Edit)
}
func DeleteTaskRoute(r *gin.RouterGroup, taskHandler *handler.TaskHandler) {
	r.DELETE("Task/:id/Delete/:task_id", taskHandler.Delete)
}
