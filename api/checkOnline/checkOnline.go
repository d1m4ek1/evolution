package checkonline

import (
	"fmt"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"net/http"
	"time"
)

// Path to error
const (
	pathToError string = "api/checkOnline -> Function "
)

const (
	errorCheckOnline string = pathToError + "CheckOnline"
)

func CheckOnline(_ http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")
	if token != nil && userId != nil {
		if token.Value != "" && userId.Value != "" {

			var user string

			database.Tables.QueryRow(database.SelectLoginForIdToken, userId.Value, token.Value).Scan(&user)

			if user != "" {
				if _, err := database.Tables.Exec(database.UpdateNetworkStatusOnline, userId.Value); err != nil {
					fmt.Println(newerror.Wrap(errorCheckOnline, "Query at db: 2", err))
				}

				t := time.NewTimer(10 * time.Second)
				<-t.C
				if _, err := database.Tables.Exec(database.UpdateNetworkStatusOffline, userId.Value); err != nil {
					fmt.Println(newerror.Wrap(errorCheckOnline, "Query at db: 3", err))
				}
			}
		}
	}
}
