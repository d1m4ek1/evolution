package models

import (
	newerror "iNote/www/backend/pkg/newerror"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func SelectProfileData(ctx *sqlx.DB, userID int64) ([]string, error) {
	var aboutme []string
	if err := ctx.Get(pq.Array(&aboutme), `
		SELECT 
		    aboutme
		FROM 
		    settings 
		WHERE 
		    settings_id=(SELECT settings_id FROM identifiers WHERE user_id=$1)`, userID); err != nil {
		newerror.NewAppError("ctx.Get", err, pathToLogFile, isTimeAmPm)
		return nil, err
	}
	return aboutme, nil
}
