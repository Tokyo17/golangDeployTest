package main

import (
	"encoding/json"
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

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	data := []struct {
		Name string
		Age  int
	}{
		{"Richard Grayson", 24},
		{"Jason Todd", 23},
		{"Tim Drake", 22},
		{"Damian Wayne", 21},
	}
	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		w.Write([]byte("post"))
	case "GET":

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInBytes)
	default:
		http.Error(w, "", http.StatusBadRequest)
	}

}
