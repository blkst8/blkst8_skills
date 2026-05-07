package service

import (
	"context"
	"fmt"
	"log"

	"github.com/example/invoice-service/internal/models"
	"github.com/example/invoice-service/internal/repository"
)

// InvoiceService defines business logic operations for invoices.
type InvoiceService interface {
	GetInvoice(ctx context.Context, id int64) (*models.Invoice, error)
	ListInvoices(ctx context.Context) ([]*models.Invoice, error)
	CreateInvoice(ctx context.Context, invoice *models.Invoice) error
	ReconcileInvoices(ctx context.Context) error
}

type invoiceService struct {
	repo repository.InvoiceRepository
}

// NewInvoiceService returns a new InvoiceService with the given repository.
func NewInvoiceService(repo repository.InvoiceRepository) InvoiceService {
	return &invoiceService{repo: repo}
}

func (s *invoiceService) GetInvoice(ctx context.Context, id int64) (*models.Invoice, error) {
	inv, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting invoice: %w", err)
	}
	return inv, nil
}

func (s *invoiceService) ListInvoices(ctx context.Context) ([]*models.Invoice, error) {
	invoices, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing invoices: %w", err)
	}
	return invoices, nil
}

func (s *invoiceService) CreateInvoice(ctx context.Context, invoice *models.Invoice) error {
	if invoice.CustomerID == 0 {
		return fmt.Errorf("customer_id is required")
	}
	if invoice.Amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}
	if invoice.Status == "" {
		invoice.Status = "pending"
	}
	if err := s.repo.Create(ctx, invoice); err != nil {
		return fmt.Errorf("creating invoice: %w", err)
	}
	return nil
}

func (s *invoiceService) ReconcileInvoices(ctx context.Context) error {
	invoices, err := s.repo.ListUnreconciled(ctx)
	if err != nil {
		return fmt.Errorf("fetching unreconciled invoices: %w", err)
	}

	for _, inv := range invoices {
		// Business logic: mark paid invoices as reconciled
		if inv.Status == "paid" {
			if err := s.repo.MarkReconciled(ctx, inv.ID); err != nil {
				log.Printf("failed to reconcile invoice %d: %v", inv.ID, err)
				continue
			}
			log.Printf("reconciled invoice %d", inv.ID)
		}
	}

	return nil
}
