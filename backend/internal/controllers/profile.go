package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	userdata "iNote/www/backend/api/userData"
	newerror "iNote/www/backend/pkg/NewError"
	"strconv"
)

func ProfileTemplate(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		var userIDConv int64
		var err error
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")

		if userId != "" {
			userIDConv, err = strconv.ParseInt(userId, 10, 0)
			if err != nil {
				newerror.Wrap("strconv.ParseInt", err)
				return
			}
		}

		userUrlId := context.Param("userName")

		profileDefaultData := userdata.GetUserDataStatic(ctx, token, userUrlId, userIDConv)

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
