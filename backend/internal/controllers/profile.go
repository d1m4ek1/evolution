package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	userdata "iNote/www/backend/api/userData"
)

func ProfileTemplate(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		userUrlId := context.Param("userName")

		profileDefaultData := userdata.GetUserDataStatic(ctx, token, userUrlId, context)

		replyBasedOnToken(ctx, &ReplyBaseOnToken{
			Define:    "profile",
			UserUrlId: userUrlId,
			Token:     token,
			UserId:    userId,
			Profile:   profileDefaultData,
			Context:   context,
		})
	})
}
