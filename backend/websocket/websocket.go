package websocket

import (
	"encoding/json"
	"fmt"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/NewError"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
)

var users = make(map[int64]WSConnect)

var upgraded = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	Subprotocols:    []string{"contact"},
}

type WSConnect struct {
	Id        int64           `json:"id"`
	Name      string          `json:"name"`
	NetStatus string          `json:"netStatus"`
	Conn      *websocket.Conn `json:"conn"`
}

type GettedMessage struct {
	ChatId         string `json:"chatId"`
	SenderId       int64  `json:"sender_id"`
	RecipientId    int64  `json:"recipient_id"`
	IsMessageCheck bool   `json:"isMessageCheck"`
	Message        string `json:"message"`
	Date           string `json:"date"`
}

func appendClient(conn *websocket.Conn, login string, userID int64) {
	if _, ok := users[userID]; !ok {
		users[userID] = WSConnect{
			Id:        userID,
			Name:      login,
			NetStatus: "online",
			Conn:      conn,
		}
	}
}

func listenConnect(ctx *sqlx.DB, conn *websocket.Conn, login string, userID int64) {
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			delete(users, userID)

			if err := models.SetNetworkStatusOffline(ctx, userID); err != nil {
				newerror.Wrap("models.SetNetworkStatusOnline", err)
				return
			}

			break
		}

		if message != nil {
			var messageJSON GettedMessage
			if err := json.Unmarshal(message, &messageJSON); err != nil {
				newerror.Wrap("json.Unmarshal", err)
				return
			}
			if messageJSON.IsMessageCheck {
				if err := models.SetMessage(ctx, messageJSON.ChatId, messageJSON.Message); err != nil {
					newerror.Wrap("models.SetMessage", err)
					return
				}

				if _, ok := users[messageJSON.RecipientId]; ok {
					if err := users[messageJSON.RecipientId].Conn.WriteMessage(mt,
						[]byte(fmt.Sprintf(`{"checked":true, "chatId": %s}`, messageJSON.ChatId))); err != nil {
						newerror.Wrap("users[messageJSON.RecipientId].Conn.WriteMessage", err)
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
					newerror.Wrap("json.Marshal", err)
					return
				}

				if err := models.SetNewMessage(ctx, messageJSON.ChatId, string(messageString)); err != nil {
					newerror.Wrap("models.SetNewMessage", err)
					return
				}

				messageFull, err := json.Marshal(gin.H{
					"chatId":    messageJSON.ChatId,
					"sender_id": messageJSON.SenderId,
					"message":   messageJSON.Message,
					"date":      messageJSON.Date,
				})
				if err != nil {
					newerror.Wrap("json.Marshal", err)
					return
				}

				if _, ok := users[messageJSON.RecipientId]; ok {
					if err := users[messageJSON.RecipientId].Conn.WriteMessage(mt, messageFull); err != nil {
						newerror.Wrap("users[messageJSON.RecipientId].Conn.WriteMessage", err)
						return
					}
				}
			}
		}
	}
}

func WebSocketConnect(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userID, _ := context.Cookie("userId")

		if token != "" && userID != "" {
			userIDConv, err := strconv.ParseInt(userID, 10, 0)
			if err != nil {
				newerror.Wrap("strconv.ParseInt", err)
				return
			}

			conn, err := upgraded.Upgrade(context.Writer, context.Request, nil)
			if err != nil {
				newerror.Wrap("upgraded.Upgrade", err)
				return
			}
			defer conn.Close()

			login, err := models.SelectLoginByIdToken(ctx, userIDConv, token)
			if err != nil {
				newerror.Wrap("models.SelectLoginByIdToken", err)
				return
			}

			if err := models.SetNetworkStatusOnline(ctx, userIDConv); err != nil {
				newerror.Wrap("models.SetNetworkStatusOnline", err)
				return
			}

			appendClient(conn, login, userIDConv)
			listenConnect(ctx, conn, login, userIDConv)
		}
	})
}
