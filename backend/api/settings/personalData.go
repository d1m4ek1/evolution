package settings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/NewError"
	"net/http"
	"strconv"
	"strings"
)

type personalData struct {
	BackupKeys bool   `json:"bcpk"`
	Email      string `json:"eml"`
}

func (v *personalData) validBackupKey(arr []string) {
	if len(arr) != 0 && len(arr) < 5 {
		v.BackupKeys = true
	} else {
		v.BackupKeys = false
	}
}

func getOldPersonalData(ctx *sqlx.DB, context *gin.Context, userId int64) {
	var userData personalData
	var backupKeys []string
	var err error

	userData.Email, backupKeys, err = models.SelectPersonalData(ctx, userId)
	if err != nil {
		newerror.Wrap("models.SelectPersonalData", err)
		return
	}

	userData.validBackupKey(backupKeys)

	context.JSON(http.StatusOK, userData)
}

func GetPersonalData(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")
		userIDConv, err := strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.Wrap("strconv.ParseInt", err)
			return
		}

		if token != "" && userId != "" {
			var user models.CheckSignin = models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.Wrap("user.CheckUserOnSignin", err)
				return
			}

			if user.Autorize {
				getOldPersonalData(ctx, context, userIDConv)
			} else {
				context.Redirect(http.StatusMovedPermanently, "signin")
			}
		}
	})
}

func SetPersonalData(ctx *sqlx.DB, context *gin.Context, userID int64) {
	var pdQuerys = [2]string{"password", "email"}
	var backupKey string

	if context.Query("backupkey_one") != "" && context.Query("backupkey_two") != "" &&
		context.Query("backupkey_three") != "" && context.Query("backupkey_four") != "" {

		backupKey = fmt.Sprintf("{%s, %s, %s, %s}",
			context.Query("backupkey_one"),
			context.Query("backupkey_two"),
			context.Query("backupkey_three"),
			context.Query("backupkey_four"))

		if err := models.SetBackupKeys(ctx, backupKey, userID); err != nil {
			newerror.Wrap("Query at db: 1", err)
			return
		}
	}

	var personaltpl []string
	for _, v := range pdQuerys {
		if context.Query(v) != "" {
			personaltpl = append(personaltpl, fmt.Sprintf("%s='%s'", v, context.Query(v)))
		}
	}

	if err := models.SetPersonalData(ctx, strings.Join(personaltpl, ", "), userID); err != nil {
		newerror.Wrap("models.SetPersonalData", err)
		return
	}
}

func SavePersonalData(ctx *sqlx.DB) gin.HandlerFunc {
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
				SetPersonalData(ctx, context, userIDConv)
			} else {
				context.Redirect(http.StatusMovedPermanently, "/signin")
			}
		}
	})
}
