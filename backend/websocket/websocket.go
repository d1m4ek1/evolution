package websocket

import (
	"encoding/json"
	"fmt"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/newerror"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
)

func listenConnect(ctx *sqlx.DB, conn *websocket.Conn, login string, userID int64) {
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			delete(users, userID)

			if err := models.SetNetworkStatusOffline(ctx, userID); err != nil {
				newerror.NewAppError("models.SetNetworkStatusOnline", err, pathToLogFile, isTimeAmPm)
				return
			}

			break
		}

		if message != nil {
			var messageJSON GettedMessage
			if err := json.Unmarshal(message, &messageJSON); err != nil {
				newerror.NewAppError("json.Unmarshal", err, pathToLogFile, isTimeAmPm)
				return
			}

			if messageJSON.IsMessageCheck {
				if err := models.SetMessage(ctx, messageJSON.ChatId, messageJSON.Message); err != nil {
					newerror.NewAppError("models.SetMessage", err, pathToLogFile, isTimeAmPm)
					return
				}

				if _, ok := users[messageJSON.RecipientId]; ok {
					if err := users[messageJSON.RecipientId].Conn.WriteMessage(mt,
						[]byte(fmt.Sprintf(`{"checked":true, "chatId": %s}`, messageJSON.ChatId))); err != nil {
						newerror.NewAppError("users[messageJSON.RecipientId].Conn.WriteMessage", err, pathToLogFile, isTimeAmPm)
						return
					}
				}
			} else {
				messageString, err := json.Marshal(gin.H{
					"sender_id": messageJSON.SenderId,
					"message":   messageJSON.Message,
					"date":      messageJSON.Date,
				})
				if err != nil {
					newerror.NewAppError("json.Marshal", err, pathToLogFile, isTimeAmPm)
					return
				}

				if err := models.SetNewMessage(ctx, messageJSON.ChatId, string(messageString)); err != nil {
					newerror.NewAppError("models.SetNewMessage", err, pathToLogFile, isTimeAmPm)
					return
				}

				messageFull, err := json.Marshal(gin.H{
					"chatId":    messageJSON.ChatId,
					"sender_id": messageJSON.SenderId,
					"message":   messageJSON.Message,
					"date":      messageJSON.Date,
				})
				if err != nil {
					newerror.NewAppError("json.Marshal", err, pathToLogFile, isTimeAmPm)
					return
				}

				if _, ok := users[messageJSON.RecipientId]; ok {
					if err := users[messageJSON.RecipientId].Conn.WriteMessage(mt, messageFull); err != nil {
						newerror.NewAppError("users[messageJSON.RecipientId].Conn.WriteMessage", err, pathToLogFile, isTimeAmPm)
						return
					}
				}
			}
		}
	}
}
