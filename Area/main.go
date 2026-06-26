package main

import "fmt"

const pi =  3.14159

func main() {
	var A , B , C float64
	fmt.Scan(&A , &B , &C)

	fmt.Printf("TRIANGULO: %.3f\n", float64(0.5 * A * C))
	fmt.Printf("CIRCULO: %.3f\n",float64(pi * C * C))
	fmt.Printf("TRAPEZIO: %.3f\n",float64((A + B) / 2) * C)
	fmt.Printf("QUADRADO: %.3f\n",float64(B * B))
	fmt.Printf("RETANGULO: %.3f\n",float64(A * B))


}
