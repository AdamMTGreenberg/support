package main

import (
	"fmt"
)

type Message struct {
	Name string      `json:"name`
	Data interface{} `json:"data`
}

func main() {
	msgChan := make(chan string)
	go func() {
		msgChan <- "Hello"
	}()
	msg := <-msgChan
	fmt.Println(msg)
}
