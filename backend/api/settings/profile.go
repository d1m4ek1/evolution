package settings

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/models"
	"iNote/www/backend/pkg/NewError"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var keyQueryFiles = [2]string{"logo", "banner"}
var keyQueryUser = [1]string{"name"}
var keyQuerySettings = [5]string{"language", "theme_page", "aboutme_title", "aboutme_content"}
var keyQueryConnection = [5]string{"telegram", "instagram", "facebook", "vk", "tiktok"}

func GetProfileSettings(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")
		userIDConv, err := strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.Wrap("strconv.ParseInt", err)
			return
		}

		if token != "" && userId != "" {
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.Wrap("user.CheckUserOnSignin", err)
				return
			}

			if user.Autorize {
				userData, err := models.SelectProfileSettings(ctx, userIDConv)
				if err != nil {
					newerror.Wrap("models.SelectProfileSettings", err)
					return
				}
				context.JSON(http.StatusOK, userData)

			} else {
				context.Redirect(http.StatusMovedPermanently, "/signin")
			}
		}
	})
}

func exists(s string, key string) bool {
	if _, err := os.Stat("./user_images/profile/" + key + "/saved/" + s); !os.IsNotExist(err) {
		return false
	}
	return true
}

func storee(bv []byte) string {
	hasher := sha1.New()
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}

func validFile(s string, id int64) (string, error) {
	var filePath string

	if strings.Contains(s, ".png") {
		sha := storee([]byte(strings.Replace(s, ".png", "", -1)))
		filePath = "id" + fmt.Sprint(id) + "_" + sha + ".png"

	} else if strings.Contains(s, ".jpeg") {
		sha := storee([]byte(strings.Replace(s, ".jpeg", "", -1)))
		filePath = "id" + fmt.Sprint(id) + "_" + sha + ".jpeg"

	} else if strings.Contains(s, ".jpg") {
		sha := storee([]byte(strings.Replace(s, ".jpg", "", -1)))
		filePath = "id" + fmt.Sprint(id) + "_" + sha + ".jpg"
	} else if strings.Contains(s, ".gif") {
		sha := storee([]byte(strings.Replace(s, ".gif", "", -1)))
		filePath = "id" + fmt.Sprint(id) + "_" + sha + ".gif"
	} else {
		return "", errors.New("файл не является изображением")
	}

	return filePath, nil
}

func saveImage(ctx *sqlx.DB, context *gin.Context, keyFile string, userID int64) {
	oldFilePath, err := models.SelectFilePath(ctx, keyFile, userID)
	if err != nil {
		newerror.Wrap("models.SelectFilePath", err)
		return
	}

	src, hdr, err := context.Request.FormFile(keyFile)
	if err != nil {
		newerror.Wrap("context.Request.FormFile", err)
		return
	}
	defer src.Close()

	if exists(oldFilePath, keyFile) {
		os.Remove("user_images/profile/" + keyFile + "/saved/" + oldFilePath)
	}

	filePath, err := validFile(hdr.Filename, userID)
	if err != nil {
		newerror.Wrap("validFile", err)
		return
	}

	if filePath != "" {
		filePath = strings.Replace(filePath, "=", "", -1)

		dst, err := os.Create("user_images/profile/" + keyFile + "/saved/" + filePath)
		if err != nil {
			newerror.Wrap("os.Create", err)
			return
		}
		defer dst.Close()

		if err := models.SetFilePath(ctx, keyFile, filePath, userID); err != nil {
			newerror.Wrap("models.SetFilePath", err)
			return
		}

		if _, err := io.Copy(dst, src); err != nil {
			newerror.Wrap("io.Copy", err)
			return
		}
	}
}

func saveNewSettings(ctx *sqlx.DB, context *gin.Context, userID int64) {
	for _, v := range keyQueryFiles {
		if context.Query(v) != "" {
			saveImage(ctx, context, v, userID)
		}
	}

	var tplSettings, tplConnection, tplUser []string

	for _, v := range keyQuerySettings {
		if v == "aboutme_title" || v == "aboutme_content" {
			if context.Query(v) != "" {
				var aboutmeValueArrs string
				if v == "aboutme_title" {
					aboutmeValueArrs = "aboutme[1]"
				} else {
					aboutmeValueArrs = "aboutme[2]"
				}
				tplSettings = append(tplSettings, fmt.Sprintf("%s='%s'", aboutmeValueArrs, context.Query(v)))
			}
		} else {
			if context.Query(v) != "" {
				tplSettings = append(tplSettings, fmt.Sprintf("%s='%s'", v, context.Query(v)))
			}
		}
	}
	for _, v := range keyQueryConnection {
		if context.Query(v) != "" {
			tplConnection = append(tplConnection, fmt.Sprintf("%s='%s'", v, context.Query(v)))
		}
	}
	for _, v := range keyQueryUser {
		if context.Query(v) != "" {
			tplUser = append(tplUser, fmt.Sprintf("%s='%s'", v, context.Query(v)))
		}
	}

	if err := models.SetSettingsProfile(ctx,
		strings.Join(tplSettings, ","),
		strings.Join(tplConnection, ","),
		strings.Join(tplUser, ","), userID); err != nil {
		newerror.Wrap("models.SetSettingsProfile", err)
		return
	}
}

func SaveProfileSettings(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, _ := context.Cookie("token")
		userId, _ := context.Cookie("userId")
		userIDConv, err := strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.Wrap("strconv.ParseInt", err)
			return
		}

		if token != "" && userId != "" {
			user := models.CheckSignin{
				Id:       userIDConv,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.Wrap("user.CheckUserOnSignin", err)
				return
			}

			if user.Autorize {
				saveNewSettings(ctx, context, userIDConv)
			} else {
				context.Redirect(http.StatusMovedPermanently, "/signin")
				return
			}
		}
	})
}
