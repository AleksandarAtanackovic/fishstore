package main

import (
	"database/sql"
	"log"

	"github.com/AleksandarAtanackovic/fishstore/backend/internal/orders"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "postgres://postgres:postgres@localhost:5432/fishstore?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	orders.InitDB(db)

	router := gin.Default()
	router.GET("/api/v1/orders", orders.GetOrders)
	router.GET("/api/v1/orders/:id", orders.GetOrderByApiId)
	router.GET("/api/v1/orders/unfinished", orders.GetUnfinishedOrders)
	router.PATCH("/api/v1/orders/:id", orders.TogglePrepared)
	router.POST("/api/v1/orders", orders.AddOrder)
	router.Run("0.0.0.0:9090")
}
