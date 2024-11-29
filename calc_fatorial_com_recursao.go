// É a capacidade de uma função de chamar a si mesma

// Calcular fatorial

package main

import "fmt"

func fatorial(x int) int { // função fatorial vai recerber um valor (x) inteiro e tbm drvolver um inteiro

	if x == 0 {
		return 1
	}

	return x * fatorial(x-1)
}

func main() {

	fmt.Println(fatorial(9))
}
