package orders

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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

func GetOrders(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, orders)
}

func GetUnfinishedOrders(context *gin.Context) {
	var unfinished []order
	for _, order := range orders {
		if !order.Completed {
			unfinished = append(unfinished, order)
		}
	}
	context.IndentedJSON(http.StatusOK, unfinished)
}

func AddOrder(context *gin.Context) {
	var newOrder order

	if err := context.BindJSON(&newOrder); err != nil {
		return
	}

	if newOrder.CreatedAt.IsZero() {
		newOrder.CreatedAt = time.Now()
	}

	if !newOrder.FishType.IsValid() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid fish type"})
		return
	}

	orders = append(orders, newOrder)

	context.IndentedJSON(http.StatusCreated, newOrder)
}

func GetOrderByApiId(context *gin.Context) {
	id := context.Param("id")
	order, err := getOrderById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Order not found"})
	}

	context.IndentedJSON(http.StatusOK, order)
}

func TogglePrepared(context *gin.Context) {
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

func (f Fish) IsValid() bool {
	switch f {
	case Saran, Oslic, Pastrmka:
		return true
	}
	return false
}
