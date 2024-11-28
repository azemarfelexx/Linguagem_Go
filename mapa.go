package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])

	y := make(map[int]int)
	y[1] = 20
	fmt.Println(y[1])

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	fmt.Println(elemento["He"])

}
