package main

import "fmt"

func main() {
	//fmt.Println("Hello, World!")
	//fmt.Println("Welcome to Go programming.")

	//var name string = "Dip"
	//fmt.Println("Hello, " + name + "!")

	var age int

	fmt.Scan(&age)
	fmt.Printf("You are %d years old.\n", age)
}