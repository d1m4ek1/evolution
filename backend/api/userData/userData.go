package userdata

import (
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/NewError"
	"iNote/www/backend/pkg/general"
)

type DataArray struct {
	Position []string `json:"position"`
	Audience []string `json:"audience"`
}

func profileDefault(ctx *sqlx.DB, id string) (general.ProfileData, int64) {
	var err error
	var pdd general.ProfileData
	var dataArray DataArray

	isVerify, err := models.CheckVerifByCustomID(ctx, id)
	if err != nil {
		newerror.Wrap("models.CheckVerifByCustomID", err)
		return general.ProfileData{}, 0
	}

	if isVerify != 0 {
		pdd.Name, pdd.NetworkStatus, pdd.Logo, pdd.Banner, pdd.Verif,
			dataArray.Position, dataArray.Audience, err = models.SelectProfileDefault(ctx, isVerify)
		if err != nil {
			newerror.Wrap("models.SelectProfileDefault", err)
			return general.ProfileData{}, 0
		}

		pdd.Audience = len(dataArray.Audience)
		pdd.ValidLogoBanner(pdd.Logo, pdd.Banner)

		return pdd, isVerify
	}

	return general.ProfileData{}, 0
}

func GetUserDataStatic(ctx *sqlx.DB, token, userUrlId string, userID int64) general.ProfileData {
	profileDefaultData, isVerify := profileDefault(ctx, userUrlId)

	if token != "" {
		if isVerify != 0 && token != "" {
			if err := profileDefaultData.ProfileUser(ctx, isVerify, token); err != nil {
				newerror.Wrap("profileDefaultData.ProfileUser", err)
				return general.ProfileData{}
			}

			profileDefaultData.NetworkStatus = "online"
		}
	}

	return profileDefaultData
}
