package authorization

import (
	"fmt"
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/general"
	"iNote/www/backend/pkg/newerror"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func createAccount(ctx *sqlx.DB, context *gin.Context, s general.SignUpData) {
	user, err := models.CheckLogin(ctx, s.Login)
	if err != nil {
		newerror.NewAppError("models.CheckLogin", err, pathToLogFile, isTimeAmPm)
		return
	}

	if user {
		context.JSON(http.StatusOK, gin.H{
			"error": "Логин занят!",
		})
		return
	}

	if !user {
		identificateID, err := models.CreateAccount(ctx, s.Login, s.Password, s.Email, s.Token, s.Nickname)
		if err != nil {
			newerror.NewAppError("models.CreateAccount", err, pathToLogFile, isTimeAmPm)
			return
		}

		context.SetCookie("token", s.Token, 0, "/", "", false, true)
		context.SetCookie("userId", fmt.Sprint(identificateID), 0, "/", "", false, false)

		context.JSON(http.StatusOK, gin.H{
			"aut": true,
		})
	}
}

func SignUp(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		signUpData := general.SignUpData{}
		signUpData.ValidData(&general.SignUpData{
			Nickname: context.Query("nickname"),
			Email:    context.Query("email"),
			Login:    context.Query("login"),
			Password: context.Query("password"),
			Token:    createToken(),
		})

		if signUpData.Login != "" && signUpData.Password != "" && signUpData.Nickname != "" && signUpData.Email != "" && signUpData.Token != "" {
			createAccount(ctx, context, signUpData)
		} else {
			context.JSON(http.StatusOK, gin.H{
				"error": "Некорректное значение",
			})
		}
	})
}
