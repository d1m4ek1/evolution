package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/NewError"
	"net/http"
	"strconv"
)

func ConfirmPassword(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")
		userIDConv, err := strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.Wrap("strconv.ParseInt", err)
			return
		}

		if token != "" && userId != "" {
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}

			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.Wrap("user.CheckUserOnSignin", err)
				return
			}

			if user.Autorize {
				confirmPassword := models.ConfirmitadePassword{}
				if err := confirmPassword.ConfirmPassword(ctx, userIDConv, token, context.Query("conf_pass")); err != nil {
					newerror.Wrap("confirmPassword.ConfirmPassword", err)
					return
				}

				context.JSON(http.StatusOK, confirmPassword)
			} else {
				context.Redirect(http.StatusMovedPermanently, "/signin")
			}
		}
	})
}
