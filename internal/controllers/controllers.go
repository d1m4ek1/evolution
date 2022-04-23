package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	userdata "iNote/www/api/userData"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"iNote/www/pkg/general"

	"github.com/gorilla/mux"
)

// Path to error
const (
	pathToError string = "internal/controllers -> Function "
)

// ERRORS
const (
	errorHomeTemplate          string = pathToError + "HomeTemplate"
	errorMessagesTemplate      string = pathToError + "MessagesTemplate"
	errorProfileTemplate       string = pathToError + "ProfileTemplate"
	errorAutorizTemplate       string = pathToError + "AutorizTemplate"
	errorGetUserData           string = pathToError + "getUserData"
	errorSettingsTemplate      string = pathToError + "SettingsTemplate"
	errorMusicTemplate         string = pathToError + "MusicTemplate"
	errorSubscriptionsTemplate string = pathToError + "SubscriptionsTemplate"
	errorOrdersTemplate        string = pathToError + "OrdersTemplate"
	errorShopTemplate          string = pathToError + "ShopTemplate"
)

// Path to main templates
const (
	homeTpl          string = "./ui/html/index.html"
	autorizTpl       string = "./ui/html/Autoriz.html"
	profileTpl       string = "./ui/html/Profile.html"
	messagesTpl      string = "./ui/html/Messages.html"
	directoryTpl     string = "./ui/html/Directory.html"
	settingsTpl      string = "./ui/html/Settings.html"
	musicTpl         string = "./ui/html/Music.html"
	subscriptionsTpl string = "./ui/html/Subscriptions.html"
	ordersTpl        string = "./ui/html/Orders.html"
	shopTpl          string = "./ui/html/Shop.html"
)

type ReplyDaseOnToken struct {
	Path, Define, ErrorTpl, UserUrlId string
	Token, UserId                     *http.Cookie
	W                                 http.ResponseWriter
	Profile                           general.ProfileData
	Settings                          general.SettingsData
}

type Redirector struct {
	Token, UserId *http.Cookie
	W             http.ResponseWriter
	R             *http.Request
}

func replyBasedOnToken(r *ReplyDaseOnToken) {
	tpl, err := template.ParseFiles(r.Path)
	if err != nil {
		fmt.Println(newerror.Wrap(r.ErrorTpl, "load tpl", err))
		return
	}

	data := getUserData(r.Token)

	if r.Token != nil {

		if data.UserId == "" && !data.Auth {
			tpl.ExecuteTemplate(r.W, r.Define, nil)

		} else if data.UserId == r.UserId.Value && data.Auth {

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
			tpl.ExecuteTemplate(r.W, r.Define, Data)
		}
	} else {
		Data := struct {
			HeaderData  general.HeaderData
			ProfileData general.ProfileData
		}{
			HeaderData:  data,
			ProfileData: r.Profile,
		}
		tpl.ExecuteTemplate(r.W, r.Define, Data)
	}
}

