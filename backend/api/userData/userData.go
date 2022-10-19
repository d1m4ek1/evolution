package userdata

import (
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/general"
	"iNote/www/backend/pkg/newerror"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const pathToLogFile string = "backend/logs/logs.txt"
const isTimeAmPm bool = true

type DataArray struct {
	Position []string `json:"position"`
	Audience []string `json:"audience"`
}

func profileDefault(ctx *sqlx.DB, id string) (general.ProfileData, int64, error) {
	var err error
	var pdd general.ProfileData
	var dataArray DataArray

	isVerify, err := models.CheckVerifByCustomID(ctx, id)
	if err != nil {
		newerror.NewAppError("models.CheckVerifByCustomID", err, pathToLogFile, isTimeAmPm)
		return general.ProfileData{}, 0, err
	}

	if isVerify != 0 {
		pdd.Name, pdd.NetworkStatus, pdd.Logo, pdd.Banner, pdd.Verif,
			dataArray.Position, dataArray.Audience, err = models.SelectProfileDefault(ctx, isVerify)
		if err != nil {
			newerror.NewAppError("models.SelectProfileDefault", err, pathToLogFile, isTimeAmPm)
			return general.ProfileData{}, 0, err
		}

		pdd.Audience = len(dataArray.Audience)
		pdd.ValidLogoBanner(pdd.Logo, pdd.Banner)

		return pdd, isVerify, nil
	}

	return general.ProfileData{}, 0, nil
}

func GetUserDataStatic(ctx *sqlx.DB, token, userUrlId string, context *gin.Context) general.ProfileData {
	profileDefaultData, isVerify, err := profileDefault(ctx, userUrlId)
	if err != nil {
		context.Redirect(http.StatusFound, "/")
	}

	if token != "" {
		if isVerify != 0 && token != "" {
			if err := profileDefaultData.ProfileUser(ctx, isVerify, token); err != nil {
				newerror.NewAppError("profileDefaultData.ProfileUser", err, pathToLogFile, isTimeAmPm)
				return general.ProfileData{}
			}

			profileDefaultData.NetworkStatus = "online"
		}
	}

	return profileDefaultData
}
