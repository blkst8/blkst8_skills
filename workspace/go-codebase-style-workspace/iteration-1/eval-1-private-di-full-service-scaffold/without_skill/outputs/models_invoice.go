package models

import "time"

// Invoice represents an invoice record.
type Invoice struct {
	ID          int64     `db:"id" json:"id"`
	CustomerID  int64     `db:"customer_id" json:"customer_id"`
	Amount      float64   `db:"amount" json:"amount"`
	Status      string    `db:"status" json:"status"`
	IsReconciled bool     `db:"is_reconciled" json:"is_reconciled"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
