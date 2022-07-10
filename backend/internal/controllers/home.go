package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"iNote/www/backend/pkg/general"
	"net/http"
)

func HomeTemplate(ctx *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		token, _ := c.Cookie("token")
		data := defineHeaderForAutorize(ctx, token)

		Data := struct {
			HeaderData general.HeaderData
		}{
			HeaderData: data,
		}

		c.HTML(http.StatusOK, "index", Data)
	})
}
