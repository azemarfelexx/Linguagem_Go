// Pacote hash:

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// Criar a hash
	h := crc32.NewIEEE()
	// Escrever os dados no hash
	h.Write([]byte("Código com o pacote hash"))
	// Calcular o hash
	v := h.Sum32()
	fmt.Println(v)
}

// Pacote Crypto:

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	h := sha1.New()
	h.Write([]byte("Código com o pacote crypto"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)

}
