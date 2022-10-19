package models

import (
	newerror "iNote/www/backend/pkg/newerror"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProfileData struct {
	Aboutme             pq.StringArray `json:"aboutme" db:"aboutme"`
	VisitorIsAuthorized bool           `json:"visitorIsAuthorized"`
	Name                string         `json:"name"`
	Logo                string         `json:"logo"`
	Banner              string         `json:"banner"`
	NetworkStatus       string         `json:"networkStatus" db:"net_status"`
	URLID               string         `json:"urlId"`
}

func SelectProfileData(ctx *sqlx.DB, userID int64) (ProfileData, error) {
	var profileData ProfileData
	if err := ctx.Get(&profileData, `
		SELECT 
		    aboutme,
				logo,
				banner
		FROM 
		    settings
		WHERE 
		    settings_id=(SELECT settings_id FROM identifiers WHERE user_id=$1)`, userID); err != nil {
		newerror.NewAppError("ctx.Get", err, pathToLogFile, isTimeAmPm)
		return ProfileData{}, err
	}

	if err := ctx.Get(&profileData, `
	SELECT 
			name,
			net_status
	FROM 
			users
	WHERE 
			user_id=$1`, userID); err != nil {
		newerror.NewAppError("ctx.Get", err, pathToLogFile, isTimeAmPm)
		return ProfileData{}, err
	}
	return profileData, nil
}
