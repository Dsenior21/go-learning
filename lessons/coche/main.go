package main

import "fmt"

type coche struct {
	numPuertas int
}

func suma(a int, b int) {
	valores := a + b
	fmt.Println(valores)

}

func (c *coche) puertas() int {
	i := 4
	c.numPuertas = i + 1
	return c.numPuertas
}

func main() {
	suma(10, 20)
	fmt.Println()

}
