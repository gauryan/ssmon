package main

import (
	"github.com/gauryan/ssmon/routes"
	"github.com/gauryan/ssmon/database"
	"github.com/gauryan/ssmon/store"
)


func main() {
	app := routes.Router()
	database.Init()
	store.Init()
	app.Listen(":3000")
}
