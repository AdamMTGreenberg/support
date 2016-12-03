package main

import (
	"net/http"
)

type Channel struct {
    Name string `json:"name`
    Id string `json:"id`
}

func main() {
    router := NewRouter()

    router.Handle("channel add", addChannel)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}