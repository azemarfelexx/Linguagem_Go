// Defer Escalona as funções

package main

import "fmt"

func dia1() {

	fmt.Println("Domingo")

}
func dia2() {
	fmt.Println("Segunda-Feira")
}

func main() {

	defer dia2()
	defer dia1()
}

// Panic: Por erro do programador / Erro de execuçao de tmpo, etc.
// Recover: Interrompe o panic

package main

import "fmt"

func main() {

	defer func() {

		x := recover()

		fmt.Println(x)

	}()

	panic("Pânico")
}
