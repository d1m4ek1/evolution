package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	newerror "iNote/www/backend/pkg/NewError"
)

type SettingsProfileData struct {
	Name           string `json:"name" db:"name"`
	Logo           string `json:"logo" db:"logo"`
	Banner         string `json:"banner" db:"banner"`
	ThemePage      string `json:"themePage" db:"theme_page"`
	Language       string `json:"language" db:"language"`
	AboutMeTitle   string `json:"aboutMeTitle" db:"about_me_title"`
	AboutMeContent string `json:"aboutMeContent" db:"about_me_content"`
	Telegram       string `json:"telegram" db:"telegram"`
	Instagram      string `json:"instagram" db:"instagram"`
	Facebook       string `json:"facebook" db:"facebook"`
	Vk             string `json:"vk" db:"vk"`
	Tiktok         string `json:"tiktok" db:"tiktok"`
}

func SetBackupKeys(ctx *sqlx.DB, tplKeys string, userID int64) error {
	if _, err := ctx.DB.Exec(`
		UPDATE 
		    users_data 
		SET 
		    backup_keys=$1 
		WHERE 
		    id=$2`, tplKeys, userID); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}

func SetPersonalData(ctx *sqlx.DB, tplPersonal string, userID int64) error {
	if _, err := ctx.DB.Exec(fmt.Sprintf(`
		UPDATE 
				users_data 
		SET 
				%s 
		WHERE 
				id=$1`, tplPersonal), userID); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}

func SelectPersonalData(ctx *sqlx.DB, userID int64) (email string, backupKeys []string, err error) {
	if err := ctx.DB.QueryRow(`
		SELECT 
		    email, 
		    backup_keys 
		FROM 
		    users_data 
		WHERE 
		    id=$1`, userID).Scan(&email, pq.Array(&backupKeys)); err != nil {
		newerror.Wrap("ctx.DB.QueryRow", err)
		return "", nil, err
	}
	return email, backupKeys, nil
}

func SetSettingsProfile(ctx *sqlx.DB, tplSettings, tplConnection, tplUser string, userID int64) error {
	if tplSettings != "" {
		if _, err := ctx.DB.Exec(fmt.Sprintf(`
		UPDATE 
				settings 
		SET 
				%s 
		WHERE 
				settings_id=(SELECT settings_id FROM identifiers WHERE user_id=$1)`, tplSettings), userID); err != nil {
			newerror.Wrap("ctx.DB.Exec", err)
			return err
		}
	}

	if tplConnection != "" {
		if _, err := ctx.DB.Exec(fmt.Sprintf(`
		UPDATE 
				connection 
		SET 
				%s 
		WHERE 
				connection_id=(SELECT connection_id FROM identifiers WHERE user_id=$1)`, tplConnection), userID); err != nil {
			newerror.Wrap("ctx.DB.Exec", err)
			return err
		}
	}

	if tplUser != "" {
		if _, err := ctx.DB.Exec(fmt.Sprintf(`
		UPDATE 
				users 
		SET 
				%s 
		WHERE 
				user_id=$1`, tplUser), userID); err != nil {
			newerror.Wrap("ctx.DB.Exec", err)
			return err
		}
	}

	return nil
}

func SelectFilePath(ctx *sqlx.DB, keyFile string, userID int64) (path string, err error) {
	if err := ctx.Get(&path, fmt.Sprintf(`
	SELECT 
			sgs.%s 
	FROM 
			settings sgs, 
			identifiers ids 
	WHERE 
			ids.user_id=$1 
		AND 
			ids.settings_id=sgs.settings_id`, keyFile), userID); err != nil {
		newerror.Wrap("ctx.Get", err)
		return "", err
	}
	return path, nil
}

func SetFilePath(ctx *sqlx.DB, keyFile, filePath string, userID int64) error {
	if _, err := ctx.DB.Exec(fmt.Sprintf(`
		UPDATE 
				settings 
		SET 
				%s=$1
		WHERE 
			settings_id=(SELECT settings_id FROM identifiers WHERE user_id=$2)`, keyFile), filePath, userID); err != nil {
		newerror.Wrap("ctx.DB.Exec", err)
		return err
	}
	return nil
}

func SelectProfileSettings(ctx *sqlx.DB, userID int64) (spd SettingsProfileData, err error) {
	if err := ctx.Get(&spd, `
		SELECT 
		    name 
		FROM 
		    users 
		WHERE 
		    user_id=$1`, userID); err != nil {
		newerror.Wrap("ctx.Get", err)
		return SettingsProfileData{}, err
	}

	if err := ctx.Get(&spd, `
		SELECT 
		    t1.logo AS logo, 
		    t1.banner AS banner, 
		    t1.aboutme[1] AS about_me_title, 
		    t1.aboutme[2] AS about_me_content, 
		    t1.theme_page AS theme_page, 
		    t1.language AS language
		FROM 
		    settings t1,
		    identifiers t2 
		WHERE 
		    t2.user_id=$1 
		  AND 
		    t2.settings_id=t1.settings_id`, userID); err != nil {
		newerror.Wrap("ctx.Get", err)
		return SettingsProfileData{}, err
	}

	if err := ctx.Get(&spd, `
		SELECT 
		    t1.telegram AS telegram, 
		    t1.instagram AS instagram, 
		    t1.facebook AS facebook, 
		    t1.vk AS vk, 
		    t1.tiktok AS tiktok
		FROM 
		    connection t1, 
		    identifiers t2 
		WHERE 
		    t2.user_id=$1 
		  AND 
		    t2.connection_id=t1.connection_id`, userID); err != nil {
		newerror.Wrap("ctx.Get", err)
		return SettingsProfileData{}, err
	}

	return spd, nil
}
