// 1 - Mostrar os números entre 1 e 100, divisíveis por 3:

package main

import (
	"fmt"
)

func main() {

	for x := 0; x < 100; x++ {

		if x%3 == 0 {

			fmt.Println(x)

		} else {
			continue
		}

	}
}

// 2 - Imprimir os números de "1" a "100". 
// Se o número divisível por "3", imprimir "Pin" em vez do número, se divisível por "5", imprimir "Pan", se devisível por "3" e "5", imprimir "Pin-Pan".

package main

import "fmt"

func main() {

	for a := 0; a < 100; a++ {

		if a%3 == 0 && a%5 == 0 {

			fmt.Println("Pin-Pan")

		} else if a%3 == 0 {
			fmt.Println("Pin")
		} else if a%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(a)
		}

	}

}
