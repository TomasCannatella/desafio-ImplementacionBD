package repository

import (
	"database/sql"

	"app/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*c).Id = int(id)

	return
}

// FindTotalAmountGroupByCondition returns the total amount of customers group by condition.
func (r *CustomersMySQL) FindTotalAmountGroupByCondition() (c []internal.CustomerTotalAmountGroupByCondition, err error) {
	query := "SELECT if(c.condition=0,'Inactivo','Activo') as `condition`, sum(i.total) as `total` FROM customers c INNER JOIN invoices i ON c.id = i.customer_id GROUP BY c.condition;"
	rows, err := r.db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var ct internal.CustomerTotalAmountGroupByCondition
	for rows.Next() {
		err = rows.Scan(&ct.Condition, &ct.TotalAmount)
		if err != nil {
			return
		}
		c = append(c, ct)
	}
	return
}

// FindActiveWithHighestAmountSpent returns the customers that spent the highest amount of money in purchases.
func (r *CustomersMySQL) FindActiveWithHighestAmountSpent() (c []internal.CustomerActiveWithHighestAmountSpent, err error) {
	query := "SELECT c.id,c.first_name,c.last_name,SUM(i.total) as amount FROM customers c INNER JOIN invoices i ON c.id = i.customer_id GROUP BY c.id,c.first_name,c.last_name ORDER BY SUM(i.total) DESC"

	rows, err := r.db.Query(query)
	if err != nil {
		return
	}

	var ca internal.CustomerActiveWithHighestAmountSpent
	var id int
	for rows.Next() {
		err = rows.Scan(&id, &ca.FirstName, &ca.LastName, &ca.TotalAmount)
		if err != nil {
			return
		}
		c = append(c, ca)
	}
	return
}
