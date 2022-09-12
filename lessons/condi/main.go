package main

import "fmt"

func main() {

	numerolf := -3

	if numerolf < 0 {
		fmt.Println("número negativo")
	} else if numerolf > 0 {
		fmt.Println("número positivo")
	} else if numerolf == 0 {
		fmt.Println("número igual a 0")
	}

	for numeroWhile := 0; numeroWhile < 3; numeroWhile++ {
		fmt.Println(numeroWhile)
	}

	numerowhile2 := 2
	for numerowhile2 < 3 {
		fmt.Println(numerowhile2)
		numerowhile2++
	}

	for numeroFor := 0; numeroFor <= 3; numeroFor++ {
		fmt.Println(numeroFor)
	}
	fmt.Println("escriba la estación en que se encuentra")
	var estacion string
	fmt.Scanln(&estacion)

	switch estacion {
	case "invierno":
		fmt.Println("Hace mucho frío")
	case "verano":
		fmt.Println("hace mucho calor")
	case "primavera":
		fmt.Println("hace bonito día")
	case "otoño":
		fmt.Println("Es hora de dormir")
	default:
		fmt.Println("Eso no es una estación")
	}
}
