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

	database.Tables.QueryRow(database.SelectIdForIdOrCustomId, id, id).Scan(&user)

	if user != "" {
		if err := database.Tables.QueryRow(database.SelectProfileData, id).Scan(&pdd.Logo, &pdd.Banner, &pdd.Name, pq.Array(&pdd.Position), pq.Array(&da.Audience), &pdd.Verif, &pdd.NetworkStatus); err != nil {
			fmt.Println(newerror.Wrap(errorProfileDefault, "Query at db: 2", err))
		}

		pdd.Audience = len(da.Audience)

		return pdd, user
	}

	return general.ProfileData{}, ""
}

func GetUserDataStatic(token, _ *http.Cookie, userUrlId string) general.ProfileData {
	profileDefaultData, newUserId := profileDefault(userUrlId)

	if token != nil {
		if newUserId != "" && token.Value != "" {
			profileDefaultData.ProfileUser(newUserId, token.Value)
		}
	}

	return profileDefaultData
}
