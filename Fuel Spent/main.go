package main
import (
	"fmt"
)

func main() {
  var time , avgSpeed int
  fmt.Scan(&time , &avgSpeed)
  fmt.Printf("%.3f\n", float64(time * avgSpeed) / 12.0)
  
//   avg := float64(avgSpeed) / 12.0

//   fmt.Printf("%.3f\n", avg *float64(time))


}