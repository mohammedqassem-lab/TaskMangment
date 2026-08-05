package route

import (
	"TaskMangment/Internal/handler"

	"github.com/gin-gonic/gin"
)

func CreateProject(r *gin.RouterGroup, ProjectHandler *handler.ProjectHandler) {
	r.POST("/Project/:id/create", ProjectHandler.Create)
}
func GetById(r *gin.RouterGroup, ProjectHAndler *handler.ProjectHandler) {
	r.GET("/Project/:id/GetbyId/:projectid", ProjectHAndler.GetById)
}
func Get(r *gin.RouterGroup, ProjectHandler *handler.ProjectHandler) {
	r.GET("/Project/:id/Get", ProjectHandler.Get)
}
func Update(r *gin.RouterGroup, ProjectHandler *handler.ProjectHandler) {
	r.PUT("/Project/:id/Update", ProjectHandler.Update)
}
func Delete(r *gin.RouterGroup, ProjectHandler *handler.ProjectHandler) {
	r.DELETE("/Project/:id/Delete/:projectid", ProjectHandler.Delete)
}
