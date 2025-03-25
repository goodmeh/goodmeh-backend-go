package main

import (
	"context"
	"goodmeh/app/router"
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
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	init := deps.Initialize(conn, ctx)
	app := router.Init(init)

	app.Run(":" + port)
}