func getUserData(token *http.Cookie) general.HeaderData {
	var headerData general.HeaderData
	if token != nil {
		if err := database.Tables.QueryRow(`SELECT ud.id, ud.user_custom_id FROM users_data ud WHERE ud.token=$1`, token.Value).Scan(&headerData.UserId, &headerData.CustomId); err != nil {
			fmt.Println(newerror.Wrap(errorGetUserData, "Query at db: 1", err))
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

func redirectUser(r *Redirector) {
	if r.Token == nil && r.UserId == nil {
		http.Redirect(r.W, r.R, "/signin", http.StatusSeeOther)
		return
	} else {
		if !general.ValidateUser(r.Token, r.UserId) {
			http.SetCookie(r.W, &http.Cookie{
				Name:   "token",
				Value:  "",
				Path:   "/",
				MaxAge: -1,

				HttpOnly: true,
			})
			http.SetCookie(r.W, &http.Cookie{
				Name:   "userId",
				Value:  "",
				Path:   "/",
				MaxAge: -1,

				HttpOnly: true,
			})
			http.Redirect(r.W, r.R, "/signin", http.StatusSeeOther)
		}
	}
}

func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(homeTpl)
	if err != nil {
		fmt.Println(newerror.Wrap(errorHomeTemplate, "load tpl", err))
		return
	}

	token, _ := r.Cookie("token")

	if token != nil {
		data := getUserData(token)

		if data.UserId == "" && !data.Auth {
			tpl.ExecuteTemplate(w, "index", nil)

		} else if data.UserId != "" && data.Auth {
			data.Auth = true
			Data := struct {
				HeaderData general.HeaderData
			}{
				HeaderData: data,
			}
			tpl.ExecuteTemplate(w, "index", Data)
		}
	} else {
		tpl.ExecuteTemplate(w, "index", nil)
	}
}

func ProfileTemplate(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	vars := mux.Vars(r)
	userUrlId := vars["userName"]

	profileDefaultData := userdata.GetUserDataStatic(token, userId, userUrlId)

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:      profileTpl,
		Define:    "profile",
		ErrorTpl:  errorProfileTemplate,
		UserUrlId: userUrlId,
		Token:     token,
		UserId:    userId,
		W:         w,
		Profile:   profileDefaultData,
	})
}

func AutorizTemplate(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:     autorizTpl,
		Define:   "autoriz",
		ErrorTpl: errorAutorizTemplate,
		Token:    token,
		UserId:   userId,
		W:        w,
	})
}

func MessagesTemplate(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	redirectUser(&Redirector{
		Token:  token,
		UserId: userId,
		W:      w,
		R:      r,
	})

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:     messagesTpl,
		Define:   "messages",
		ErrorTpl: errorMessagesTemplate,
		Token:    token,
		UserId:   userId,
		W:        w,
	})
}

func DirectoryTemplate(w http.ResponseWriter, r *http.Request) {

	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:     directoryTpl,
		Define:   "directory",
		ErrorTpl: errorMessagesTemplate,
		Token:    token,
		UserId:   userId,
		W:        w,
	})
}

func SettingsTemplate(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	redirectUser(&Redirector{
		Token:  token,
		UserId: userId,
		W:      w,
		R:      r,
	})

	vars := mux.Vars(r)
	title := vars["settingsType"]

	if title == "" {
		title = "profile"
	}

	var settingsData general.SettingsData
	settingsData.SetTitle(title)

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:     settingsTpl,
		Define:   "settings",
		ErrorTpl: errorSettingsTemplate,
		Token:    token,
		UserId:   userId,
		W:        w,
		Settings: settingsData,
	})
}

func MusicTemplate(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	redirectUser(&Redirector{
		Token:  token,
		UserId: userId,
		W:      w,
		R:      r,
	})

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:     musicTpl,
		Define:   "music",
		ErrorTpl: errorMusicTemplate,
		Token:    token,
		UserId:   userId,
		W:        w,
	})
}

func SubscriptionsTemplate(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	redirectUser(&Redirector{
		Token:  token,
		UserId: userId,
		W:      w,
		R:      r,
	})

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:     subscriptionsTpl,
		Define:   "subscriptions",
		ErrorTpl: errorSubscriptionsTemplate,
		Token:    token,
		UserId:   userId,
		W:        w,
	})
}

func OrdersTemplate(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	redirectUser(&Redirector{
		Token:  token,
		UserId: userId,
		W:      w,
		R:      r,
	})

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:     ordersTpl,
		Define:   "orders",
		ErrorTpl: errorOrdersTemplate,
		Token:    token,
		UserId:   userId,
		W:        w,
	})
}

func ShopTemplate(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Cookie("token")
	userId, _ := r.Cookie("userId")

	redirectUser(&Redirector{
		Token:  token,
		UserId: userId,
		W:      w,
		R:      r,
	})

	replyBasedOnToken(&ReplyDaseOnToken{
		Path:     shopTpl,
		Define:   "shop",
		ErrorTpl: errorShopTemplate,
		Token:    token,
		UserId:   userId,
		W:        w,
	})
}
