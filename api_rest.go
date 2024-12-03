// http://pokeapi.co/api/v2/pokedex/kanto/
// GET: traz uma listagem de registros
// POST: Adiciona um novo registro
// DELETE: remove um registro
// PUT e PAT: Editar um registro
// Vamos trabalhar com 3 estruturas: 1: Response, 2: Pokemon, 3: Pokemon Specie

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Estruturar:

type response struct {
	Nome    string    `json:"name"`
	Pokemon []pokemon `json:"pokemon_entries"`
}

type pokemon struct {
	Numero  int            `json:"entry_number"`
	Especie pokemonSpecies `json:"pokemonSpecies"`
}

type pokemonSpecies struct {
	Nome string `json:"name"`
}

func main() {
	httpResponse, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/") // mapear o site

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(httpResponse.Body) // dados --> etão em bytes
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData)) // transforma os dados em strings

	var responseObject response //Criar a variável responseObject para trabalhar com o responseData, para decodificar o json
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Nome)
	fmt.Println(responseObject.Pokemon)

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Especie.Nome)
	}
}
