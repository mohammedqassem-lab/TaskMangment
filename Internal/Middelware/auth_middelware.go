package middelware

import (
	auth "TaskMangment/Internal/Auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleeare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"massage": "no jwt in the header of requst",
			})
			ctx.Abort()
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		clims, err := auth.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"massage": "you are not auth",
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", clims.UserID)
		ctx.Next()
	}
}
