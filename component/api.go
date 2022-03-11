package fungsi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Name struct {
	Name string
	Age  int
}

func Api() {
	http.HandleFunc("/api", ApiTest)
	http.HandleFunc("/postapi", postApi)
	http.HandleFunc("/getapi", getApi)

	http.ListenAndServe(":8080", nil)
}

func ApiTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	var method = fmt.Sprintf("ini %s", r.Method)
	fmt.Fprintf(w, method)
}

func postApi(w http.ResponseWriter, r *http.Request) {
	//# mengambil data dari body tapi dalam byte
	w.Header().Set("Content-Type", "application/json")

	read, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error)
	}

	//#  merubah data ke string dan memasukan data body tadi ke struct
	var names []Name
	json.Unmarshal(read, &names)
	fmt.Println(names)

	//# menampilkan data
	json.NewEncoder(w).Encode(names)
}

func getApi(w http.ResponseWriter, r *http.Request) {

	// fmt.Println(key)
	// json.NewEncoder(w).Encode(key)
}
