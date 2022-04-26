package signup

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

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

func completeTheRestTables(id int, s general.SignUpData) error {
	if _, err := database.Tables.Exec(database.InsertNewUser, id, s.Nickname); err != nil {
		fmt.Println(errorCreateAccount, "Query at db: 1", err)
		return fmt.Errorf("<- completeTheRestTables")
	}

	var connectionId string = setConnectionIdentificate(id)
	var settingsId string = setSettingsIdentificate(id)

	if _, err := database.Tables.Exec(database.InsertIdentifiers, id, connectionId, settingsId); err != nil {
		fmt.Println(errorCreateAccount, "Query at db: 2", err)
		return fmt.Errorf("<- completeTheRestTables")
	}

	if _, err := database.Tables.Exec(database.InsertConnection, connectionId); err != nil {
		fmt.Println(errorCreateAccount, "Query at db: 3", err)
		return fmt.Errorf("<- completeTheRestTables")
	}
	if _, err := database.Tables.Exec(database.InsertSettings, settingsId); err != nil {
		fmt.Println(errorCreateAccount, "Query at db: 3", err)
		return fmt.Errorf("<- completeTheRestTables")
	}

	return nil
}

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
		var identificate int
		err := database.Tables.QueryRow(database.InsertNewUserData, s.Login, s.Password, s.Email, s.Token).Scan(&identificate)
		if err != nil {
			fmt.Println(newerror.Wrap(errorCreateAccount, "Query at db: 1", err))
		}

		if err := completeTheRestTables(identificate, s); err != nil {
			fmt.Println(newerror.Wrap(errorCreateAccount, "resInserted", err))
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
