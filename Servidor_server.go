package main

import (
	"fmt"
	"net/http"
)

// Um conceito fundamental sem servidores net/http são funções (que estão guardados em nosso pacote http)

func ola(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ola\n")
}

// As funções que servem como manipuladores recebem a http.RespondeWriter e a http.Request como argumentos.
// O gravador de resposta é usado para preencher a resposta http. Aqui, a resposta simples é apenas "Olá\n".

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

// Esse manipulador faz algo um pouco mais sofisticado, lendo todos os cabeçalhos de solicitação http e ecoando-os no corpo da resposta.

func main() {
	http.HandleFunc("/Ola", ola)
	http.HandleFunc("/cabecalho", cabecalhos)
	http.ListenAndServe(":8090", nil)
}

/* Um manipulador ("função") é um objeto que implementa http. Handler.
Uma maneira comum de escrever um manipulador (handler) é usar o http.HandleFunc adaptando as funções com a assinatura própria.
Registramos nossos manipuladores nas rotas do servidor usando a http.HandleFunc, rota da função que ele dee chamar: "/ola" e "/cabecalho".
Finalmente, chamamos o ListenAndServe com a porta ":08090" e um manipulador nil para que seja usado o roteador padrão que acabamos de configurar

Acesse http://localhost:8090/Ola
Acesse http://localhost:8090/cabecalho
*/
