//A palavra reservada para goroutines é go
//Goroutines é, basicamente, o ato de fazer progresso com o seu código em mais de uma tarefa simultaneamente

package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var escreva string
	fmt.Scanln(&escreva)
}
