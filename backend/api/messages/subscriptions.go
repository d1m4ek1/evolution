package messages

import (
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/newerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func getUserAuthData(context *gin.Context) (string, int64, error) {
	token, _ := context.Cookie("token")
	userId, _ := context.Cookie("userId")
	if token != "" && userId != "" {
		userIDConv, err := strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
			return "", 0, err
		}

		return token, userIDConv, nil
	}

	return "", 0, nil
}

func GetUserCardMessages(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, userID, err := getUserAuthData(context)
		if err != nil {
			newerror.NewAppError("getUserAuthData", err, pathToLogFile, isTimeAmPm)
		}

		if token == "" && userID == 0 {
			context.JSON(http.StatusOK, gin.H{
				"isAuthorized": false,
			})
			return
		}

		user := models.CheckSignin{
			Id:       userID,
			Token:    token,
			Autorize: false,
		}
		if err := user.CheckUserOnSignin(ctx); err != nil {
			newerror.NewAppError("user.CheckUserOnSignin", err, pathToLogFile, isTimeAmPm)
			return
		}

		if user.Autorize {
			isCardSubscriptions, isCardSubscribers, err := models.SelectUserCardMessages(ctx, user.Id)
			if err != nil {
				newerror.NewAppError("models.SelectUserCardMessages", err, pathToLogFile, isTimeAmPm)
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"isCardSubscriptions": isCardSubscriptions,
				"isCardSubscribers":   isCardSubscribers,
				"isAuthorized":        true,
			})
		} else {
			context.Redirect(http.StatusFound, "/signin")
		}
	})
}
