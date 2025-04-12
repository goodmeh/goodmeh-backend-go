package main

import (
	"context"
	"goodmeh/app/router"
	"goodmeh/app/socket"
	"goodmeh/deps"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	// deps.InitLog()
}

func main() {
	port := os.Getenv("PORT")
	ctx := context.Background()

	// Initialize the database connection pool instead of a single connection
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	// Optional: Configure the pool size based on your needs
	poolConfig.MaxConns = 10 // Default is usually 4

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// Verify the connection works
	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	// Initialize websocket server
	socketServer := socket.NewServer()

	init := deps.Initialize(pool, ctx, &socketServer)
	app := router.Init(init, ctx)

	app.Run(":" + port)
}
