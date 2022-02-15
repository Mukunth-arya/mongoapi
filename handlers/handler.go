package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"name/controller"
	"name/models"
)

func getalldata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	dataget := controller.Getallsource()

	json.NewEncoder(w).Encode(dataget)
}

func postvalue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	if r.Body == nil {

		json.NewEncoder(w).Encode("Hey please enter the value")

	}
	var post models.Data
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {

		log.Fatal(err)
	}
	controller.Insertdataone(post)
	json.NewEncoder(w).Encode(post)
}
func markupto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	controller.Efficentparameter(params["id"])
	json.NewEncoder(w).Encode(params)
}
func deletesingle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	controller.Deleteone(params["id"])
	json.NewEncoder(w).Encode(params)
}
func deleteall(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	value := controller.DeleteAll()
	json.NewEncoder(w).Encode(value)
}
