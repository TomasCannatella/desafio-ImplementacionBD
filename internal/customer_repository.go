package internal

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)
	// FindTotalAmountGroupByCondition returns the total amount of customers group by condition.
	FindTotalAmountGroupByCondition() (c []CustomerTotalAmountGroupByCondition, err error)
	// FindActiveWithHighestAmountSpent returns the customers that spent the highest amount of money in purchases.
	FindActiveWithHighestAmountSpent() (c []CustomerActiveWithHighestAmountSpent, err error)
}
