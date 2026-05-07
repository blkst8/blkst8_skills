package models

import "time"

// Client represents a client in the system.
type Client struct {
	ID        uint32     `db:"id"         json:"id"`
	Name      string     `db:"name"       json:"name"`
	Email     string     `db:"email"      json:"email"`
	Phone     string     `db:"phone"      json:"phone"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
