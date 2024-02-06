package internal

// RepositoryInvoice is the interface that wraps the basic methods that an invoice repository should implement.
type RepositoryInvoice interface {
	// FindAll returns all invoices
	FindAll() (i []Invoice, err error)
	// Save saves an invoice
	Save(i *Invoice) (err error)
	// UpdateInvoicesTotal updates the total of an invoice
	UpdateInvoicesTotal() (invoiceUpdated int64, err error)
}
