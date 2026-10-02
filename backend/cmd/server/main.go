package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/AleksandarAtanackovic/fishstore/backend/internal/orders"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:password@localhost:5432/fishstore?sslmode=disable"
	}

	ctx := context.Background()

	// pgxpool.New samo pripremi pool, još se ne povezuje na bazu
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("invalid database url: %v", err)
	}
	defer pool.Close()

	// Ping zaista otvori konekciju i proveri lozinku.
	// Ako baza ne odgovori za 5 sekundi, odustajemo.
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	log.Println("connected to database")

	router := gin.Default()
	router.GET("/api/v1/orders", orders.GetOrders)
	router.GET("/api/v1/orders/:id", orders.GetOrderByApiId)
	router.GET("/api/v1/orders/unfinished", orders.GetUnfinishedOrders)
	router.PATCH("/api/v1/orders/:id", orders.TogglePrepared)
	router.POST("/api/v1/orders", orders.AddOrder)
	router.Run("0.0.0.0:9090")
}
