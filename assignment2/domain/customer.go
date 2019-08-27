
package domain

type Customer struct {
	ID string
	Name string
	Email string
}

type CustomerStore interface {
	Create(Customer) error
	Delete(string) error
	Update(string, Customer) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}
