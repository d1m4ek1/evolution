package general

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	newerror "iNote/www/backend/pkg/NewError"
)

const (
	pathToError string = "pkg/general -> Function "
)

const (
	errorProfileDefault string = pathToError + "profileDefault"
)

type SignUpData struct {
	Nickname string
	Email    string
	Login    string
	Password string
	UserID   string
	Token    string
}

type HeaderData struct {
	UserId    string
	UserUrlId string
	CustomId  sql.NullString
	Auth      bool
}

type ProfileData struct {
	Name          string   `json:"name"`
	Logo          string   `json:"logo"`
	Banner        string   `json:"banner"`
	Verif         bool     `json:"verif"`
	NetworkStatus string   `json:"netStatus"`
	Position      []string `json:"position"`
	Audience      int      `json:"audience"`
	Auth          bool     `json:"auth"`
}

func (P *ProfileData) ValidLogoBanner(logo string, banner string) {
	if banner != "not_banner.png" {
		P.Banner = "/user_images/profile/banner/saved/" + banner
	} else {
		P.Banner = "/user_images/profile/banner/notBanner/not_banner.png"
	}
	if logo != "not_logo.png" {
		P.Logo = "/user_images/profile/logo/saved/" + logo
	} else {
		P.Logo = "/user_images/profile/logo/notLogo/not_logo.png"
	}
}

type SettingsData struct {
	Title string
}

func (S *SettingsData) SetTitle(s string) {
	switch s {
	case "profile":
		S.Title = "Настройки профиля"
	case "shop":
		S.Title = "Настройки магазина"
	case "page-appearance":
		S.Title = "Настройки внешнего вида страницы"
	case "personal-data":
		S.Title = "Настройки персональных данных"
	}
}

func (S *SignUpData) ValidData(s *SignUpData) {
	if len(s.Nickname) <= 110 && s.Nickname != "" {
		S.Nickname = s.Nickname
	}
	if len(s.Email) <= 64 && s.Email != "" {
		S.Email = s.Email
	}
	if len(s.Login) <= 100 && s.Login != "" {
		S.Login = s.Login
	}
	if len(s.Password) == 32 && s.Password != "" {
		S.Password = s.Password
	}
	if len(s.UserID) <= 32 && s.UserID != "" {
		S.UserID = s.UserID
	}
	if len(s.Token) == 40 && s.Token != "" {
		S.Token = s.Token
	}
}

func (PD *ProfileData) ProfileUser(ctx *sqlx.DB, id int64, token string) error {
	if err := ctx.Get(&PD.Auth, `SELECT count(*) = 1 FROM users_data WHERE id=$1 AND token=$2`, id, token); err != nil {
		newerror.Wrap("ctx.Get", err)
		return err
	}
	return nil
}

func ValidateUser(ctx *sqlx.DB, token, userId string) bool {
	var valid int32

	if err := ctx.DB.QueryRow(`SELECT COUNT(*) FROM users_data WHERE id=$1 AND token=$2`, userId, token).Scan(&valid); err != nil {
		return false
	} else {
		if valid != 0 {
			return true
		}
	}

	return false
}
