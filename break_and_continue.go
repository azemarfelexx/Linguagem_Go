package main

import "fmt"

func main() {

	x := 0

	for {

		if x < 20 {

			x++

			fmt.Println("x < 20")
		} else {

			break
		}

	}

	for x := 0; x < 20; x++ {

		if x == 3 { // ele vai pular o 3

			continue

		}

		fmt.Println(x)

	}

}
