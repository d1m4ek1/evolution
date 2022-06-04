package settings

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"iNote/www/api/authorization"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"io"
	"net/http"
	"os"
	"strings"
)

// Path to error
const (
	pathToError string = "api/settings/profile -> Function "
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

type settingsProfileData struct {
	Name           string `json:"name"`
	Logo           string `json:"logo"`
	Banner         string `json:"banner"`
	ThemePage      string `json:"themePage"`
	Language       string `json:"language"`
	AboutMeTitle   string `json:"aboutMeTitle"`
	AboutMeContent string `json:"aboutMeContent"`
	Telegram       string `json:"telegram"`
	Instagram      string `json:"instagram"`
	Facebook       string `json:"facebook"`
	Vk             string `json:"vk"`
	Tiktok         string `json:"tiktok"`
}

func getOldProfileSettings(userId string, w http.ResponseWriter) {
	var userData settingsProfileData

	if err := database.Tables.QueryRow(`SELECT name FROM users u 
		WHERE user_id=$1`, userId).Scan(&userData.Name); err != nil {
		fmt.Println(newerror.Wrap(errorGetSettings, "Query at db: 1", err))
		return
	}

	if err := database.Tables.QueryRow(`SELECT sgs.logo, sgs.banner, sgs.aboutme[1], sgs.aboutme[2], sgs.theme_page, sgs.language 
		FROM settings sgs,identifiers ids 
		WHERE ids.user_id=$1 AND ids.settings_id=sgs.settings_id`, userId).Scan(&userData.Logo, &userData.Banner,
		&userData.AboutMeTitle, &userData.AboutMeContent, &userData.ThemePage, &userData.Language); err != nil {
		fmt.Println(newerror.Wrap(errorGetSettings, "Query at db: 3", err))
		return
	}

	if err := database.Tables.QueryRow(`SELECT ctn.telegram, ctn.instagram, ctn.facebook, ctn.vk, ctn.tiktok FROM connection ctn, identifiers ids 
		WHERE ids.user_id=$1 AND ids.connection_id=ctn.connection_id`, userId).Scan(&userData.Telegram,
		&userData.Instagram, &userData.Facebook, &userData.Vk, &userData.Tiktok); err != nil {
		fmt.Println(newerror.Wrap(errorGetSettings, "Query at db: 3", err))
		return
	}

	if err := json.NewEncoder(w).Encode(userData); err != nil {
		fmt.Println(newerror.Wrap(errorGetSettings, "json", err))
		return
	}
}

func GetProfileSettings(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	if token != nil && userId != nil {
		var user authorization.CheckSignin = authorization.CheckSignin{
			Id:       userId.Value,
			Token:    token.Value,
			Autorize: false,
		}
		user.CheckSignin(&user)

		if user.Autorize {
			getOldProfileSettings(userId.Value, w)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
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
	database.Tables.QueryRow(fmt.Sprintf(`SELECT sgs.%s FROM settings sgs, identifiers ids 
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

			var idSettings string
			database.Tables.QueryRow(`SELECT settings_id FROM identifiers WHERE user_id=$1`, userId).Scan(&idSettings)
			database.Tables.Exec(fmt.Sprintf(`UPDATE settings SET %s=$1 WHERE settings_id=$2`, keyFile), filePath, idSettings)

			io.Copy(dst, src)
		}
	}
}

func saveNewSettings(userId string, r *http.Request) {
	for _, v := range keyParamFiles {
		if r.URL.Query().Get(v) != "" {
			saveImage(r, v, userId)
		}
	}
	for _, v := range keyParamSettings {
		if r.URL.Query().Get(v) != "" {
			if _, err := database.Tables.Exec(fmt.Sprintf(`UPDATE settings SET %s=$1
			FROM identifiers ids, settings sgs WHERE ids.user_id=$2
			AND sgs.settings_id=ids.settings_id;`, v), r.URL.Query().Get(v), userId); err != nil {
				fmt.Println(newerror.Wrap(errorSetSettings, "Query at db: 1", err))
			}
		}
	}
	for _, v := range keyParamConnection {
		if r.URL.Query().Get(v) != "" {
			if _, err := database.Tables.Exec(fmt.Sprintf(`UPDATE connection SET %s=$1
			FROM identifiers ids, connection ctc WHERE ids.user_id=$2
			AND ctc.connection_id=ids.connection_id;`, v), r.URL.Query().Get(v), userId); err != nil {
				fmt.Println(newerror.Wrap(errorSetSettings, "Query at db: 2", err))
			}
		}
	}
	for _, v := range keyParamUser {
		if r.URL.Query().Get(v) != "" {
			if _, err := database.Tables.Exec(fmt.Sprintf(`UPDATE users SET %s=$1 WHERE user_id=$2`, v), r.URL.Query().Get(v), userId); err != nil {
				fmt.Println(newerror.Wrap(errorSetSettings, "Query at db: 3", err))
			}
		}
	}
}

func SaveProfileSettings(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	if token != nil && userId != nil {
		var user authorization.CheckSignin = authorization.CheckSignin{
			Id:       userId.Value,
			Token:    token.Value,
			Autorize: false,
		}
		user.CheckSignin(&user)

		if user.Autorize {
			saveNewSettings(userId.Value, r)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
	}
}
