package router

import "Golang_assignment/assignment3/Controller"
import	"github.com/gorilla/mux"


var Mux = mux.NewRouter()


func Routers() {

	Mux.HandleFunc("/create/", controller.Create).Methods("GET")
	Mux.HandleFunc("/update/", controller.Update).Methods("GET")
	Mux.HandleFunc("/getbyid/", controller.GetById).Methods("GET")
	Mux.HandleFunc("/getall/", controller.GetAll).Methods("GET")
	Mux.HandleFunc("/delete/", controller.Delete).Methods("GET")

}
