package main

import "fmt"

const pi = 3.14159

func main() {

	var radius int 
	fmt.Scan(&radius)
	Volume := (4.0 / 3.0) * pi * float64(radius*radius*radius)
	fmt.Printf("VOLUME = %.3f\n", Volume)


} 