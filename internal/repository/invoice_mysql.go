package repository

import (
	"database/sql"

	"app/internal"
)

// NewInvoicesMySQL creates new mysql repository for invoice entity.
func NewInvoicesMySQL(db *sql.DB) *InvoicesMySQL {
	return &InvoicesMySQL{db}
}

// InvoicesMySQL is the MySQL repository implementation for invoice entity.
type InvoicesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all invoices from the database.
func (r *InvoicesMySQL) FindAll() (i []internal.Invoice, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var iv internal.Invoice
		// scan the row into the invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		// append the invoice to the slice
		i = append(i, iv)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the invoice into the database.
func (r *InvoicesMySQL) Save(i *internal.Invoice) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
		(*i).Datetime, (*i).Total, (*i).CustomerId,
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
	(*i).Id = int(id)

	return
}

// UpdateInvoicesTotal updates the total of an invoice.
func (r *InvoicesMySQL) UpdateInvoicesTotal() (invoiceUpdated int64, err error) {

	query := `
		UPDATE invoices i1 
		SET total = (SELECT SUM(s.quantity * p.price) AS total 
					FROM (SELECT i.id 
							FROM invoices i WHERE i.id = i1.id) i 
					INNER JOIN sales s ON i.id = s.invoice_id 
					INNER JOIN products p ON s.product_id = p.id);
	`

	result, err := r.db.Exec(query)
	if err != nil {
		return
	}

	invoiceUpdated, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}
