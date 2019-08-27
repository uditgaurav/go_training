package mapstore

import (
	"assignment2/domain"
	"errors"
	"fmt"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

// NOTE : Other validations/ edge cases are not covered for following methods

func (ms *MapStore) Create(c domain.Customer) error {
	if c.ID == "" {
		return errors.New("ID not found for the customer")
	}
	ms.store[c.ID] = c
	fmt.Println("User created successfully, User ID:", c.ID)
	return nil
}

func (ms *MapStore) Delete(ID string) error {
	_, ok := ms.store[ID]
	if ok {
		delete(ms.store, ID)
	//	fmt.Println("User deleted successfully")
		return nil
	}
	return errors.New("Cannot delete: the customer does not exist")
}

func (ms *MapStore) Update(ID string, c domain.Customer) error {
	_, ok := ms.store[ID]
	if ok {
			ms.store[ID] = c
			return nil
		  }
	return errors.New("can not update: the customer does not exist")

}

func (ms *MapStore) GetById(ID string) (domain.Customer, error) {
	c, ok := ms.store[ID]
	if ok {
		return c, nil
	}
	return domain.Customer{}, errors.New("given id doesn't exists")
}

func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	var cust []domain.Customer
	if len(ms.store) != 0 {
		for _, v := range ms.store {
			cust = append(cust, v)
		}
		return cust, nil
	}
	return nil, errors.New("No record found")
}

