package websocket

import (
	"fmt"
	"net/http"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go Writer(ws)
	Reader(ws)
}
func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func WebsocketMain() {
	fmt.Println("websocket Started")
	setupRoutes()
	http.HandleFunc("/", Api)
	http.ListenAndServe(":8080", nil)
}

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello api")
}
