package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"iNote/www/api/autorization/signin"
	"iNote/www/api/autorization/signout"
	"iNote/www/api/autorization/signup"
	checkonline "iNote/www/api/checkOnline"
	"iNote/www/api/settings"
	"iNote/www/internal/controllers"
	"iNote/www/internal/database"
	newerror "iNote/www/pkg/NewError"

	"github.com/gorilla/mux"
)

// PATHS
const (
	pathToError string = "cmd/web/main -> Function "
)

// Errors
var (
	errorConnectToDatabase string = pathToError + "ConnectToDatabase"
)

func handle() {
	rtr := mux.NewRouter()

	// INDEX TEMPLATE
	rtr.HandleFunc("/", controllers.HomeTemplate)

	// AUTORIZATION ROUTERS
	rtr.HandleFunc("/signin", controllers.AutorizTemplate)
	rtr.HandleFunc("/signup", controllers.AutorizTemplate)

	// MESSAGE ROUTERS
	rtr.HandleFunc("/messages", controllers.MessagesTemplate)
	rtr.HandleFunc("/messages/favorites", controllers.MessagesTemplate)
	rtr.HandleFunc("/messages/control", controllers.MessagesTemplate)

	// SETTINGS ROUTERS
	rtr.HandleFunc("/customize", controllers.SettingsTemplate)
	rtr.HandleFunc("/customize/{settingsType}", controllers.SettingsTemplate)

	// DIRECTORY ROUTERS
	rtr.HandleFunc("/directory", controllers.DirectoryTemplate)
	rtr.HandleFunc("/directory/{dircontent}", controllers.DirectoryTemplate)

	// MUSIC ROUTERS
	rtr.HandleFunc("/music", controllers.MusicTemplate)

	// SUBSCRIPTIONS ROUTERS
	rtr.HandleFunc("/subscriptions", controllers.SubscriptionsTemplate)

	// SUBSCRIPTIONS ROUTERS
	rtr.HandleFunc("/orders", controllers.OrdersTemplate)

	// SUBSCRIPTIONS ROUTERS
	rtr.HandleFunc("/shop", controllers.ShopTemplate)

	// PROFILE ROUTERS
	rtr.HandleFunc("/{userName}", controllers.ProfileTemplate)
	rtr.HandleFunc("/{userName}/aboutMe", controllers.ProfileTemplate)

	// API ACCOUNT
	// SIGN UP
	rtr.HandleFunc("/api/create_account", signup.SignUp)
	// SIGN IN
	rtr.HandleFunc("/api/signin_account", signin.SignIn)
	// SIGN OUT
	rtr.HandleFunc("/api/signout_account", signout.SignOut)

	// API SETTINGS
	// PROFILE
	rtr.HandleFunc("/api/get_settings/profile", settings.GetProfileSettings)
	rtr.HandleFunc("/api/save_settings/profile", settings.SaveProfileSettings)
	// PERSONAL DATA
	rtr.HandleFunc("/api/get_settings/personal_data", settings.GetPersonalData)
	rtr.HandleFunc("/api/save_settings/personal_data", settings.SavePersonalData)

	// API CHECK STATUS
	rtr.HandleFunc("/api/check_status", checkonline.CheckOnline)

	http.Handle("/", rtr)

	log.Println("Function handler -> Routings initialized")

	http.Handle("/ui/images/", http.StripPrefix("/ui/images/", http.FileServer(http.Dir("./ui/images/"))))
	http.Handle("/ui/assets/", http.StripPrefix("/ui/assets/", http.FileServer(http.Dir("./ui/assets/"))))
	http.Handle("/profile/banner/", http.StripPrefix("/profile/banner/", http.FileServer(http.Dir("./profile/banner/"))))
	http.Handle("/profile/logo/", http.StripPrefix("/profile/logo/", http.FileServer(http.Dir("./profile/logo/"))))

	log.Println("Function handler -> Static files initialized")
	log.Println("Function handler -> Server started successfully")

	// localhost:8000 or 127.0.0.1:8000
	if err := http.ListenAndServe(":8000", nil); err == nil {
		log.Println("Function handler -> Server started successfully")
	} else {
		log.Println("Function handler -> Failed to start the server")
		log.Fatal(err)
	}
}

// Function creates connection for db
func openDB() *sql.DB {
	db, err := sql.Open("postgres", database.ConnectString)
	if err != nil {
		fmt.Println(newerror.Wrap(errorConnectToDatabase, "Connect to db", err))
	}

	return db
}

func main() {
	log.Println("Function main -> Function openDB -> Open database connection")

	database.Tables = openDB()
	defer database.Tables.Close()

	log.Println("Function main -> Database received, connection closed")
	log.Println("Function main -> Preparing to start the server")

	handle()
}
