package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/NewError"
	"iNote/www/backend/pkg/general"
	"net/http"
)

func createAccount(ctx *sqlx.DB, context *gin.Context, s general.SignUpData) {
	user, err := models.CheckLogin(ctx, s.Login)
	if err != nil {
		newerror.Wrap("models.CheckLogin", err)
		return
	}

	if user {
		context.JSON(http.StatusOK, gin.H{
			"error": "Логин занят!",
		})
		return
	}

	if !user {
		if err := models.CreateAccount(ctx, s.Login, s.Password, s.Email, s.Token, s.Nickname); err != nil {
			newerror.Wrap("models.CreateAccount", err)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"aut": true,
		})
	}
}

func SignUp(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		keyWords := [5]string{"nickname", "email", "login", "password", "token"}
		var valid int

		for _, v := range keyWords {
			if context.Query(v) != "" {
				valid += 1
			}
		}

		if valid == 5 {
			signUpData := general.SignUpData{}
			signUpData.ValidData(&general.SignUpData{
				Nickname: context.Query("nickname"),
				Email:    context.Query("email"),
				Login:    context.Query("login"),
				Password: context.Query("password"),
				Token:    context.Query("token"),
			})

			if signUpData.Login != "" && signUpData.Password != "" && signUpData.Nickname != "" && signUpData.Email != "" && signUpData.Token != "" {
				createAccount(ctx, context, signUpData)
			} else {
				context.JSON(http.StatusOK, gin.H{
					"error": "Некорректное значение",
				})
			}
		} else {
			context.JSON(http.StatusOK, gin.H{
				"error": "Некорректное значение",
			})
		}
	})
}
