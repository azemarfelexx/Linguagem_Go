package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type biblioteca struct {
	Livros []livros `json:"livros"`
}

type livros struct {
	Titulo        string     `json:"Titulo"`
	Autor         string     `json:"Autor"`
	Paginas       int        `json:"Paginas"`
	Editora       string     `json:"Editora"`
	Nacionalidade string     `json:"Nacionalidade"`
	Disponivel    Disponivel `json:"Disponivel"`
	Formato       Formato    `json:"Formato"`
}

type Disponivel struct {
	Skoob  string `json:"Skoob"`
	Amazon string `json:"Amazon"`
}

type Formato struct {
	Digital string `json:"Digital"`
	Fisico  string `json:"Fisico"`
}

func main() {
	jsonFile, err := os.Open("biblioteca.json")
	//se der erro ao abrir o arquivo:
	if err != nil {
		fmt.Println(err)
	}
	// se não tiver erro:
	fmt.Println("Arquivo aberto com sucesso")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var livros biblioteca

	json.Unmarshal(byteValue, &livros)

	err = json.Unmarshal(byteValue, &livros)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	for i := 0; i < len(livros.Livros); i++ {
		fmt.Println("livros Titulo: " + livros.Livros[i].Titulo)
		fmt.Println("livros Paginas: " + strconv.Itoa(livros.Livros[i].Paginas))
	}
}
