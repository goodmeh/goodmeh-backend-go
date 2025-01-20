package main

import (
	"goodmeh/app/router"
	"goodmeh/config"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	// config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Initialize()
	app := router.Init(init)

	app.Run(":" + port)
}
