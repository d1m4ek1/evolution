package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InSocialTemplate(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		chatID := context.Param("chatId")

		if chatID != "" {
			fmt.Println(chatID)
		}

		redirectUser(ctx, context, &Redirector{
			Token:  token,
			UserId: userId,
		})

		replyBasedOnToken(ctx, &ReplyBaseOnToken{
			Define:  "insocial",
			Token:   token,
			UserId:  userId,
			Context: context,
		})
	})
}
