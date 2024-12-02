package main

import (
	"fmt"
	"time"
)

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)

	go pong(c)
	go ping(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
