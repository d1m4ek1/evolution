package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/NewError"
	"net/http"
	"strconv"
)

func CheckAuthoriztion(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token := context.Query("token")
		userId := context.Query("userId")

		if token != "" && userId != "" {
			userIDConv, err := strconv.ParseInt(userId, 10, 0)
			if err != nil {
				newerror.Wrap("strconv.ParseInt", err)
				return
			}

			var authorize = models.CheckSignin{
				Id:    userIDConv,
				Token: token,
			}

			if err := authorize.CheckUserOnSignin(ctx); err != nil {
				newerror.Wrap("authorize.CheckUserOnSignin", err)
				return
			}

			context.JSON(http.StatusOK, gin.H{
				"isVerify": authorize.Autorize,
			})
		}
	})
}
