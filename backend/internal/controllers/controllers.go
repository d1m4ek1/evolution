package controllers

import (
	"iNote/www/backend/pkg/general"
	"iNote/www/backend/pkg/newerror"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const pathToLogFile string = "backend/logs/logs.txt"
const isTimeAmPm bool = true

// Path to main templates
const (
	homeTpl          string = "./ui/html/index.html"
	autorizTpl       string = "./ui/html/Authoriz.html"
	profileTpl       string = "./ui/html/Profile.html"
	directoryTpl     string = "./ui/html/Directory.html"
	settingsTpl      string = "./ui/html/Settings.html"
	subscriptionsTpl string = "./ui/html/Subscriptions.html"
	ordersTpl        string = "./ui/html/Orders.html"

	// Projects
	insocialTpl string = "./ui/html/inSocial.html"
	inmusicTpl  string = "./ui/html/inMusic.html"
	inbeatsTpl  string = "./ui/html/inBeats.html"
)

// TEMPLATES
const (
	headerTmpl       string = "./ui/templates/Headers/Header.layout.html"
	stickyheaderTmpl string = "./ui/templates/Headers/StickyHeader.layout.html"
)

type ReplyBaseOnToken struct {
	Define        string
	UserUrlId     string
	Token, UserId string
	Profile       general.ProfileData
	Settings      general.SettingsData
	Context       *gin.Context
}

type Redirector struct {
	Token, UserId string
}

func replyBasedOnToken(ctx *sqlx.DB, r *ReplyBaseOnToken) {
	data := defineHeaderForAutorize(ctx, r.Token)

	if r.Token != "" {
		if data.UserId == "" && !data.Auth {
			r.Context.HTML(http.StatusOK, r.Define, nil)
			return

		} else if data.UserId == r.UserId && data.Auth {
			data.Auth = true
			data.UserUrlId = r.UserUrlId

			Data := struct {
				HeaderData   general.HeaderData
				ProfileData  general.ProfileData
				SettingsData general.SettingsData
			}{
				HeaderData:   data,
				ProfileData:  r.Profile,
				SettingsData: r.Settings,
			}
			r.Context.HTML(http.StatusOK, r.Define, Data)
		}
	} else {
		Data := struct {
			HeaderData  general.HeaderData
			ProfileData general.ProfileData
		}{
			HeaderData:  data,
			ProfileData: r.Profile,
		}
		r.Context.HTML(http.StatusOK, r.Define, Data)
	}
}

// defineHeaderForAutorize определяет шапку страницы на авторизованную и не авторизованную
func defineHeaderForAutorize(ctx *sqlx.DB, token string) general.HeaderData {
	var headerData general.HeaderData
	if token != "" {
		if err := ctx.DB.QueryRow(`
			SELECT 
			    ud.id, 
			    ud.user_custom_id 
			FROM 
			    users_data ud 
			WHERE 
			    ud.token=$1`, token).Scan(&headerData.UserId, &headerData.CustomId); err != nil {
			newerror.NewAppError("ctx.DB.QueryRow", err, pathToLogFile, isTimeAmPm)
			return general.HeaderData{}
		}

		headerData.Auth = true

		if headerData.CustomId.Valid {
			headerData.UserId = headerData.CustomId.String
		}
		if headerData.UserId != "" {
			return headerData
		}
	}

	headerData.Auth = false
	return headerData
}

func redirectUser(ctx *sqlx.DB, context *gin.Context, redirector *Redirector) {
	if redirector.Token == "" && redirector.UserId == "" {
		context.Redirect(http.StatusFound, "/signin")
		return
	} else {
		if !general.ValidateUser(ctx, redirector.Token, redirector.UserId) {
			context.SetCookie("token", "", -1, "/", "localhost", false, true)
			context.SetCookie("userId", "", -1, "/", "localhost", false, true)
			context.Redirect(http.StatusFound, "/signin")
		}
	}
}
