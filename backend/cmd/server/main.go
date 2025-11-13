package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var c, python, java bool
var fl float64

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("I like pie", math.Pi)
	fmt.Println("Result of addition is", add(4, 5))
	word1, word2 := swap("hello", "world")
	fmt.Println("Swaping words hello and world:", word1, word2)

	var i int
	fmt.Println(i, c, python, java, fl)
	fmt.Printf("%.2f\n", fl)
}
