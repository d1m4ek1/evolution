package general

import (
	"database/sql"
	"net/http"

	"iNote/www/internal/database"
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

type UserData struct {
	UserId        int            `json:"user_id"`
	Name          string         `json:"name"`
	Birthday      sql.NullString `json:"birthday"`
	Position      []string       `json:"position"`
	MainAddress   []int          `json:"mainAddress"`
	NetworkStatus string         `json:"netStatus"`
	Logo          string         `json:"logo"`
	Banner        string         `json:"banner"`
	Audience      []string       `json:"audience"`
	Verif         sql.NullString `json:"verif"`
	FirstName     sql.NullString `json:"firstName"`
	LastName      sql.NullString `json:"lastName"`
	OldToken      string         `json:"olt"`
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
		P.Banner = "/profile/banner/saved/" + banner
	} else {
		P.Banner = "/profile/banner/notBanner/not_banner.png"
	}
	if logo != "not_logo.png" {
		P.Logo = "/profile/logo/saved/" + logo
	} else {
		P.Logo = "/profile/logo/notLogo/not_logo.png"
	}
}

type DataArray struct {
	Position []string `json:"position"`
	Audience []string `json:"audience"`
}

type SignInData struct {
	Login    string
	Password string
	NewToken string
	OldToken sql.NullString
}

type SettingsData struct {
	Title string
}

type SettingsProfileData struct {
	Name           string `json:"name"`
	Logo           string `json:"logo"`
	Banner         string `json:"banner"`
	ThemePage      string `json:"themePage"`
	Language       string `json:"language"`
	AboutMeTitle   string `json:"aboutMeTitle"`
	AboutMeContent string `json:"aboutMeContent"`
	Telegram       string `json:"telegram"`
	Instagram      string `json:"instagram"`
	Facebook       string `json:"facebook"`
	Vk             string `json:"vk"`
	Tiktok         string `json:"tiktok"`
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

func (PD *ProfileData) ProfileUser(id, token string) {
	var user string

	database.Tables.QueryRow(`SELECT login FROM users_data WHERE id=$1 AND token=$2`, id, token).Scan(&user)

	if user != "" {
		PD.Auth = true
	} else {
		PD.Auth = false
	}
}

func ValidateUser(token, userId *http.Cookie) bool {
	var valid int32

	if err := database.Tables.QueryRow(`SELECT COUNT(*) FROM users_data WHERE id=$1 AND token=$2`, userId.Value, token.Value).Scan(&valid); err != nil {
		return false
	} else {
		if valid != 0 {
			return true
		}
	}

	return false
}
