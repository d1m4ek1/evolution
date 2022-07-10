package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/NewError"
	"net/http"
	"strconv"
)

func CheckChat(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")
		userIDConv, err := strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.Wrap("strconv.ParseInt", err)
			return
		}

		if token != "" && userId != "" {
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.Wrap("user.CheckUserOnSignin", err)
				return
			}

			userIdTwo, err := strconv.ParseInt(context.Query("user_id_two"), 10, 0)
			if err != nil {
				newerror.Wrap("strconv.ParseInt", err)
				return
			}

			if user.Autorize {
				chatData, err := models.SelectChat(ctx, userIDConv, userIdTwo)
				if err != nil {
					newerror.Wrap("models.SelectChat", err)
					return
				}

				context.JSON(http.StatusOK, chatData)
			}
		}
	})
}
