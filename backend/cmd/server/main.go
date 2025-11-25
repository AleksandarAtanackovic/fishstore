package main

import (
	"github.com/AleksandarAtanackovic/fishstore/backend/internal/orders"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/v1/orders", orders.GetOrders)
	router.GET("/api/v1/orders/:id", orders.GetOrderByApiId)
	router.PATCH("/api/v1/orders/:id", orders.TogglePrepared)
	router.POST("/api/v1/orders", orders.AddOrder)
	router.Run("localhost:9090")
}
