package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Router struct {
	rules map[string]Handler
}

type Handler func(*Client, interface{})

func (e *Rounter) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printt(w, err.Error())
		return
	}
	client := NewClient(socket, e.FindHandler)
	go client.Write()
	client.Read()
}

func (r *Router) FindHandler(msgName string) (Handler, bool) {
    handler, found :=r.rules[msgName]
    return handler, found
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
