package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/AleksandarAtanackovic/fishstore/backend/internal/orders"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:password@localhost:5432/fishstore?sslmode=disable"
	}
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	//Testing database connection with dummy data

	//Executed the following queries manually:
	/*
		INSERT INTO customers (name, surname, phone_number, created_at)
		VALUES ('Aleksandar','Atanackovic','123456',NOW());

		INSERT INTO orders (customer_id, order_time, completed)
		VALUES (1, NOW(), false);

		INSERT INTO order_items (order_id, fish_type, order_type, prepared)
		VALUES
		(1,'saran','fry',false),
		(1,'oslic','clean',true);
	*/
	println("there wasnt an error with the database")
	//orders.InitDB(db)
	rows, err := db.Query("SELECT * FROM customers")

	if err != nil {
		println(err)
	}
	defer rows.Close()

	router := gin.Default()
	router.GET("/api/v1/orders", orders.GetOrders)
	router.GET("/api/v1/orders/:id", orders.GetOrderByApiId)
	router.GET("/api/v1/orders/unfinished", orders.GetUnfinishedOrders)
	router.PATCH("/api/v1/orders/:id", orders.TogglePrepared)
	router.POST("/api/v1/orders", orders.AddOrder)
	router.Run("0.0.0.0:9090")
}
