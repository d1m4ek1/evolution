package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/pkg/general"
)

func SettingsTemplate(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		title := context.Param("settingsType")
		if title == "" {
			title = "profile"
		}

		redirectUser(ctx, context, &Redirector{
			Token:  token,
			UserId: userId,
		})

		var settingsData general.SettingsData
		settingsData.SetTitle(title)

		replyBasedOnToken(ctx, &ReplyBaseOnToken{
			Define:   "settings",
			Token:    token,
			UserId:   userId,
			Settings: settingsData,
			Context:  context,
		})
	})
}
