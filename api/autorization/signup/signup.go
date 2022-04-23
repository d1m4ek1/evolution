package signup

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
	pathToError string = "api/autorization/signup -> Function "
)

const (
	errorSignUp        string = pathToError + "SignUp"
	errorCreateAccount string = pathToError + "createAccount"
)

func createAccount(s general.SignUpData, w http.ResponseWriter) {
	var user string

	database.Tables.QueryRow(database.SelectLogin, s.Login).Scan(&user)
	if user != "" {
		e := newerror.ErrorClient{Value: "Логин занят!", Number: 200}

		if err := json.NewEncoder(w).Encode(&e); err != nil {
			fmt.Println(newerror.Wrap(errorSignUp, "Json", err))
		}
		return
	}

	if user == "" {
		if _, err := database.Tables.Exec(database.InsertNewUserData, s.Login, s.Password, s.Email, s.Token); err != nil {
			fmt.Println(newerror.Wrap(errorCreateAccount, "Query at db: 1", err))
		}

		if _, err := database.Tables.Exec(database.InsertNewUser, s.Nickname); err != nil {
			fmt.Println(newerror.Wrap(errorCreateAccount, "Query at db: 2", err))
		}

		type AutBool struct {
			Aut bool `json:"aut"`
		}

		aut := AutBool{Aut: true}

		if err := json.NewEncoder(w).Encode(&aut); err != nil {
			fmt.Println(newerror.Wrap(errorCreateAccount, "Json", err))
		}
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	keyWords := [5]string{"nickname", "email", "login", "password", "token"}
	var valid int

	for _, v := range keyWords {
		if r.URL.Query().Get(v) != "" {
			valid += 1
		}
	}

	if valid == 5 {
		signUpData := general.SignUpData{}
		signUpData.ValidData(&general.SignUpData{
			Nickname: r.URL.Query().Get("nickname"),
			Email:    r.URL.Query().Get("email"),
			Login:    r.URL.Query().Get("login"),
			Password: r.URL.Query().Get("password"),
			Token:    r.URL.Query().Get("token"),
		})

		if signUpData.Login != "" && signUpData.Password != "" && signUpData.Nickname != "" && signUpData.Email != "" && signUpData.Token != "" {
			createAccount(signUpData, w)
		} else {
			e := newerror.ErrorClient{Value: "Некорректное значение", Number: 200}

			if err := json.NewEncoder(w).Encode(&e); err != nil {
				fmt.Println(newerror.Wrap(errorSignUp, "Json", err))
			}
		}
	} else {
		e := newerror.ErrorClient{Value: "Получено пустое поле", Number: 200}

		if err := json.NewEncoder(w).Encode(&e); err != nil {
			fmt.Println(newerror.Wrap(errorSignUp, "Json", err))
		}
	}
}
