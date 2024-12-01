// Select funciona como um Switch para canais;
// Select permite que você aguarde operações de vários canais;

package main

// Cada canal receberá um valor após algum tempo, para simular, por exemplo, o bloqueio de operações RPC em execução em goroutines concorrentes.

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Um"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Dois"
	}()

	for i := 0; i < 2; i++ {
		select { // O select será usado para aguardar esses dois valores simultaneamente, imprimindo cada um à medida que chega.
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)
		}

	}
}
