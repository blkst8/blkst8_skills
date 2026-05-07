package app

import (
	"github.com/jmoiron/sqlx"

	"github.com/blkst8/invoice-service/internal/services"
)

// Service holds all service instances for the application.
type Service struct {
	Invoice services.InvoiceService
}

// WithServices constructs all services and returns the bundle.
func WithServices(db *sqlx.DB, repo *Repository) *Service {
	return &Service{
		Invoice: services.NewInvoiceService(db, repo.Invoice),
	}
}
