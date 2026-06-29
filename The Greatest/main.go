package main

import (
	"fmt"
	//"math"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	// max2 := func(x, y int) int {
	// 	return (x + y + int(math.Abs(float64(x-y)))) / 2
	// }

	// greatest := max2(max2(a, b), c)

	// fmt.Printf("%d eh o maior\n", greatest)


	if (a >= b && a >= c) {
		fmt.Printf("%d eh o maior\n", a)
	} else if (b >= a && b >= c) {
		fmt.Printf("%d eh o maior\n", b)
	} else {
		fmt.Printf("%d eh o maior\n", c)
	}
}
