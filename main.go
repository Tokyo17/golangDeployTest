package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handel)
	fmt.Println("server started localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func handel(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Write([]byte("post"))
	case "GET":
		w.Write([]byte("get"))
	default:
		http.Error(w, "", http.StatusBadRequest)
	}

}
