package main

import (
	"github.com/AleksandarAtanackovic/fishstore/backend/internal/orders"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	/*
		router.GET("/api/v1/orders", orders.GetOrders)
		router.GET("/api/v1/orders/:id", orders.GetOrderByApiId)
		router.GET("/api/v1/orders/unfinished", orders.GetUnfinishedOrders)
		router.PATCH("/api/v1/orders/:id", orders.TogglePrepared)
		router.POST("/api/v1/orders", orders.AddOrder)
	*/
	router.GET("/api/v2/customers", orders.GetCustomers)
	router.POST("/api/v2/customers", orders.AddOrder)
	router.Run("localhost:9090")

}
