package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Mukunth-arya/mongoapi/router"
)

func main() {

	fmt.Println("Here we are gooing to consruct the new api")
	fmt.Println("Api for review system")
	r := mux.NewRouter()
	r.HandleFunc("/getnew", router.Getalldata).Methods("GET")
	r.HandleFunc("/getnew/addsome", router.Postvalue).Methods("POST")

	http.ListenAndServe(":7000", r)

}
