package authorization

import (
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/newerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const pathToLogFile string = "backend/logs/logs.txt"
const isTimeAmPm bool = true

func CheckAuthoriztion(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token := context.Query("token")
		userId := context.Query("userId")

		if token != "" && userId != "" {
			userIDConv, err := strconv.ParseInt(userId, 10, 0)
			if err != nil {
				newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
				return
			}

			var authorize = models.CheckSignin{
				Id:    userIDConv,
				Token: token,
			}

			if err := authorize.CheckUserOnSignin(ctx); err != nil {
				newerror.NewAppError("authorize.CheckUserOnSignin", err, pathToLogFile, isTimeAmPm)
				return
			}

			context.JSON(http.StatusOK, gin.H{
				"isVerify": authorize.Autorize,
			})
		}
	})
}
