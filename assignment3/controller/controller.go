package controller

import (
	"fmt"
	"Golang_assignment/assignment3/domain"
	"Golang_assignment/assignment3/mapstore"
	"net/http"
	"github.com/gorilla/mux"
)

type CustomerController struct {
	store domain.CustomerStore
}

var controller = CustomerController{
	store: mapstore.NewMapStore(),
}

func Create(w http.ResponseWriter, r *http.Request) {

	    cust := domain.Customer{
		ID:   "1629185",
		Name: "Udit Gaurav",
		Email:"udit.gaurav@mayadata.io",
	}

	err := controller.store.Create(cust)
	if err != nil {
		fmt.Println("Could not create the customer!!")
		}else {
		fmt.Println("Customer  created successfuly!!")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	    updateCustomer := domain.Customer{
		ID:   "213",
		Name: "Udit Gaurav",
		Email:"udit.gaurav@mayadata.io",
	}
	err := controller.store.Update("1629185", updateCustomer)
	if err != nil {
		fmt.Println("Update Failed!Try again")
	}else {
		fmt.Println("Deatils Updated", updateCustomer)
	}

}
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	del := vars["1629185"]
	err := controller.store.Delete(del)
	if err != nil {
		fmt.Println("Deletion Failed")

	}else {
		fmt.Println("DELETED!!")
	}

}

func GetById(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["1629185"]
	//fmt.Fprintf(w,"GetbyID ")
	value, err := controller.store.GetById(k)
	if err != nil {
		fmt.Println("Incorrect!check Id ")

	}else {
		fmt.Println(value)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("Getall executed")
	  value, err := controller.store.GetAll()
	if err != nil {
		fmt.Println("ERROR!TRY AGAIN")

	}else {
		fmt.Println( value)
	}

}
