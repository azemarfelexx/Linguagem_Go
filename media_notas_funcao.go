// Função também é conhecida com procedimento ou sub-rotina
// É uma parte do código
// Ela pega o dado de uma entrada e dá um dado de saída

package main

import "fmt"

func media(notas []float64) float64 {

	total := 0.0

	for _, valor := range notas {

		total += valor
	}

	return total / float64(len(notas)) // interrompe a função

}

func main() {

	notas := []float64{10, 8.3, 7.5, 8.9, 9.1, 8.8}
	fmt.Println(media(notas))

}

/*
func main() {

	notas := []float64{10, 8.3, 7.5, 8.9, 9.1, 8.8}

	total := 0.0

	for _, valor := range notas {

		total += valor

	}
	fmt.Println(total / float64(len(notas)))

}
*/
