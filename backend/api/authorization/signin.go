package authorization

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/NewError"
	"net/http"
)

type UserData struct {
	UserId int64 `json:"user_id"`
	//Name   string `json:"name"`
	//Birthday      sql.NullString `json:"birthday"`
	//Position      []string       `json:"position"`
	//MainAddress   []int          `json:"mainAddress"`
	NetworkStatus string `json:"netStatus"`
	//Logo          string         `json:"logo"`
	//Banner        string         `json:"banner"`
	//Audience      []string       `json:"audience"`
	//Verif         sql.NullString `json:"verif"`
	//FirstName     sql.NullString `json:"firstName"`
	//LastName      sql.NullString `json:"lastName"`
	Token string `json:"olt"`
}

type SignInData struct {
	Login    string
	Password string
	NewToken string
	OldToken sql.NullString
}

func getDataSignIn(ctx *sqlx.DB, context *gin.Context, s *SignInData) (UserData, bool) {
	var err error
	var data UserData

	data.UserId, data.NetworkStatus, err = models.SignInData(ctx, s.Login, s.Password, s.NewToken)
	if err != nil {
		newerror.Wrap("models.SignInData", err)
		context.JSON(http.StatusOK, gin.H{
			"error": "Неверный логин или пароль",
		})
		return UserData{}, false
	}

	data.Token = s.NewToken

	return data, true
}

func SignIn(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		signIn := context.Query("signin")

		if signIn == "true" {

			userData, jsonBool := getDataSignIn(ctx, context, &SignInData{
				Login:    context.Query("login"),
				Password: context.Query("password"),
				NewToken: context.Query("token"),
			})

			if jsonBool {
				context.JSON(http.StatusOK, userData)
			}
		}
	})
}
