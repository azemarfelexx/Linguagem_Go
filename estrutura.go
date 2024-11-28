package main

import "fmt"

type biblioteca struct {
	titulo  string
	autor   string
	paginas int
}

func main() {
	fmt.Println(biblioteca{"Gritos do Passado, ", "Camilla Lackberg, ", 366})
	fmt.Println(biblioteca{"Cordeiro Assassino, ", "Ken Englade, ", 204})
}
