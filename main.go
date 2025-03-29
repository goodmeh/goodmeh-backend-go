package main

import (
	"context"
	"goodmeh/app/router"
	"goodmeh/app/socket"
	"goodmeh/deps"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	// deps.InitLog()
}

func main() {
	port := os.Getenv("PORT")
	ctx := context.Background()

	// Initialize the database connection
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	// Initialize websocket server
	socketServer := socket.NewServer()

	init := deps.Initialize(conn, ctx, &socketServer)
	app := router.Init(init, ctx)

	app.Run(":" + port)
}
