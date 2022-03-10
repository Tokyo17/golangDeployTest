package websocket

import (
	"fmt"
	"net/http"
	"os"
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
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello api")
}
