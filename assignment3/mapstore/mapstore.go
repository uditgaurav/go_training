package mapstore

import "Golang_assignment/assignment3/domain"
import "errors"
//import "fmt"

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore { store: make(map[string]domain.Customer)}
}

func (m *MapStore)Create(c domain.Customer)error {
	if c.ID == "" {
		return errors.New("Customer Id not present!!")
	} else {
		m.store[c.ID] = c
		return nil
	}
}

func (m *MapStore) Update(ID string,c domain.Customer)error{
	if c.ID==""{
		return errors.New("Error! Try again adding valid key")
	}else{


		m.store[c.ID]=c

		return nil
	}
}

func( m *MapStore)GetAll()([]domain.Customer,error){
	var cust []domain.Customer
	for _,value :=range m.store{
		cust=append(cust,value)
	}
	return cust,nil
}

func (m *MapStore)GetById(s string)(domain.Customer,error){
	//var c domain.Customer
	//c=m.store[s]
	s1:=m.store[s]
	// fmt.Println("ID is",s1.ID)
	if s1.ID==""{
		return domain.Customer{},errors.New("Incorrect ID")
	}else{
		return m.store[s],nil
	}

}

func (m *MapStore)Delete(s string)error{
	c:=m.store[s]
	if c.ID==""{
		return errors.New("Not valid ID")
	}else{
		delete(m.store,s)
		return nil
	}

}

