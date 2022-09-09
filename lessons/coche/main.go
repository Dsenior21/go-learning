package main

import "fmt"

type coche struct {
	numPuertas int
}

func suma(a int, b int) {
	valores := a + b
	fmt.Println(valores)

}

func puertas(c *coche) int {
	(*c).numPuertas += 1
	return c.numPuertas
}

func main() {
	c := coche{}
	puertas(&c)
	fmt.Println(c)

	suma(10, 20)

}
