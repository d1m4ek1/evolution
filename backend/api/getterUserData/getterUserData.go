package getteruserdata

import (
	"iNote/www/backend/models"
	newerror "iNote/www/backend/pkg/newerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const pathToLogFile string = "backend/logs/logs.txt"
const isTimeAmPm bool = true

func GetUserDataDefault(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(context *gin.Context) {
		convSenderID, err := strconv.ParseInt(context.Query("sender_id"), 10, 0)
		if err != nil {
			newerror.NewAppError("strconv.ParseInt", err, pathToLogFile, isTimeAmPm)
			return
		}

		name, logo, err := models.SelectUserDataDefault(ctx, convSenderID)
		if err != nil {
			newerror.NewAppError("models.SelectUserDataDefault", err, pathToLogFile, isTimeAmPm)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"logo": logo,
		})
	})
}
