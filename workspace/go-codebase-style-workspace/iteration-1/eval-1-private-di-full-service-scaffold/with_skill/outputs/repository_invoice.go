package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/blkst8/invoice-service/internal/models"
)

// ErrInvoiceNotFound is returned when a requested invoice does not exist.
var ErrInvoiceNotFound = errors.New("invoice not found")

// Invoice defines the data access interface for invoices.
type Invoice interface {
	Create(ctx context.Context, invoice models.Invoice) error
	Get(ctx context.Context, id uint32) (*models.Invoice, error)
	List(ctx context.Context) ([]models.Invoice, error)
	Update(ctx context.Context, invoice models.Invoice) error
	ListPending(ctx context.Context) ([]models.Invoice, error)
}

type invoiceRepository struct {
	db *sqlx.DB
}

// NewInvoiceRepository constructs a new Invoice repository implementation.
func NewInvoiceRepository(db *sqlx.DB) Invoice {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) Create(ctx context.Context, invoice models.Invoice) error {
	query := `INSERT INTO invoices (customer_id, amount, status, due_date, created_at)
	          VALUES (:customer_id, :amount, :status, :due_date, :created_at)`
	_, err := r.db.NamedExecContext(ctx, query, &invoice)
	return err
}

func (r *invoiceRepository) Get(ctx context.Context, id uint32) (*models.Invoice, error) {
	var invoice models.Invoice
	query := `SELECT * FROM invoices WHERE id = ?`
	if err := r.db.GetContext(ctx, &invoice, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvoiceNotFound
		}
		return nil, err
	}
	return &invoice, nil
}

func (r *invoiceRepository) List(ctx context.Context) ([]models.Invoice, error) {
	var invoices []models.Invoice
	query := `SELECT * FROM invoices ORDER BY created_at DESC`
	if err := r.db.SelectContext(ctx, &invoices, query); err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *invoiceRepository) Update(ctx context.Context, invoice models.Invoice) error {
	query := `UPDATE invoices SET status = :status, updated_at = :updated_at WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, &invoice)
	return err
}

func (r *invoiceRepository) ListPending(ctx context.Context) ([]models.Invoice, error) {
	var invoices []models.Invoice
	query := `SELECT * FROM invoices WHERE status = 'pending' ORDER BY due_date ASC`
	if err := r.db.SelectContext(ctx, &invoices, query); err != nil {
		return nil, err
	}
	return invoices, nil
}
