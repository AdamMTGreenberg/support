package main

import (
	"fmt"
	"net/http"
    "time"

    "github.com/mitchellh/mapstructure"
)

type Message struct {
    Name string `json:"name`
    Data interface{} `json:"data`
}

type Channel struct {
    Name string `json:"name`
    Id string `json:"id`
}

func main() {
    router := NewRouter()

    router.Handle("channel add", addChannel)

	http.HandleFunc("/", router)
	http.ListenAndServe(":4000", nil)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "Hello from go")
// 	// var socket *websocket.Conn
// 	// var err error

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for {
// 		// msgType, msg, err := socket.ReadMessage()
//   //       if err != nil {
//   //           fmt.Println(err)
//   //           return
//   //       }
//         var inMessage Message
//         var outMessage Message

//         if err := socket.ReadJSON(&inMessage); err != nil {
//             fmt.Println(err)
//             break
//         }
//         fmt.Printf("%#v\n", inMessage)
//         switch inMessage.Name {
//         case "channel add":
//             err := addChannel(inMessage.Data)
//             if err != nil {
//                 outMessage = Message{"error", err}
//                 if err := socket.WriteJSON(outMessage); err != nil {
//                     fmt.Println(err)
//                     break
//                 }
//             }
//         case "channel subscribe":
//             go subscribeChannel(socket)
//         }
//         // fmt.Println(string(msg))
//         // if err = socket.WriteMessage(msgType, msg); err != nil {
//         //     fmt.Println(err)
//         //     return
//         // }
// 	}
// }


}