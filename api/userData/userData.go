package userdata

import (
	"fmt"
	"net/http"

	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"iNote/www/pkg/general"

	"github.com/lib/pq"
)

// Path to error
const (
	pathToError string = "api/userData/userdata -> Function "
)

const (
	errorProfileDefault string = pathToError + "profileDefault"
)

func profileDefault(id string) (general.ProfileData, string) {
	pdd := general.ProfileData{}
	da := general.DataArray{}
	var user string

	database.Tables.QueryRow(`SELECT id FROM users_data WHERE id=$1 OR user_custom_id=$2`, id, id).Scan(&user)

	if user != "" {
		if err := database.Tables.QueryRow(`SELECT logo, banner, name, position, audience, verification, network_status 
		FROM users WHERE id=$1`, id).Scan(&pdd.Logo, &pdd.Banner, &pdd.Name, pq.Array(&pdd.Position), pq.Array(&da.Audience), &pdd.Verif, &pdd.NetworkStatus); err != nil {
			fmt.Println(newerror.Wrap(errorProfileDefault, "Query at db: 2", err))
		}

		pdd.Audience = len(da.Audience)

		return pdd, user
	}

	return general.ProfileData{}, ""
}

func GetUserDataStatic(token, userId *http.Cookie, userUrlId string) general.ProfileData {
	profileDefaultData, newUserId := profileDefault(userUrlId)

	if token != nil {
		if newUserId != "" && token.Value != "" {
			profileDefaultData.ProfileUser(newUserId, token.Value)
		}
	}

	return profileDefaultData
}
