package main

import "fmt"

type coche struct {
	numPuertas int
}

func suma(a int, b int) {
	valores := a + b
	fmt.Println(valores)

}

func new() *coche {
	numPuertas = 4
}

func (c *coche) puertas() int {
	c.numPuertas += 1
	return c.numPuertas
}

// make function like puertas above to print the doors

func main() {
	coche := new() // creates a new car with 4 doors
	coche.puertas() // adds one door
	suma(10, 20)
	fmt.Println()

}
