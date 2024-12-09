// Doc: soma.go


package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 10)
	z := subtrai(30, 10, 5) // Exemplo com múltiplos valores
	a := divide(10, 2)
	fmt.Println(x, y, z, a)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func subtrai(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}

func divide(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, v := range i[1:] {
		if v == 0 {
			fmt.Println("Erro: Divisão por zero.")
			return 0
		}

		total /= v
	}
	return total
}


// Doc Soa_test.go:

package main

import "testing"

/* Padrão Triple A - AAA
A - Arrange (Preparar)
A - Act (Rodar)
A - Assert (Verificar as Asserções)
*/

func ShouldSumCorrect(t *testing.T) {
	//arrange
	teste := soma(3, 2, 1)
	//act
	resultado := 6
	// assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldSumIncorrect(t *testing.T) {
	teste := soma(3, 2, 1) // arrange
	resultado := 7         // act
	// assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldMultCorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldMultIncorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 265
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldSubCorrect(t *testing.T) {
	teste := subtrai(30, 10, 5)
	resultado := 15
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldSubInCorrect(t *testing.T) {
	teste := subtrai(30, 10, 5)
	resultado := 30
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldDivCorrect(t *testing.T) {
	teste := divide(10, 2)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func ShouldDivIncorrect(t *testing.T) {
	teste := divide(10, 2)
	resultado := 6
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

// Para executar uma esteira de testes, usar o comando: go test -v
