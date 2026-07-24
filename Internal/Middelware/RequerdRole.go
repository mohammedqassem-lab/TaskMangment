package middelware

import (
	repositry "TaskMangment/Internal/Repositry"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RequireRole(
	repo repositry.IWorkspaceRepository,
	allowedRoles ...string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		userValue, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		userID := userValue.(string)
		userIDInt, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid user ID",
			})
			return
		}
		workspaceID, err := strconv.ParseInt(
			c.Param("id"),
			10,
			64,
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid workspace id",
			})
			return
		}

		role, err := repo.GetRole(
			c.Request.Context(),
			workspaceID,
			userIDInt,
		)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   err.Error(),
				"message": "Access denied",
			})
			return
		}

		for _, allowed := range allowedRoles {

			if role == allowed {

				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Access denied",
		})
	}
}
