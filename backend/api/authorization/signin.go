package authorization

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/newerror"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UserData struct {
	UserID        int64  `json:"user_id"`
	NetworkStatus string `json:"netStatus"`
}

type SignInData struct {
	Login    string
	Password string
	NewToken string
	OldToken sql.NullString
}

func createToken() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var characters = 64

	b := make([]rune, characters)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			newerror.NewAppError("rand.Int", err, pathToLogFile, isTimeAmPm)
			return ""
		}

		b[i] = letterRunes[n.Int64()]
	}
	return string(b)
}

func getDataSignIn(ctx *sqlx.DB, context *gin.Context, s *SignInData) (UserData, bool) {
	var err error
	var data UserData

	data.UserID, data.NetworkStatus, err = models.SignInData(ctx, s.Login, s.Password, s.NewToken)
	if err != nil {
		newerror.NewAppError("models.SignInData", err, pathToLogFile, isTimeAmPm)
		context.JSON(http.StatusOK, gin.H{
			"error": "Неверный логин или пароль",
		})
		return UserData{}, false
	}

	context.SetCookie("token", s.NewToken, 0, "/", "", false, true)
	context.SetCookie("userId", fmt.Sprint(data.UserID), 0, "/", "", false, false)

	return data, true
}

func SignIn(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		signIn := context.Query("signin")

		if signIn == "true" {
			userData, jsonBool := getDataSignIn(ctx, context, &SignInData{
				Login:    context.Query("login"),
				Password: context.Query("password"),
				NewToken: createToken(),
			})

			if jsonBool {
				context.JSON(http.StatusOK, userData)
			}
		}
	})
}
