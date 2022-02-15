package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Mukunth-arya/mongoapi/handler"
)

func main() {

	fmt.Println("Here we are gooing to consruct the new api")
	fmt.Println("Api for review system")
	r := mux.NewRouter()
	r.HandleFunc("/getnew", handler.Getalldata).Methods("GET")
	r.HandleFunc("/getnew/addsome", handler.Postvalue).Methods("POST")

	http.ListenAndServe(":2000", r)

}
