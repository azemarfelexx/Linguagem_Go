package main

import "fmt"

func main() {

	x := 100
	y := 153
	z := 200

	if x%2 == 0 || y%2 == 0 {

		fmt.Println("Ao menos uma das variáveis (x ou y) possui valor ímpar")
	} else {

		fmt.Println("Ao menos uma das variáveis possui valor ímpar")
	}

	if x%2 == 0 && y%2 == 0 {

		fmt.Println("Os valores de x e y são pares")
	} else {

		fmt.Println("Ao menos uma das variáveis possui valor ímpar")
	}

	if x%2 == 0 && z%2 == 0 {

		fmt.Println("As variáveis x e z são valores pares")
	} else {

		fmt.Println("Ao menos uma das variáveis possui valor ímpar")
	}

}
