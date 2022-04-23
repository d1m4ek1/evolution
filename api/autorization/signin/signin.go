package signin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"iNote/www/pkg/general"
)

// Path to error
const (
	pathToError string = "api/autorization/signin -> Function "
)

const (
	errorSignIn  string = pathToError + "SignIn"
	errorGetData string = pathToError + "getData"
)

func getDataSignIn(s *general.SignInData, w http.ResponseWriter) (general.UserData, bool) {

	data := general.UserData{}

	err := database.Tables.QueryRow(database.SelectUserDataSignIn, s.Login, s.Password).Scan(&data.UserId, &data.NetworkStatus, &s.OldToken)
	if err != nil {
		e := newerror.ErrorClient{Value: "Неверный логин или пароль", Number: 200}
		if err := json.NewEncoder(w).Encode(&e); err != nil {
			fmt.Println(newerror.Wrap(errorGetData, "Json: 1", err))
		}
		return general.UserData{}, false
	}

	if !s.OldToken.Valid {
		if _, err := database.Tables.Exec(database.UpdateTokenNotNull, s.NewToken, data.UserId); err != nil {
			fmt.Println(newerror.Wrap(errorGetData, "Query at db: 3", err))
		}
	} else {
		data.OldToken = s.OldToken.String
	}

	if data.NetworkStatus == "offline" {
		if _, err := database.Tables.Exec(database.UpdateNetworkStatusOnline, data.UserId); err != nil {
			fmt.Println(newerror.Wrap(errorGetData, "Query at db: 2", err))
		}
	}

	return data, true
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	signIn := r.URL.Query().Get("signin")

	if signIn == "true" {

		userData, bo := getDataSignIn(&general.SignInData{
			Login:    r.URL.Query().Get("login"),
			Password: r.URL.Query().Get("password"),
			NewToken: r.URL.Query().Get("token"),
		}, w)

		if bo {
			if err := json.NewEncoder(w).Encode(&userData); err != nil {
				fmt.Println(newerror.Wrap(errorSignIn, "JSON", err))
			}
		}
	}
}
