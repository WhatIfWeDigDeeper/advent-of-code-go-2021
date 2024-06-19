package main

import (
	"fmt"
	"math"
	"math/rand"
)

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func exampleMath() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello, world.")
	exampleMath()
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
