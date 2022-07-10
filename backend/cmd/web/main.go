package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"html/template"
	"iNote/www/backend/api/authorization"
	dataprofile "iNote/www/backend/api/dataProfile"
	"iNote/www/backend/api/messages"
	"iNote/www/backend/api/settings"
	"iNote/www/backend/api/subscription"
	"iNote/www/backend/internal/controllers"
	"iNote/www/backend/internal/database"
	newerror "iNote/www/backend/pkg/NewError"
	"iNote/www/backend/websocket"
	"log"
	"net/http"
	"path/filepath"
)

func initTemplates() *template.Template {
	var files []string

	var paths = []string{
		"./ui/html/*html",
		"./ui/html/template/Headers/*.html"}

	for _, path := range paths {
		file, err := filepath.Glob(path)
		if err != nil {
			newerror.Wrap("filepath.Glob", err)
			return nil
		}

		files = append(files, file...)
	}

	tmpls, err := template.ParseFiles(files...)
	if err != nil {
		newerror.Wrap("template.ParseFiles", err)
		return nil
	}

	return tmpls
}

func handle(ctx *sqlx.DB) {
	gin.SetMode(gin.ReleaseMode)
	ginRouter := gin.Default()
	ginRouter.SetHTMLTemplate(initTemplates())

	// INDEX TEMPLATE
	ginRouter.GET("/", controllers.HomeTemplate(ctx))

	// AUTORIZATION ROUTERS
	ginRouter.GET("/signin", controllers.AutorizTemplate(ctx))
	ginRouter.GET("/signup", controllers.AutorizTemplate(ctx))

	// MESSAGE ROUTERS
	ginRouter.GET("/inSocial", controllers.InSocialTemplate(ctx))
	ginRouter.GET("/inSocial/favorites", controllers.InSocialTemplate(ctx))
	ginRouter.GET("/inSocial/chat_:chatId", controllers.InSocialTemplate(ctx))

	// MUSIC ROUTERS
	ginRouter.GET("/inMusic", controllers.InMusicTemplate(ctx))

	//SUBSCRIPTIONS ROUTERS
	ginRouter.GET("/inBeats", controllers.InBeatsTemplate(ctx))

	// SETTINGS ROUTERS
	ginRouter.GET("/customize", controllers.SettingsTemplate(ctx))
	ginRouter.GET("/customize/:settingsType", controllers.SettingsTemplate(ctx))

	// DIRECTORY ROUTERS
	ginRouter.GET("/directory", controllers.DirectoryTemplate(ctx))
	ginRouter.GET("/directory/:dircontent", controllers.DirectoryTemplate(ctx))

	// SUBSCRIPTIONS ROUTERS
	ginRouter.GET("/subscriptions", controllers.SubscriptionsTemplate(ctx))

	// SUBSCRIPTIONS ROUTERS
	ginRouter.GET("/orders", controllers.OrdersTemplate(ctx))

	//PROFILE ROUTERS
	ginRouter.GET("/:userName", controllers.ProfileTemplate(ctx))
	ginRouter.GET("/:userName/aboutMe", controllers.ProfileTemplate(ctx))

	// API ACCOUNT
	// SIGN UP
	ginRouter.POST("/api/create_account", authorization.SignUp(ctx))
	// SIGN IN
	ginRouter.GET("/api/signin_account", authorization.SignIn(ctx))
	// SIGN OUT
	ginRouter.POST("/api/signout_account", authorization.SignOut(ctx))

	// API SETTINGS
	// PROFILE
	ginRouter.GET("/api/get_settings/profile", settings.GetProfileSettings(ctx))
	ginRouter.POST("/api/save_settings/profile", settings.SaveProfileSettings(ctx))
	// PERSONAL DATA
	ginRouter.GET("/api/get_settings/personal_data", settings.GetPersonalData(ctx))
	ginRouter.POST("/api/save_settings/personal_data", settings.SavePersonalData(ctx))

	// API PROFILE
	ginRouter.GET("/api/get_data_profile", dataprofile.ControlDataProfile(ctx))

	// API CONFIRM
	ginRouter.GET("/api/confirm", authorization.ConfirmPassword(ctx))

	// API CHECK AUTORIZATION
	ginRouter.GET("/api/check_authorization", authorization.CheckAuthoriztion(ctx))

	// API SUBSRIBERS
	ginRouter.POST("/api/append_subscriber", subscription.AppendSubscription(ctx))
	ginRouter.DELETE("/api/delete_subscriber", subscription.DeleteSubscription(ctx))
	ginRouter.GET("/api/check_subscriber", subscription.CheckSubscriber(ctx))
	ginRouter.GET("/api/count_subscriber", subscription.CountSubscriber(ctx))

	// API MESSAGES
	ginRouter.GET("/api/get_user_card_messages", messages.GetUserCardMessages(ctx))
	ginRouter.GET("/api/check_chat", messages.CheckChat(ctx))

	// WEBSOCKET
	ginRouter.GET("/websocket/connect", websocket.WebSocketConnect(ctx))

	log.Println("Function handler -> Routings initialized")

	ginRouter.StaticFS("/ui/images/", http.Dir("./ui/images/"))
	ginRouter.StaticFS("/ui/assets/", http.Dir("./ui/assets/"))
	ginRouter.StaticFS("/user_images/profile/banner/", http.Dir("./user_images/profile/banner/"))
	ginRouter.StaticFS("/user_images/profile/logo/", http.Dir("./user_images/profile/logo/"))

	log.Println("Function handler -> Static files initialized")
	log.Println("Function handler -> Server started successfully")

	// localhost:8000 or 127.0.0.1:8000
	if err := ginRouter.Run(":8000"); err != nil {
		newerror.Wrap("ginRouter.Run(\":8000\")", err)
	}
}

func main() {
	log.Println("Function main -> Function openDB -> Open database connection")

	var ctx database.Context
	ctx.DB = ctx.Init()
	defer ctx.DB.Close()

	log.Println("Function main -> Database received, connection closed")
	log.Println("Function main -> Preparing to start the server")

	handle(ctx.DB)
}
