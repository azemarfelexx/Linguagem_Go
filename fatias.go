package main

import "fmt"

func main() {

	arr := [8]float64{1, 2, 3, 4, 5, 6, 7, 8}

	fatia1 := arr[0:7]
	fatia2 := append(fatia1, 9, 10)

	fmt.Println(fatia1, fatia2)

}
