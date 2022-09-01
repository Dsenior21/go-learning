package main

import (
	"fmt"
)

func main() {

	menu()
	var selec string
	var num1 int
	var num2 int
	fmt.Scanln(&selec)

	switch selec {
	case "1":
		Suma(num1, num2, selec)

	case "2":
		resta(num1, num2, selec)

	case "3":
		multiplicar(num1, num2, selec)

	case "4":
		dividir(num1, num2, selec)

	default:
		fmt.Println("Selección no válida")
		fmt.Println(menu)

	}

}

func Suma(num1 int, num2 int, selec string) {
	fmt.Println("Digite el primer número")
	fmt.Scanln(&num1)
	fmt.Println("Digite el segundo número")
	fmt.Scanln(&num2)
	fmt.Println("el resultado de la suma es ", num1+num2)
}

func resta(num1 int, num2 int, selec string) {
	fmt.Println("Digite el primer número")
	fmt.Scanln(&num1)
	fmt.Println("Digite el segundo número")
	fmt.Scanln(&num2)
	fmt.Println("el resultado de la suma es ", num1-num2)
}

func multiplicar(num1 int, num2 int, selec string) {
	fmt.Println("Digite el primer número")
	fmt.Scanln(&num1)
	fmt.Println("Digite el segundo número")
	fmt.Scanln(&num2)
	fmt.Println("el resultado de la suma es ", num1*num2)
}

func dividir(num1 int, num2 int, selec string) {
	fmt.Println("Digite el primer número")
	fmt.Scanln(&num1)
	fmt.Println("Digite el segundo número")
	fmt.Scanln(&num2)
	fmt.Println("el resultado de la suma es ", num1/num2)
}

func menu() {
	menu :=
		`
Welcome, select one option
1) suma
2) resta
3) multiplicar
4) dividir
`
	fmt.Print(menu)
}
