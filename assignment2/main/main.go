package main

import (
	"assignment2/domain"
	"assignment2/mapstore"

	//"assignment2/mapstore"
	"fmt"
)
	type CustomerController struct {

		store domain.CustomerStore
	}

	func(cc *CustomerController) Create(c domain.Customer) {

		err := cc.store.Create(c)
		if err != nil {
			fmt.Println("error", err)
		} else {
			fmt.Println("user created")
		}

	}

	func(cc *CustomerController) update(id string,c domain.Customer) {
		err := cc.store.Update(id,c)
		if err != nil {
			fmt.Println("error", err)
		} else {
			fmt.Println("user updated")
		}
	}

	func (cc CustomerController) Delete(c domain.Customer) {
	err := cc.store.Delete(c.ID)
	if err != nil {
		fmt.Println("Could not delete a customer:", err)
		return
	}
	fmt.Println("User deleted Successfully")
}
func (cc CustomerController) getAll () {
	val,err := cc.store.GetAll()
	if err != nil {
		fmt.Println("no record found")
		return
	}
	for _,i := range val{
		fmt.Println(i)
	}

}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}
	customer := domain.Customer{
		ID:   "1",
		Name: "Anupriya Gupta",
	}
	customer2:= domain.Customer{
		ID:   "1",
		Name: "Udit",
	}

	controller.Create(customer)
	controller.getAll()
	controller.update("1",customer2)
	controller.Delete(customer)
	controller.getAll()

}