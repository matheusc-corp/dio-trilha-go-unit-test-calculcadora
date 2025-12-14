package main

import (
	"fmt"
)

func somar(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}

	return total
}

func subtrair(i ...int) int {
	total := i[0]

	for j := 1; j < len(i); j++ {
		total -= i[j]
	}

	return total
}

func multiplicar(i ...int) int {
	total := 1

	for _, v := range i {
		total *= v
	}

	return total
}

func dividir(i ...int) int {
	total := i[0]

	for j := 1; j < len(i); j++ {

		if i[j] == 0 {
			fmt.Print("Impossivel dividir por zero!")

			return 0
		}

		total /= i[j]
	}

	return total
}

func main() {
	fmt.Println("Soma: ", somar(2, 3, 5))
	fmt.Println("Resto: ", subtrair(8, 2))
	fmt.Println("Produto: ", multiplicar(9, 4, -1))
	fmt.Print("Quociente: ", dividir(20, 2, 2))
}
