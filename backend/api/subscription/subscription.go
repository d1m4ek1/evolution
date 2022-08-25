package subscription

import (
	"fmt"
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/newerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const pathToLogFile string = "backend/logs/logs.txt"
const isTimeAmPm bool = true

func getUserAuthData(context *gin.Context) (string, int64, error) {
	token, _ := context.Cookie("token")
	userId, _ := context.Cookie("userId")

	var userIDConv int64
	var err error

	if userId != "" {
		userIDConv, err = strconv.ParseInt(userId, 10, 0)
		if err != nil {
			newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
			return "", 0, err
		}
	}

	return token, userIDConv, nil
}

func AppendSubscription(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, userID, err := getUserAuthData(context)
		if err != nil {
			newerror.NewAppError("getUserAuthData", err, pathToLogFile, isTimeAmPm)
			return
		}

		fmt.Println(token, userID)

		if token != "" && userID != 0 {
			user := models.CheckSignin{
				Id:       userID,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.NewAppError("user.CheckUserOnSignin", err, pathToLogFile, isTimeAmPm)
				return
			}

			if user.Autorize {
				appendUserID, err := strconv.ParseInt(context.Query("append_id"), 10, 0)
				if err != nil {
					newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
					return
				}

				if err := models.SetAppendSubscriber(ctx, userID, appendUserID); err != nil {
					newerror.NewAppError("models.SetAppendSubscriber", err, pathToLogFile, isTimeAmPm)
					return
				}
			} else {
				context.Redirect(http.StatusMovedPermanently, "/signin")
			}
		} else {
			context.Redirect(http.StatusMovedPermanently, "/signin")
		}
	})
}

func DeleteSubscription(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, userID, err := getUserAuthData(context)
		if err != nil {
			newerror.NewAppError("getUserAuthData", err, pathToLogFile, isTimeAmPm)
		}

		if token != "" && userID != 0 {
			user := models.CheckSignin{
				Id:       userID,
				Token:    token,
				Autorize: false,
			}
			if err := user.CheckUserOnSignin(ctx); err != nil {
				newerror.NewAppError("user.CheckUserOnSignin", err, pathToLogFile, isTimeAmPm)
				return
			}

			if user.Autorize {
				deleteUserID, err := strconv.ParseInt(context.Query("delete_id"), 10, 0)
				if err != nil {
					newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
					return
				}

				if err := models.SetDeleteSubscriber(ctx, userID, deleteUserID); err != nil {
					newerror.NewAppError("models.SetDeleteSubscriber", err, pathToLogFile, isTimeAmPm)
					return
				}
			} else {
				context.Redirect(http.StatusMovedPermanently, "/signin")
			}
		} else {
			context.Redirect(http.StatusMovedPermanently, "/signin")
		}
	})
}

func CheckSubscriber(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		token, userID, err := getUserAuthData(context)
		if err != nil {
			newerror.NewAppError("getUserAuthData", err, pathToLogFile, isTimeAmPm)
		}

		if token != "" && userID != 0 {
			isSubscriber, err := models.SelectSubscriber(ctx, userID, context.Query("check_id"))
			if err != nil {
				newerror.NewAppError("models.SelectSubscriber", err, pathToLogFile, isTimeAmPm)
				context.JSON(http.StatusOK, gin.H{
					"isSubscriber": false,
				})
				return
			}

			context.JSON(http.StatusOK, gin.H{
				"isSubscriber": isSubscriber,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"isSubscriber": false,
			})
		}
	})
}

func CountSubscriber(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		userID, err := strconv.ParseInt(context.Query("check_id"), 10, 0)
		if err != nil {
			newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
			return
		}

		isCount, err := models.SelectCountSubscriber(ctx, userID)
		if err != nil {
			newerror.NewAppError("models.SelectSubscriber", err, pathToLogFile, isTimeAmPm)
			context.JSON(http.StatusOK, gin.H{
				"isCount": 0,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"isCountSubscriber": isCount.Int64,
		})
	})
}
