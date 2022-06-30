package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	newerror "iNote/www/backend/pkg/NewError"
)

func SelectProfileData(ctx *sqlx.DB, userID int64) ([]string, error) {
	var aboutme []string
	if err := ctx.Get(pq.Array(&aboutme), `
		SELECT 
		    aboutme
		FROM 
		    settings 
		WHERE 
		    id=$1`, userID); err != nil {
		newerror.Wrap("ctx.Get", err)
		return nil, err
	}
	return aboutme, nil
}
