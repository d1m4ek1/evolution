package setprofile

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"iNote/www/pkg/general"
	"io"
	"net/http"
	"os"
	"strings"
)

// Path to error
const (
	pathToError string = "api/settings/setProfile -> Function "
)

const (
	errorGetSettings string = pathToError + "GetSettings"
	errorSetSettings string = pathToError + "SetSettings"
)

var keyParamFiles = [2]string{"logo", "banner"}
var keyParamUser = [1]string{"name"}
var keyParamSettings = [5]string{"language", "themePage", "aboutme_title",
	"aboutme_content"}
var keyParamConnection = [5]string{"telegram", "instagram", "facebook", "vk", "tiktok"}

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

func exists(s string, key string) bool {
	if _, err := os.Stat("./profile/" + key + "/saved/" + s); !os.IsNotExist(err) {
		return false
	}
	return true
}

func storee(bv []byte) string {
	hasher := sha1.New()
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}

func validFile(s, id string) (string, error) {
	var filePath string

	if strings.Contains(s, ".png") {
		sha := storee([]byte(strings.Replace(s, ".png", "", -1)))
		filePath = "id" + id + "_" + sha + ".png"

	} else if strings.Contains(s, ".jpeg") {
		sha := storee([]byte(strings.Replace(s, ".jpeg", "", -1)))
		filePath = "id" + id + "_" + sha + ".jpeg"

	} else if strings.Contains(s, ".jpg") {
		sha := storee([]byte(strings.Replace(s, ".jpg", "", -1)))
		filePath = "id" + id + "_" + sha + ".jpg"
	} else if strings.Contains(s, ".gif") {
		sha := storee([]byte(strings.Replace(s, ".gif", "", -1)))
		filePath = "id" + id + "_" + sha + ".gif"
	} else {
		return "", errors.New("Файл не является изображением")
	}

	return filePath, nil
}

func saveImage(r *http.Request, keyFile string, userId string) {
	var oldFilePath string
	database.Tables.QueryRow(fmt.Sprintf(`SELECT sgs.%s FROM settings sgs,identifiers ids 
	WHERE ids.user_id=$1 AND ids.settings_id=sgs.settings_id`, keyFile), userId).Scan(&oldFilePath)

	if r.Method == "POST" {
		src, hdr, err := r.FormFile(keyFile)
		if err != nil {
			fmt.Println(err, errors.New("1"))
			return
		}
		defer src.Close()

		if exists(oldFilePath, keyFile) {
			os.Remove("profile/" + keyFile + "/saved/" + oldFilePath)
		}

		filePath, err := validFile(hdr.Filename, userId)
		if err != nil {
			fmt.Println(err, errors.New("3"))
			return
		}

		if filePath != "" {
			filePath = strings.Replace(filePath, "=", "", -1)

			dst, err := os.Create("profile/" + keyFile + "/saved/" + filePath)
			if err != nil {

				fmt.Println(err, errors.New("4"))
				return
			}
			defer dst.Close()

			database.Tables.Exec(fmt.Sprintf(`UPDATE settings SET %s=$1 FROM identifiers WHERE 
			identifiers.user_id=$2 AND identifiers.settings_id=settings.settings_id;`, keyFile), filePath, userId)

			io.Copy(dst, src)
		}
	}
}

func SaveSettings(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	if token != nil && userId != nil {
		for _, v := range keyParamFiles {
			if r.URL.Query().Get(v) != "" {
				saveImage(r, v, userId.Value)
			}
		}
		for _, v := range keyParamSettings {
			if r.URL.Query().Get(v) != "" {
				if _, err := database.Tables.Exec(fmt.Sprintf(`UPDATE users SET %s=$1
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
					fmt.Println(newerror.Wrap(errorSetSettings, "Query at db: 2", err))
				}
			}
		}
		for _, v := range keyParamUser {
			if r.URL.Query().Get(v) != "" {
				if _, err := database.Tables.Exec(fmt.Sprintf(`UPDATE users SET %s=$1 WHERE user_id=$2`, v), r.URL.Query().Get(v), userId.Value); err != nil {
					fmt.Println(newerror.Wrap(errorSetSettings, "Query at db: 3", err))
				}
			}
		}
	}
}