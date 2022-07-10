package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/NewError"
	"net/http"
	"strconv"
)

func getUserAuthData(context *gin.Context) (string, int64, error) {
	token, _ := context.Cookie("token")
	userId, _ := context.Cookie("userId")
	userIDConv, err := strconv.ParseInt(userId, 10, 0)
	if err != nil {
		newerror.Wrap("strconv.ParseInt", err)
		return "", 0, err
	}

	return token, userIDConv, nil
}

func GetUserCardMessages(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, userID, err := getUserAuthData(context)
		if err != nil {
			newerror.Wrap("getUserAuthData", err)
		}

		user := models.CheckSignin{
			Id:       userID,
			Token:    token,
			Autorize: false,
		}
		if err := user.CheckUserOnSignin(ctx); err != nil {
			newerror.Wrap("user.CheckUserOnSignin", err)
			return
		}

		if user.Autorize {
			isCardSubscriptions, isCardSubscribers, err := models.SelectUserCardMessages(ctx, user.Id)
			if err != nil {
				newerror.Wrap("models.SelectUserCardMessages", err)
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"isCardSubscriptions": isCardSubscriptions,
				"isCardSubscribers":   isCardSubscribers,
			})
		} else {
			context.Redirect(http.StatusMovedPermanently, "/signin")
		}
	})
}
