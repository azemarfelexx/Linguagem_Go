package main

import "fmt"

func main() {

	for hour := 0; hour <= 12; hour++ {

		fmt.Println("Hora(s): ", hour)

		for min := 0; min < 60; min++ {

			fmt.Println(min)
		}
	}
}
