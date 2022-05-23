package settings

import (
	"encoding/json"
	"fmt"
	checksignin "iNote/www/api/autorization/checkSignin"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"net/http"

	"github.com/lib/pq"
)

// Path to error
const (
	pathToErrorPD string = "api/settings/personalData -> Function "
)

const (
	errorGetOldPersonalData string = pathToErrorPD + "getOldPersonalData"
	errorGetPersonalData    string = pathToErrorPD + "GetPersonalData"
	errorsavePersonalData   string = pathToErrorPD + "savePersonalData"
)

type validBackupKeys struct {
	BackupKeys []string
}

type personalData struct {
	BackupKeys bool   `json:"bcpk"`
	Email      string `json:"eml"`
}

func (v *personalData) validBackupKey(arr []string) {
	if len(arr) != 0 && len(arr) < 5 {
		v.BackupKeys = true
	} else {
		v.BackupKeys = false
	}
}

func getOldPersonalData(userId string, w http.ResponseWriter) {
	var userData personalData
	var backupKeysArr validBackupKeys

	if err := database.Tables.QueryRow(`SELECT email, backup_keys FROM users_data 
		WHERE id=$1`, userId).Scan(&userData.Email, pq.Array(&backupKeysArr.BackupKeys)); err != nil {
		fmt.Println(newerror.Wrap(errorGetOldPersonalData, "Query at db: 1", err))
		return
	}

	userData.validBackupKey(backupKeysArr.BackupKeys)

	if err := json.NewEncoder(w).Encode(userData); err != nil {
		fmt.Println(newerror.Wrap(errorGetOldPersonalData, "json", err))
		return
	}
}

func GetPersonalData(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	if token != nil && userId != nil {
		var user checksignin.CheckSignin = checksignin.CheckSignin{
			Id:       userId.Value,
			Token:    token.Value,
			Autorize: false,
		}
		user.CheckSignin(&user)

		if user.Autorize {
			getOldPersonalData(userId.Value, w)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}
	}
}

func savePersonalData(userId string, r *http.Request) {
	var pdParams = []string{"backupkey_one", "backupkey_two", "backupkey_three", "backupkey_four", "new_password", "email"}
	var backupKey string

	if r.URL.Query().Get(pdParams[0]) != "" && r.URL.Query().Get(pdParams[1]) != "" &&
		r.URL.Query().Get(pdParams[2]) != "" && r.URL.Query().Get(pdParams[3]) != "" {

		backupKey = fmt.Sprintf("{%s, %s, %s, %s}",
			r.URL.Query().Get(pdParams[0]),
			r.URL.Query().Get(pdParams[1]),
			r.URL.Query().Get(pdParams[2]),
			r.URL.Query().Get(pdParams[3]))

		if _, err := database.Tables.Exec(`UPDATE users_data SET backup_keys=$1 WHERE id=$2`, backupKey, userId); err != nil {
			fmt.Println(newerror.Wrap(errorsavePersonalData, "Query at db: 1", err))
			return
		}
	}
}

func SavePersonalData(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	if token != nil && userId != nil {
		var user checksignin.CheckSignin = checksignin.CheckSignin{
			Id:       userId.Value,
			Token:    token.Value,
			Autorize: false,
		}
		user.CheckSignin(&user)

		if user.Autorize {
			savePersonalData(userId.Value, r)
		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}
	}
}
