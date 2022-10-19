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
		token, err := context.Cookie("token")
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"isVerify": false,
			})
			return
		}

		userID, err := context.Cookie("userId")
		if err != nil {
			newerror.NewAppError("context.Cookie", err, pathToLogFile, isTimeAmPm)
			return
		}

		if token != "" && userID != "" {
			userIDConv, err := strconv.ParseInt(userID, 10, 0)
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
