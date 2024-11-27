package main

import "fmt"

func main() {

	var x [10]float64

	x[0] = 8.0
	x[1] = 8.2
	x[2] = 7.8
	x[3] = 9.3
	x[4] = 8.4
	x[5] = 9.8
	x[6] = 7.9
	x[7] = 8.1
	x[8] = 9.0
	x[9] = 9.2

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println("A mésdia das notas é: ", total/float64(len(x)))
}
