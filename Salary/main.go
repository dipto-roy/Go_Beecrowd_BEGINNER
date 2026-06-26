package main
import "fmt"

func main () {
	var a , b int
	var c float64

	fmt.Scan(&a , &b , &c)
	fmt.Printf("NUMBER = %d\n",a)
	sum := float64(b) * c
	fmt.Printf("SALARY = U$ %.2f\n", sum)


}	
