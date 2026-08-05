package middelware

import (
	logfile "TaskMangment/Internal/LogFile"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			fmt.Println("ErrorMiddleware: ", c.Errors.Last().Error())
			logfile.LogErr(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}
