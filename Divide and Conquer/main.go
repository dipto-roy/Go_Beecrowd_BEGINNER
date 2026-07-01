package main
import (
	"fmt"
)

func main() {
	var tt int
	fmt.Scan(&tt)

	for i := 0; i < tt; i++ {
		var a , b int 
		fmt.Scan(&a , &b)

		isValid := a % b == 0

		if isValid {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}