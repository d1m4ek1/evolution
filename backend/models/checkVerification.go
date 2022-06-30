package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	newerror "iNote/www/backend/pkg/NewError"
)

func CheckVerifByCustomID(ctx *sqlx.DB, customID string) (isVerify int64, err error) {
	if err := ctx.Get(&isVerify, `
		SELECT 
		    id 
		FROM 
		    users_data 
		WHERE 
		    id=$1 OR user_custom_id=$2`, customID, customID); err != nil {
		newerror.Wrap("ctx.Get", err)
		return 0, err
	}
	return isVerify, nil
}

func CheckVerifByIDUser(ctx *sqlx.DB, checkStrID string, checkIntID int64) (isVerify bool, err error) {
	var newID string

	if checkStrID != "" {
		newID = checkStrID
	}

	if checkIntID != 0 {
		newID = fmt.Sprint(checkIntID)
	}

	if err := ctx.Get(&isVerify, `
		SELECT 
		    count(*) = 1 
		FROM 
		    users_data 
		WHERE 
		    id=$1`, newID); err != nil {
		newerror.Wrap("ctx.Get", err)
		return false, err
	}
	return isVerify, nil
}

func SelectLoginByIdToken(ctx *sqlx.DB, userID int64, token string) (login string, err error) {
	if err := ctx.Get(&login, `
		SELECT
		    login
		FROM
				users_data
		WHERE
		    id=$1 
		  AND
		    token=$2`, userID, token); err != nil {
		newerror.Wrap("ctx.Get", err)
		return "", err
	}
	return login, nil
}

func VerifUser(ctx *sqlx.DB, login, password string) (isVerify bool, err error) {
	if err := ctx.DB.QueryRow(`
		SELECT 
		    count(*) = 1
		FROM 
		    users_data
		WHERE 
		    login=$1 
		  AND 
		    password=$2`,
		login, password).Scan(&isVerify); err != nil {
		newerror.Wrap("ctx.DB.QueryRow", err)
		return false, err
	}

	return isVerify, nil
}
