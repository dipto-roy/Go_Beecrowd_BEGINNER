package main

import "fmt"

const n = 3.14159

func main() {

	var radius float64
	fmt.Scan(&radius)
	area := n * radius * radius
	fmt.Printf("A=%.4f\n", area)

}
