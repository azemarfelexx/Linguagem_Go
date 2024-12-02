// Criar um site/servidor estático
// O que o código fará: Servir arquivos html de um local específico no disco.

// 1º Passo: Criar um diretório (uma pasta) para armazenar o projeto;
// 2º Passo: Criar um arquivo html - arquivo.html (cdm: type nul > arquivo.html ou cdm: type arquivo.html)
// 3º Passo: Abrir os arquivos e escrover os códigos.

package main

// Importar os pacotes que serão usados:
import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./estatico"))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requisição recebida: %s", r.URL.Path)
		fs.ServeHTTP(w, r)
	}))
	log.Print("Listening on: 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Para acessar: http://localhost:3000/arquivo.html
