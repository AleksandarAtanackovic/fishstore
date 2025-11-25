package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//The beginning idea for data model

type Fish string

const (
	Saran    Fish = "saran"
	Oslic    Fish = "oslic"
	Pastrmka Fish = "pastrmka"
)

type customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
}

type order struct {
	ID        string    `json:"id"`
	Customer  customer  `json:"customer"`
	CreatedAt time.Time `json:"created_at"`
	FishType  Fish      `json:"fish_type"`
	OrderType string    `json:"order_type"`
	Prepared  bool      `json:"prepared"`
	Completed bool      `json:"completed"`
}

//Building a simple RESTful API

var orders = []order{}

func getOrders(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, orders)
	fmt.Println("Printing the request context:")
	fmt.Println(context.Params)
}

func addOrder(context *gin.Context) {
	var newOrder order

	if err := context.BindJSON(&newOrder); err != nil {
		return
	}

	orders = append(orders, newOrder)

	if newOrder.CreatedAt.IsZero() {
		newOrder.CreatedAt = time.Now()
	}

	context.IndentedJSON(http.StatusCreated, newOrder)
}

func getOrder(context *gin.Context) {
	id := context.Param("id")
	order, err := getOrderById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Order not found"})
	}

	context.IndentedJSON(http.StatusOK, order)
}

func togglePrepared(context *gin.Context) {
	id := context.Param("id")
	order, err := getOrderById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Order not found"})
	}
	order.Prepared = !order.Prepared

	context.IndentedJSON(http.StatusOK, order)
}

func getOrderById(id string) (*order, error) {
	for i, t := range orders {
		if t.ID == id {
			return &orders[i], nil
		}
	}

	return nil, errors.New("order not found")
}

func main() {
	router := gin.Default()
	router.GET("/api/v1/orders", getOrders)
	router.GET("/api/v1/orders/:id", getOrder)
	router.PATCH("/api/v1/orders/:id", togglePrepared)
	router.POST("/api/v1/orders", addOrder)
	router.Run("localhost:9090")
}
