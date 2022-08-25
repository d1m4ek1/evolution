package dataprofile

import (
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/newerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const pathToLogFile string = "backend/logs/logs.txt"
const isTimeAmPm bool = true

// Path to error
const (
	pathToError string = "api/dataProfile -> Function "
)

const (
	errorSendAllData string = pathToError + "sendAllData"
)

type DataProfile struct {
	AboutMeTitle   string `json:"aboutmeTitle"`
	AboutMeContent string `json:"aboutmeContent"`
}

func sendAllData(ctx *sqlx.DB, context *gin.Context, userId string) {
	var dataProfile DataProfile

	userIDConv, err := strconv.ParseInt(userId, 10, 0)
	if err != nil {
		newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
		return
	}

	aboutme, err := models.SelectProfileData(ctx, userIDConv)
	if err != nil {
		newerror.NewAppError("Qmodels.SelectProfileData", err, pathToLogFile, isTimeAmPm)
		return
	}

	dataProfile.AboutMeTitle = aboutme[0]
	dataProfile.AboutMeContent = aboutme[1]

	context.JSON(http.StatusOK, dataProfile)
}

func ControlDataProfile(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		if context.Query("get_data") == "all" {
			sendAllData(ctx, context, context.Query("user_id"))
			return
		}
	})
}
