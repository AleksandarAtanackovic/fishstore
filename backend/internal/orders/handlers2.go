package orders

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type order struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FishType  Fish      `json:"fish_type"`
	OrderType OrderType `json:"order_type"`
	Ready     bool      `json:"prepared"`
	Completed bool      `json:"completed"`
	OrderTime time.Time `json:"order_time"`
}

type customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	Order       order  `json:"order"`
}

//Building a simple RESTful API

var customers = []customer{}

func GetCustomers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, customers)
}

/*
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

	if !newOrder.FishType.isValidFish() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid fish type"})
		return
	}

	if !newOrder.OrderType.isValidOrderType() {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid order type"})
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

func (f Fish) isValidFish() bool {
	switch f {
	case Saran, Oslic, Pastrmka:
		return true
	}
	return false
}

func (f OrderType) isValidOrderType() bool {
	switch f {
	case Fry, Clean, Fresh:
		return true
	}
	return false
}
*/
