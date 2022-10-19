package dataprofile

import (
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/general"
	"iNote/www/backend/pkg/newerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const pathToLogFile string = "backend/logs/logs.txt"
const isTimeAmPm bool = true

func sendAllData(ctx *sqlx.DB, context *gin.Context, userID string) {
	visitorIDCookie, _ := context.Cookie("userId")
	var err error

	userIDConv, err := strconv.ParseInt(userID, 10, 0)
	if err != nil {
		newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
		return
	}

	profileData, err := models.SelectProfileData(ctx, userIDConv)
	if err != nil {
		newerror.NewAppError("Qmodels.SelectProfileData", err, pathToLogFile, isTimeAmPm)
		return
	}

	if visitorID, _ := models.CheckVerifByCustomID(ctx, visitorIDCookie); visitorID != 0 {
		profileData.VisitorIsAuthorized = true
	}

	profileData.Banner, profileData.Logo = general.ValidLogoBanner(profileData.Logo, profileData.Banner)

	profileData.URLID = userID

	context.JSON(http.StatusOK, profileData)
}

func ControlDataProfile(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		if context.Query("get_data") == "all" {
			sendAllData(ctx, context, context.Query("user_id"))
			return
		}
	})
}
