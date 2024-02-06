package internal

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	// FindAll returns all customers
	FindAll() (c []Customer, err error)
	// Save saves a customer
	Save(c *Customer) (err error)
	// FindTotalAmountGroupByCondition returns the total amount of customers group by condition
	FindTotalAmountGroupByCondition() (c []CustomerTotalAmountGroupByCondition, err error)
	// FindActiveWithHighestAmountSpent returns the active customer with the highest amount spent
	FindActiveWithHighestAmountSpent() (c []CustomerActiveWithHighestAmountSpent, err error)
}
