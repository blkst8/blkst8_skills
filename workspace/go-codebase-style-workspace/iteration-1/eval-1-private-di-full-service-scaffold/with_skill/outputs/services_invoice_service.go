package services

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/blkst8/invoice-service/internal/log"
	"github.com/blkst8/invoice-service/internal/models"
	"github.com/blkst8/invoice-service/internal/repository"
)

// InvoiceService defines the business logic interface for invoices.
type InvoiceService interface {
	Create(ctx context.Context, invoice models.Invoice) error
	Get(ctx context.Context, id uint32) (*models.Invoice, error)
	List(ctx context.Context) ([]models.Invoice, error)
	Reconcile(ctx context.Context) error
}

type invoiceService struct {
	db   *sqlx.DB
	repo repository.Invoice
}

// NewInvoiceService constructs a new InvoiceService implementation.
func NewInvoiceService(db *sqlx.DB, repo repository.Invoice) InvoiceService {
	return &invoiceService{db: db, repo: repo}
}

func (s *invoiceService) Create(ctx context.Context, invoice models.Invoice) error {
	invoice.Status = "pending"
	invoice.CreatedAt = time.Now()
	return s.repo.Create(ctx, invoice)
}

func (s *invoiceService) Get(ctx context.Context, id uint32) (*models.Invoice, error) {
	return s.repo.Get(ctx, id)
}

func (s *invoiceService) List(ctx context.Context) ([]models.Invoice, error) {
	return s.repo.List(ctx)
}

// Reconcile scans all pending invoices and marks those past their due date as
// overdue. It is intended to be called by the background worker every 5 minutes.
func (s *invoiceService) Reconcile(ctx context.Context) error {
	pending, err := s.repo.ListPending(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	for _, invoice := range pending {
		if now.After(invoice.DueDate) {
			invoice.Status = "overdue"
			invoice.UpdatedAt = &now
			if err := s.repo.Update(ctx, invoice); err != nil {
				log.Logger.Error("failed to mark invoice as overdue",
					zap.Uint32("invoice_id", invoice.ID),
					zap.Error(err),
				)
			}
		}
	}

	return nil
}
