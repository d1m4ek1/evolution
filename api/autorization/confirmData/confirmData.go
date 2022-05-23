package confirmdata

import (
	"encoding/json"
	"fmt"
	checksignin "iNote/www/api/autorization/checkSignin"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"net/http"
)

// Path to error
const (
	pathToError string = "api/autorization/confirmData -> Function "
)

const (
	errorConfPass        string = pathToError + "confPass"
	errorConfirmPassword string = pathToError + "ConfirmPassword"
)

type confirmitedPassword struct {
	Password bool `json:"pass"`
}

func (c *confirmitedPassword) confPass(i, t, p string) {
	if err := database.Tables.QueryRow(`SELECT count(*) = 1 FROM users_data 
	WHERE id=$1 AND token=$2 AND password=$3`, i, t, p).Scan(&c.Password); err != nil {
		fmt.Println(newerror.Wrap(errorConfPass, "Query at db: 1", err))
	}
}

func ConfirmPassword(w http.ResponseWriter, r *http.Request) {
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
			confPass := confirmitedPassword{}
			confPass.confPass(userId.Value, token.Value, r.URL.Query().Get("conf_pass"))

			if err := json.NewEncoder(w).Encode(confPass); err != nil {
				fmt.Println(newerror.Wrap(errorConfirmPassword, "json", err))
				return
			}

		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		}
	}
}
