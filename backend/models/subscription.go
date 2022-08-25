package models

import (
	"database/sql"
	newerror "iNote/www/backend/pkg/newerror"

	"github.com/jmoiron/sqlx"
)

func SelectSubscriber(ctx *sqlx.DB, userID int64, checkUserID string) (isSubscriber bool, err error) {
	if err := ctx.Get(&isSubscriber, `
	SELECT
	    $1=any(subscribers) 
	FROM 
	    users 
	WHERE 
	    user_id=$2`, userID, checkUserID); err != nil {
		newerror.NewAppError("ctx.Get", err, pathToLogFile, isTimeAmPm)
		return false, err
	}
	return isSubscriber, nil
}

func SetAppendSubscriber(ctx *sqlx.DB, userID int64, appendUserID int64) error {
	if _, err := ctx.DB.Exec(`
	UPDATE 
	    users
	SET
	    subscribers=array_append(subscribers, $1)
	WHERE user_id=$2`, userID, appendUserID); err != nil {
		newerror.NewAppError("ctx.DB.Exec", err, pathToLogFile, isTimeAmPm)
		return err
	}

	if _, err := ctx.DB.Exec(`
	UPDATE 
	    users
	SET
	    subscriptions=array_append(subscriptions, $1)
	WHERE user_id=$2`, appendUserID, userID); err != nil {
		newerror.NewAppError("ctx.DB.Exec", err, pathToLogFile, isTimeAmPm)
		return err
	}
	return nil
}

func SetDeleteSubscriber(ctx *sqlx.DB, userID int64, deleteUserId int64) error {
	if _, err := ctx.DB.Exec(`
	UPDATE 
	    users
	SET
	    subscribers=array_remove(subscribers, $1)
	WHERE user_id=$2`, userID, deleteUserId); err != nil {
		newerror.NewAppError("ctx.DB.Exec", err, pathToLogFile, isTimeAmPm)
		return err
	}

	if _, err := ctx.DB.Exec(`
	UPDATE 
	    users
	SET
	    subscriptions=array_remove(subscriptions, $1)
	WHERE user_id=$2`, deleteUserId, userID); err != nil {
		newerror.NewAppError("ctx.DB.Exec", err, pathToLogFile, isTimeAmPm)
		return err
	}

	return nil
}

func SelectCountSubscriber(ctx *sqlx.DB, userID int64) (isCount sql.NullInt64, err error) {
	if err := ctx.Get(&isCount, `
	SELECT
	    array_length(subscribers, 1)
	FROM
	    users
	WHERE
	    user_id=$1`, userID); err != nil {
		newerror.NewAppError("ctx.Get", err, pathToLogFile, isTimeAmPm)
		return sql.NullInt64{}, err
	}
	return isCount, nil
}
