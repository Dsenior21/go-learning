package main

import (
	"fmt"
)

func main() {
	menu :=
		`
Welcome, select one option
1) suma
2) resta
3) multiplicar
4) dividir
`
	fmt.Print(menu)

	var selec string
	var num1 int
	var num2 int
	fmt.Scanln(&selec)

	switch selec {
	case "1":
		fmt.Println("Digite el primer número")
		fmt.Scanln(&num1)
		fmt.Println("Digite el segundo número")
		fmt.Scanln(&num2)
		fmt.Println("el resultado de la suma es ", num1+num2)

	case "2":
		fmt.Println("Digite el primer número")
		fmt.Scanln(&num1)
		fmt.Println("Digite el segundo número")
		fmt.Scanln(&num2)
		fmt.Println("el resultado de la suma es ", num1-num2)

	case "3":
		fmt.Println("Digite el primer número")
		fmt.Scanln(&num1)
		fmt.Println("Digite el segundo número")
		fmt.Scanln(&num2)
		fmt.Println("el resultado de la suma es ", num1*num2)

	case "4":
		fmt.Println("Digite el primer número")
		fmt.Scanln(&num1)
		fmt.Println("Digite el segundo número")
		fmt.Scanln(&num2)
		fmt.Println("el resultado de la suma es ", num1/num2)

	default:
		fmt.Println("Selección no válida")
		fmt.Println(menu)

	}

}
