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
		var err error

		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")
		userIDConv, err := strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.Wrap("strconv.ParseInt", err)
			return
		}

		if token != "" && userId != "" {
			var userIdTwo int64
			var chatID int64
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.Wrap("user.CheckUserOnSignin", err)
				return
			}

			if context.Query("user_id_two") != "" {
				userIdTwo, err = strconv.ParseInt(context.Query("user_id_two"), 10, 0)
				if err != nil {
					newerror.Wrap("strconv.ParseInt", err)
					return
				}
			}

			if context.Query("chat_id") != "" {
				chatID, err = strconv.ParseInt(context.Query("chat_id"), 10, 0)
				if err != nil {
					newerror.Wrap("strconv.ParseInt", err)
					return
				}
			}

			if user.Autorize {
				chatData, err := models.SelectChat(ctx, userIDConv, userIdTwo, chatID)
				if err != nil {
					newerror.Wrap("models.SelectChat", err)
					return
				}

				context.JSON(http.StatusOK, chatData)
			}
		}
	})
}

func GetAllChats(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		if token != "" && userId != "" {
			userIDConv, err := strconv.ParseInt(userId, 10, 0)
			if err != nil {
				newerror.Wrap("strconv.ParseInt", err)
				return
			}
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.Wrap("user.CheckUserOnSignin", err)
				return
			}

			if user.Autorize {
				chatDataItems, err := models.SelectChatItems(ctx, userIDConv)
				if err != nil {
					newerror.Wrap("models.SelectChatItems", err)
					return
				}

				if chatDataItems != nil {
					context.JSON(http.StatusOK, chatDataItems)
				} else {
					context.JSON(http.StatusOK, gin.H{
						"error": "Вы еще не переписывались",
					})
				}
			}
		}
	})
}
