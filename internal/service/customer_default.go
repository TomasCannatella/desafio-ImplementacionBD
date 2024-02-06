package service

import "app/internal"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

// FindTotalAmountGroupByCondition returns the total amount of customers group by condition.
func (s *CustomersDefault) FindTotalAmountGroupByCondition() (c []internal.CustomerTotalAmountGroupByCondition, err error) {
	c, err = s.rp.FindTotalAmountGroupByCondition()
	return
}

// FindActiveWithHighestAmountSpent returns the customers that spent the highest amount of money in purchases.
func (s *CustomersDefault) FindActiveWithHighestAmountSpent() (c []internal.CustomerActiveWithHighestAmountSpent, err error) {
	c, err = s.rp.FindActiveWithHighestAmountSpent()
	return
}
