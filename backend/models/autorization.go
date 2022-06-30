package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	newerror "iNote/www/backend/pkg/NewError"
	"math/rand"
	"time"
)

type CheckSignin struct {
	Id       int64
	Token    string
	Autorize bool
}

type ConfirmitadePassword struct {
	Password bool `json:"pass"`
}

func (c *CheckSignin) CheckUserOnSignin(ctx *sqlx.DB) error {
	if err := ctx.Get(&c.Autorize, `
		SELECT 
		    count(*) <> 0 
		FROM 
		    users_data 
		WHERE 
		    id=$1 
		  AND 
		    token=$2;`, c.Id, c.Token); err != nil {
		newerror.Wrap("ctx.Get", err)
		return err
	}
	return nil
}

func (c *ConfirmitadePassword) ConfirmPassword(ctx *sqlx.DB, userId int64, token, confPass string) error {
	if err := ctx.Get(&c.Password, `
		SELECT 
		    count(*) = 1 
		FROM 
		    users_data 
		WHERE 
		    id=$1 
		  AND 
		    token=$2 
		  AND 
		    password=$3`, userId, token, confPass); err != nil {
		newerror.Wrap("ctx.Get", err)
		return err
	}
	return nil
}

func SignInData(ctx *sqlx.DB, login, password, newToken string) (userID int64, netStatus string, err error) {
	isVerify, err := VerifUser(ctx, login, password)
	if err != nil {
		newerror.Wrap("VerifUser", err)
		return 0, "", nil
	}

	if isVerify {
		if err := ctx.DB.QueryRow(`
			SELECT
			    ud.id,
			    u.net_status
			FROM
			    users_data AS ud,
					users AS u
			WHERE
			    ud.login=$1
			  AND
			    ud.password=$2
				AND
			    ud.id=u.id`, login, password).Scan(&userID, &netStatus); err != nil {
			newerror.Wrap("ctx.DB.QueryRow", err)
			return 0, "", nil
		}

		if userID != 0 && newToken != "" {
			if err := SetNewToken(ctx, userID, newToken); err != nil {
				newerror.Wrap("SetNewToken", err)
				return 0, "", nil
			}

			if err := SetNetworkStatusOnline(ctx, userID); err != nil {
				newerror.Wrap("SetNetworkStatusOnline", err)
				return 0, "", err
			}
		}
	}

	return userID, netStatus, nil
}

func SetNewToken(ctx *sqlx.DB, userID int64, newToken string) error {
	if _, err := ctx.DB.Exec(`
		UPDATE 
		    users_data 
		SET 
		    token=$1 
		WHERE 
		    users_data.id=$2`, newToken, userID); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}

func SetNetworkStatusOnline(ctx *sqlx.DB, userID int64) error {
	if _, err := ctx.DB.Exec(`
		UPDATE 
		    users 
		SET 
		    net_status='online' 
		WHERE 
		    user_id=$1 
		  AND 
		    net_status<>'online'`, userID); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}

func SetNetworkStatusOffline(ctx *sqlx.DB, userID int64) error {
	if _, err := ctx.DB.Exec(`
		UPDATE 
		    users 
		SET 
		    net_status=DEFAULT 
		WHERE 
		    id=$1 
		  AND 
		    net_status<>'offline'`, userID); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}

func UserSignOut(ctx *sqlx.DB, userId string) error {
	if _, err := ctx.DB.Exec(`
		UPDATE 
		    users 
		SET 
		    net_status=DEFAULT 
		WHERE 
		    id=$1 
		  AND 
		    net_status<>'offline'`, userId); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}

	if _, err := ctx.DB.Exec(`
		UPDATE 
		    users_data 
		SET 
		    token=DEFAULT 
		WHERE 
		    id=$1`, userId); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}

func CheckLogin(ctx *sqlx.DB, login string) (bool, error) {
	var user bool

	if err := ctx.Get(&user, `
		SELECT 
		    count(*) = 1 
		FROM 
		    users_data 
		WHERE 
		    login=$1`, login); err != nil {
		newerror.Wrap("ctx.Get", err)
		return user, err
	}
	return user, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func setConnectionIdentificate(id int) string {
	var length int = 63

	rand.Seed(time.Now().UnixNano())

	var token = make([]rune, length)

	for i := range token {
		token[i] = letters[rand.Intn(len(letters))]
	}

	return string(token) + fmt.Sprint(id)
}

func setSettingsIdentificate(id int) string {
	var length int = 63

	rand.Seed(time.Now().UnixNano())

	var token = make([]rune, length)

	for i := range token {
		token[i] = letters[rand.Intn(len(letters))]
	}

	return string(token) + fmt.Sprint(id)
}

func CreateAccount(ctx *sqlx.DB, login, password, email, token, nickname string) error {
	var identificate int
	if err := ctx.Get(&identificate, `
		INSERT INTO 
		    users_data (login, password, email, token) 
		VALUES 
		    ($1, $2, $3, $4) 
		RETURNING 
		    id`, login, password, email, token); err != nil {
		newerror.Wrap("ctx.Get", err)
		return err
	}

	if _, err := ctx.DB.Exec(`INSERT INTO users (user_id, name) VALUES ($1, $2)`, identificate, nickname); err != nil {
		newerror.Wrap("ctx.Get", err)
		return err
	}

	var connectionId string = setConnectionIdentificate(identificate)
	var settingsId string = setSettingsIdentificate(identificate)

	if _, err := ctx.DB.Exec(`
		INSERT INTO 
		    identifiers (user_id, connection_id, settings_id) 
		VALUES 
		    ($1, $2, $3)`,
		identificate,
		connectionId,
		settingsId); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}

	if _, err := ctx.DB.Exec(`
		INSERT INTO 
		    connection (connection_id) 
		VALUES 
		    ($1)`,
		connectionId); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}

	if _, err := ctx.DB.Exec(`
		INSERT INTO 
		    settings (settings_id) 
		VALUES 
		    ($1)`,
		settingsId); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}

	return nil
}
