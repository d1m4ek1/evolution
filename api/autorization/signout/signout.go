package signout

import (
	"fmt"
	"net/http"

	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
)

const (
	pathToError string = "api/autorization/signout -> Function "
)

const (
	errorSignOut string = pathToError + "SignOut"
)

func SignOut(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("token")

	if token != nil {
		var userID string
		if err := database.Tables.QueryRow(`SELECT ud.id FROM users_data ud WHERE token=$1`, token.Value).Scan(&userID); err != nil {
			fmt.Println(newerror.Wrap(errorSignOut, "Query at db: 1", err))
		}
		if _, err := database.Tables.Exec(`UPDATE users SET network_status='offline' WHERE id=$1`, userID); err != nil {
			fmt.Println(newerror.Wrap(errorSignOut, "Query at db: 2", err))
		}
		if _, err := database.Tables.Exec(`UPDATE users_data SET token=null WHERE id=$1`, userID); err != nil {
			fmt.Println(newerror.Wrap(errorSignOut, "Query at db: 2", err))
		}
	}
}
