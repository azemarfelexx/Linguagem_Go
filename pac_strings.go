//vou utilizar a função contains / que procura dentro de uma string outra string menor
// Ex.: procurar dentro da palavra terra se ela tem, dois r

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Azemar", "ze"))
	fmt.Println(strings.Contains("Terra", "rs"))

}
