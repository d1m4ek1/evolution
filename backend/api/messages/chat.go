package messages

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

func CheckChat(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		var err error

		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")
		userIDConv, err := strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
			return
		}

		if token != "" && userId != "" {
			var chatID int64
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.NewAppError("user.CheckUserOnSignin", err, pathToLogFile, isTimeAmPm)
				return
			}

			if context.Query("chat_id") != "" {
				chatID, err = strconv.ParseInt(context.Query("chat_id"), 10, 0)
				if err != nil {
					newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
					return
				}
			}

			if user.Autorize {
				chatData, err := models.SelectChat(ctx, userIDConv, user.Id, chatID)
				if err != nil {
					newerror.NewAppError("models.SelectChat", err, pathToLogFile, isTimeAmPm)
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
				newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
				return
			}
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.NewAppError("user.CheckUserOnSignin", err, pathToLogFile, isTimeAmPm)
				return
			}

			if user.Autorize {
				chatDataItems, err := models.SelectChatItems(ctx, userIDConv)
				if err != nil {
					newerror.NewAppError("models.SelectChatItems", err, pathToLogFile, isTimeAmPm)
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
