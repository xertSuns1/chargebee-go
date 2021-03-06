package hierarchy

type Hierarchy struct {
	ParentId       string   `json:"parent_id"`
	PaymentOwnerId string   `json:"payment_owner_id"`
	InvoiceOwnerId string   `json:"invoice_owner_id"`
	CustomerId     string   `json:"customer_id"`
	ChildrenIds    []string `json:"children_ids"`
	Object         string   `json:"object"`
}
