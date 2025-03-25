package main

import (
	"goodmeh/app/router"
	"goodmeh/deps"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	// deps.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := deps.Initialize()
	app := router.Init(init)

	app.Run(":" + port)
}
