package authorization

import (
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/newerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func ConfirmPassword(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userID, _ := context.Cookie("userId")
		userIDConv, err := strconv.ParseInt(userID, 10, 0)
		if err != nil {
			newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
			return
		}

		if token != "" && userID != "" {
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}

			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.NewAppError("user.CheckUserOnSignin", err, pathToLogFile, isTimeAmPm)
				return
			}

			if user.Autorize {
				confirmPassword := models.ConfirmitadePassword{}
				if err := confirmPassword.ConfirmPassword(ctx, userIDConv, token, context.Query("conf_pass")); err != nil {
					newerror.NewAppError("confirmPassword.ConfirmPassword", err, pathToLogFile, isTimeAmPm)
					return
				}

				context.JSON(http.StatusOK, confirmPassword)
			} else {
				context.Redirect(http.StatusFound, "/signin")
			}
		}
	})
}
