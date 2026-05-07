package models

import "time"

// Invoice represents a billing invoice in the system.
type Invoice struct {
	ID         uint32     `db:"id"          json:"id"`
	CustomerID uint32     `db:"customer_id" json:"customer_id"`
	Amount     float64    `db:"amount"      json:"amount"`
	Status     string     `db:"status"      json:"status"`
	DueDate    time.Time  `db:"due_date"    json:"due_date"`
	CreatedAt  time.Time  `db:"created_at"  json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"  json:"updated_at,omitempty"`
}
