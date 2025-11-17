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
	//Practicing part of the code
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

	var k, j int = 1, 2
	var c, python, java = true, false, "no!"

	fmt.Println(k, j, c, python, java)

	//App part
	var Name, Surname string
	fmt.Println("Enter name:")
	fmt.Scanf("%v", &Name)
	fmt.Println("Enter surname:")
	fmt.Scanf("%v", &Surname)

	fmt.Printf("You have entered the name %v and surname %v! Welcome to our fishstore!\n", Name, Surname)
}
