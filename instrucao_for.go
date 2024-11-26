package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		fmt.Println(i)

		if i%2 == 0 {

			fmt.Println("Número Par")

		} else {
			fmt.Println("Número Ímpar")
		}

	}
}

// infromar se o número é par ou ímpar
// resto é representado por %
