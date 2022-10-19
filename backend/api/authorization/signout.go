package authorization

import (
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/newerror"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SignOut(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		if token != "" {
			if err := models.UserSignOut(ctx, userId); err != nil {
				newerror.NewAppError("models.UserSignOut", err, pathToLogFile, isTimeAmPm)
				return
			}

			context.SetCookie("token", "", -1, "/", "localhost", false, true)
			context.SetCookie("userId", "", -1, "/", "localhost", false, true)
		}
	})
}
