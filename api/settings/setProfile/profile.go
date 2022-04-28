package setprofile

import (
	"encoding/json"
	"fmt"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"iNote/www/pkg/general"
	"net/http"
)

// Path to error
const (
	pathToError string = "api/settings/setProfile -> Function "
)

const (
	errorGetSettings string = pathToError + "GetSettings"
	errorSetSettings string = pathToError + "SetSettings"
)

func GetSettings(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	if token != nil && userId != nil {
		var userData general.SettingsProfileData

		if err := database.Tables.QueryRow(`SELECT name FROM users u 
		WHERE user_id=$1`, userId.Value).Scan(&userData.Name); err != nil {
			fmt.Println(newerror.Wrap(errorGetSettings, "Query at db: 1", err))
		}

		if err := database.Tables.QueryRow(`SELECT sgs.logo, sgs.banner, sgs.aboutme[1], sgs.aboutme[2], sgs.theme_page, sgs.language 
		FROM settings sgs,identifiers ids 
		WHERE ids.user_id=$1 AND ids.settings_id=sgs.settings_id`, userId.Value).Scan(&userData.Logo, &userData.Banner,
			&userData.AboutMeTitle, &userData.AboutMeContent, &userData.ThemePage, &userData.Language); err != nil {
			fmt.Println(newerror.Wrap(errorGetSettings, "Query at db: 3", err))
		}

		if err := database.Tables.QueryRow(`SELECT ctn.telegram, ctn.instagram, ctn.facebook, ctn.vk, ctn.tiktok FROM connection ctn, identifiers ids 
		WHERE ids.user_id=$1 AND ids.connection_id=ctn.connection_id`, userId.Value).Scan(&userData.Telegram,
			&userData.Instagram, &userData.Facebook, &userData.Vk, &userData.Tiktok); err != nil {
			fmt.Println(newerror.Wrap(errorGetSettings, "Query at db: 3", err))
		}

		if err := json.NewEncoder(w).Encode(userData); err != nil {
			fmt.Println(newerror.Wrap(errorGetSettings, "json", err))
		}
	}
}

func SetSettings(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	var keyParamSettings = []string{"name", "language", "themePage", "logo", "banner", "aboutme_title",
		"aboutme_content"}
	var keyParamConnection = []string{"telegram", "instagram", "facebook", "vk", "tiktok"}

	if token != nil && userId != nil {
		for _, v := range keyParamSettings {
			if r.URL.Query().Get(v) != "" {
				if _, err := database.Tables.Exec(fmt.Sprintf(`UPDATE settings SET %s=$1 
				FROM identifiers WHERE identifiers.user_id=$2 
				AND settings.settings_id=identifiers.settings_id;`, v), r.URL.Query().Get(v), userId.Value); err != nil {
					fmt.Println(newerror.Wrap(errorSetSettings, "Query at db: 1", err))
				}
			}
		}
		for _, v := range keyParamConnection {
			if r.URL.Query().Get(v) != "" {
				if _, err := database.Tables.Exec(fmt.Sprintf(`UPDATE connection SET %s=$1 
				FROM identifiers WHERE identifiers.user_id=$2 
				AND connection.connection_id=identifiers.connection_id;`, v), r.URL.Query().Get(v), userId.Value); err != nil {
					fmt.Println(newerror.Wrap(errorSetSettings, "Query at db: 1", err))
				}
			}
		}
	}
}
