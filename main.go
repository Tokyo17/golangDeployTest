package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handel)
	fmt.Println("server started localhost:8000")
	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, nil)
	fmt.Println(port)
}

func handel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	switch r.Method {
	case "POST":
		w.Write([]byte("post"))
	case "GET":
		w.Write([]byte("get"))
	default:
		http.Error(w, "", http.StatusBadRequest)
	}

}
