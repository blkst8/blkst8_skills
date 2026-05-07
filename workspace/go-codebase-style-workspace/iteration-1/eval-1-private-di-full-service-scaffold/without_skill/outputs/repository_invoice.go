package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/example/invoice-service/internal/models"
)

// InvoiceRepository defines the data access interface for invoices.
type InvoiceRepository interface {
	GetByID(ctx context.Context, id int64) (*models.Invoice, error)
	List(ctx context.Context) ([]*models.Invoice, error)
	Create(ctx context.Context, invoice *models.Invoice) error
	Update(ctx context.Context, invoice *models.Invoice) error
	ListUnreconciled(ctx context.Context) ([]*models.Invoice, error)
	MarkReconciled(ctx context.Context, id int64) error
}

type invoiceRepository struct {
	db *sql.DB
}

// NewInvoiceRepository returns a new InvoiceRepository backed by the given database.
func NewInvoiceRepository(db *sql.DB) InvoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) GetByID(ctx context.Context, id int64) (*models.Invoice, error) {
	const query = `
		SELECT id, customer_id, amount, status, is_reconciled, created_at, updated_at
		FROM invoices
		WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, id)
	inv := &models.Invoice{}
	err := row.Scan(
		&inv.ID,
		&inv.CustomerID,
		&inv.Amount,
		&inv.Status,
		&inv.IsReconciled,
		&inv.CreatedAt,
		&inv.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("invoice %d not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("querying invoice: %w", err)
	}
	return inv, nil
}

func (r *invoiceRepository) List(ctx context.Context) ([]*models.Invoice, error) {
	const query = `
		SELECT id, customer_id, amount, status, is_reconciled, created_at, updated_at
		FROM invoices
		ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("listing invoices: %w", err)
	}
	defer rows.Close()

	var invoices []*models.Invoice
	for rows.Next() {
		inv := &models.Invoice{}
		if err := rows.Scan(
			&inv.ID,
			&inv.CustomerID,
			&inv.Amount,
			&inv.Status,
			&inv.IsReconciled,
			&inv.CreatedAt,
			&inv.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scanning invoice row: %w", err)
		}
		invoices = append(invoices, inv)
	}
	return invoices, rows.Err()
}

func (r *invoiceRepository) Create(ctx context.Context, invoice *models.Invoice) error {
	const query = `
		INSERT INTO invoices (customer_id, amount, status, is_reconciled, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)`

	now := time.Now()
	invoice.CreatedAt = now
	invoice.UpdatedAt = now

	result, err := r.db.ExecContext(ctx, query,
		invoice.CustomerID,
		invoice.Amount,
		invoice.Status,
		invoice.IsReconciled,
		invoice.CreatedAt,
		invoice.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("creating invoice: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("getting last insert id: %w", err)
	}
	invoice.ID = id
	return nil
}

func (r *invoiceRepository) Update(ctx context.Context, invoice *models.Invoice) error {
	const query = `
		UPDATE invoices
		SET customer_id = ?, amount = ?, status = ?, is_reconciled = ?, updated_at = ?
		WHERE id = ?`

	invoice.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, query,
		invoice.CustomerID,
		invoice.Amount,
		invoice.Status,
		invoice.IsReconciled,
		invoice.UpdatedAt,
		invoice.ID,
	)
	if err != nil {
		return fmt.Errorf("updating invoice: %w", err)
	}
	return nil
}

func (r *invoiceRepository) ListUnreconciled(ctx context.Context) ([]*models.Invoice, error) {
	const query = `
		SELECT id, customer_id, amount, status, is_reconciled, created_at, updated_at
		FROM invoices
		WHERE is_reconciled = FALSE
		ORDER BY created_at ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("listing unreconciled invoices: %w", err)
	}
	defer rows.Close()

	var invoices []*models.Invoice
	for rows.Next() {
		inv := &models.Invoice{}
		if err := rows.Scan(
			&inv.ID,
			&inv.CustomerID,
			&inv.Amount,
			&inv.Status,
			&inv.IsReconciled,
			&inv.CreatedAt,
			&inv.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scanning invoice row: %w", err)
		}
		invoices = append(invoices, inv)
	}
	return invoices, rows.Err()
}

func (r *invoiceRepository) MarkReconciled(ctx context.Context, id int64) error {
	const query = `UPDATE invoices SET is_reconciled = TRUE, updated_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("marking invoice %d reconciled: %w", id, err)
	}
	return nil
}
