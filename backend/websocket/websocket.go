package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/NewError"
	"strconv"
)

var users = make(map[string]*WSConnect)

var upgraded = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WSConnect struct {
	Name string          `json:"json"`
	Conn *websocket.Conn `json:"conn"`
}

func appendClient(conn *websocket.Conn, login string) {
	if _, ok := users[login]; !ok {
		users[login] = &WSConnect{
			Name: login,
			Conn: conn,
		}
	}
}

func listenConnect(ctx *sqlx.DB, conn *websocket.Conn, login string, userID int64) {
	for {
		mt, _, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			delete(users, login)

			if err := models.SetNetworkStatusOffline(ctx, userID); err != nil {
				newerror.Wrap("models.SetNetworkStatusOnline", err)
				return
			}

			break
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

			appendClient(conn, login)
			listenConnect(ctx, conn, login, userIDConv)
		}
	})
}
