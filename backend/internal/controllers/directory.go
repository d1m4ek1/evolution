package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func DirectoryTemplate(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		replyBasedOnToken(ctx, &ReplyBaseOnToken{
			Define:  "directory",
			Token:   token,
			UserId:  userId,
			Context: context,
		})
	})
}
