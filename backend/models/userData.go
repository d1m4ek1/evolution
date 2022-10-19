package models

import (
	newerror "iNote/www/backend/pkg/newerror"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func SelectProfileDefault(ctx *sqlx.DB, userID int64) (name, netStatus, logo, banner string, verify bool, position, audience []string, err error) {
	if err := ctx.DB.QueryRow(`
		SELECT 
    		name, 
    		position, 
    		subscribers, 
    		verification, 
    		net_status 
		FROM 
		    users 
		WHERE 
		    id=$1`, userID).Scan(&name,
		pq.Array(&position), pq.Array(&audience),
		&verify, &netStatus); err != nil {
		newerror.NewAppError("ctx.DB.QueryRow", err, pathToLogFile, isTimeAmPm)
		return "", "", "", "", false, nil, nil, err
	}

	if err := ctx.DB.QueryRow(`
		SELECT 
		    sgs.logo, 
		    sgs.banner 
		FROM 
		    settings sgs,
		    identifiers ids 
		WHERE 
		    ids.user_id=$1 
		  AND 
		    ids.settings_id=sgs.settings_id`, userID).Scan(&logo, &banner); err != nil {
		newerror.NewAppError("ctx.DB.QueryRow", err, pathToLogFile, isTimeAmPm)
		return "", "", "", "", false, nil, nil, err
	}

	return name, netStatus, logo, banner, verify, position, audience, nil
}

func SelectUserDataDefault(ctx *sqlx.DB, userID int64) (name, logo string, err error) {
	if err := ctx.DB.QueryRow(`
	SELECT
		name
	FROM
		users
	WHERE
		id=$1`, userID).Scan(&name); err != nil {
		newerror.NewAppError("ctx.DB.QueryRow", err, pathToLogFile, isTimeAmPm)
		return "", "", nil
	}

	if err := ctx.DB.QueryRow(`
		SELECT 
		    sgs.logo
		FROM 
		    settings sgs,
		    identifiers ids 
		WHERE 
		    ids.user_id=$1 
		  AND 
		    ids.settings_id=sgs.settings_id`, userID).Scan(&logo); err != nil {
		newerror.NewAppError("ctx.DB.QueryRow", err, pathToLogFile, isTimeAmPm)
		return "", "", err
	}

	return name, logo, nil
}

func SelectUserName(ctx *sqlx.DB, userID int64) (string, error) {
	var userName string
	if err := ctx.Get(&userName, `
	SELECT
		name
	FROM
		users
	WHERE
		user_id=$1`, userID); err != nil {
		newerror.NewAppError("ctx.Get", err, pathToLogFile, isTimeAmPm)
		return "", err
	}
	return userName, nil
}
