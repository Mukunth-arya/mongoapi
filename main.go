package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Here we are gooing to consruct the new api")
	fmt.Println("Api for review system")
	r := mux.NewRouter()
	r.HandleFunc("/data", datasample).Methods("GET")
	r.HandleFunc("/getnew", getalldata).Methods("GET")
	r.HandleFunc("/getnew/addsome", postvalue).Methods("POST")
	r.HandleFunc("/getnew/{id}", markupto).Methods("PUT")
	r.HandleFunc("/getnew/{id}", deletesingle).Methods("DELETE")
	r.HandleFunc("/getnew/addsome/delall", deleteall).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9090", r))

}
func datasample(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("heythere")

}
