package authorization

import (
	"fmt"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
)

const (
	errorCheckSignin string = pathToError + "CheckSignin"
)

type CheckSignin struct {
	Id       string
	Token    string
	Autorize bool
}

func (c *CheckSignin) CheckSignin(user *CheckSignin) {
	if err := database.Tables.QueryRow(`SELECT count(*) <> 0 FROM users_data 
	WHERE id=$1 AND token=$2;`, user.Id, user.Token).Scan(&c.Autorize); err != nil {
		fmt.Println(newerror.Wrap(errorCheckSignin, "Query at db: 1", err))
	}
}
