package main

import (
	"github.com/gauryan/ssmon/config"
	"github.com/gauryan/ssmon/database"
	"github.com/gauryan/ssmon/routes"
	"github.com/gauryan/ssmon/store"
	// "fmt"
)

func main() {
	app := routes.Router()
	app.Static("/", "./static")
	database.Init()
	store.Init()
	app.Listen(":" + config.Config("APP_PORT"))
}
