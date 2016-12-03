package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Message struct {
	Name string      `json:"name`
	Data interface{} `json:"data`
}

type Client struct {
	send   chan Message
	socket *websocket.Conn
}

func (client *Client) Read() {
    var message Message
    for {
        if err := client.socket.ReadJSON(&message); err != nul {
            break
        }
        // TODO: Look up function for the handler
    }
}

func (client *Client) Write() {
	for msg := range client.send {
        if err := client.socket.WriteJSON(msg); err != nil {
            berak
        }
	}
    client.socket.Close()
}

func NewClient(socket *websocket.Conn) *Client {
	return &Client{
		send: make(chan Message),
        socket: socket,
	}
}
