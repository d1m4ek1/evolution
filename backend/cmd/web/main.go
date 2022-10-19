package main

import (
	"iNote/www/backend/internal/database"
	"iNote/www/backend/internal/handlers"
	"log"
)

func main() {
	log.Println("Function main -> Function openDB -> Open database connection")

	var ctx database.Context
	ctx.DB = ctx.InitDatabase()
	defer ctx.DB.Close()

	log.Println("Function main -> Database received, connection closed")
	log.Println("Function main -> Preparing to start the server")

	handlers.InitHandlers(ctx.DB)
}
