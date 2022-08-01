package models

import (
	newerror "iNote/www/backend/pkg/NewError"

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
		newerror.Wrap("ctx.DB.QueryRow", err)
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
		newerror.Wrap("ctx.DB.QueryRow", err)
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
		newerror.Wrap("ctx.DB.QueryRow", err)
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
		newerror.Wrap("ctx.DB.QueryRow", err)
		return "", "", err
	}

	return name, logo, nil
}
