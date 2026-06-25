package main
import "fmt"

func main () {
	var name string
	var salary , bonus float64
	fmt.Scan(&name , &salary , &bonus)
    sum := salary + bonus * 0.15
	fmt.Printf("TOTAL = R$ %.2f\n", sum)


}