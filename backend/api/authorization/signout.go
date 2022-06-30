package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/NewError"
)

func SignOut(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		if token != "" {
			if err := models.UserSignOut(ctx, userId); err != nil {
				newerror.Wrap("models.UserSignOut", err)
				return
			}
		}
	})
}
