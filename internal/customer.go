package internal

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

// CustomerTotalAmountGroupByCondition is the struct that represents the total amount of customers group by condition.
type CustomerTotalAmountGroupByCondition struct {
	// Condition is the condition of the customer.
	Condition string
	// TotalAmount is the total amount of customers.
	TotalAmount float64
}

// CustomerActiveWithHighestAmountSpent is the struct that represents the customers that spent the highest amount of money in purchases.
type CustomerActiveWithHighestAmountSpent struct {
	// Id is the unique identifier of the customer.
	Id int
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// TotalAmount is the total amount of customers.
	TotalAmount float64
}
