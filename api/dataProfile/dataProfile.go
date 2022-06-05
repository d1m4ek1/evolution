package dataprofile

import (
	"encoding/json"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"
	"log"
	"net/http"

	"github.com/lib/pq"
)

// Path to error
const (
	pathToError string = "api/dataProfile -> Function "
)

const (
	errorSendAllData string = pathToError + "sendAllData"
)

type DataProfile struct {
	AboutMeTitle   string `json:"aboutmeTitle"`
	AboutMeContent string `json:"aboutmeContent"`
}

func sendAllData(w http.ResponseWriter, userId string) {
	var aboutme []string
	var dataProfile DataProfile

	if err := database.Tables.QueryRow(`SELECT aboutme
	FROM settings WHERE id=$1`, userId).Scan(pq.Array(&aboutme)); err != nil {
		log.Println(newerror.Wrap(errorSendAllData, "Query at db: 1", err))
	}

	dataProfile.AboutMeTitle = aboutme[0]
	dataProfile.AboutMeContent = aboutme[1]

	if err := json.NewEncoder(w).Encode(&dataProfile); err != nil {
		log.Println(newerror.Wrap(errorSendAllData, "json.NewEncoder", err))
	}
}

func ControlDataProfile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("get_data") == "all" {
		sendAllData(w, r.URL.Query().Get("user_id"))
		return
	}
}
