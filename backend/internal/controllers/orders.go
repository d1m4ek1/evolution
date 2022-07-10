package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func OrdersTemplate(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		redirectUser(ctx, context, &Redirector{
			Token:  token,
			UserId: userId,
		})

		replyBasedOnToken(ctx, &ReplyBaseOnToken{
			Define:  "orders",
			Token:   token,
			UserId:  userId,
			Context: context,
		})
	})
}
