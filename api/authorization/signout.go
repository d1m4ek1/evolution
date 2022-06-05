package authorization

import (
	"fmt"
	"net/http"

	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
)

const (
	errorSignOut string = pathToError + "SignOut"
)

func SignOut(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("token")

	if token != nil {
		var userID string
		if err := database.Tables.QueryRow(database.SelectUserIdForToken, token.Value).Scan(&userID); err != nil {
			fmt.Println(newerror.Wrap(errorSignOut, "Query at db: 1", err))
		}
		if _, err := database.Tables.Exec(database.UpdateNetworkStatusOffline, userID); err != nil {
			fmt.Println(newerror.Wrap(errorSignOut, "Query at db: 2", err))
		}
		if _, err := database.Tables.Exec(database.UpdateTokenNull, userID); err != nil {
			fmt.Println(newerror.Wrap(errorSignOut, "Query at db: 2", err))
		}
	}
}
