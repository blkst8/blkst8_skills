package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/blkst8/client-service/internal/models"
)

// ErrClientNotFound is returned when a client cannot be found.
var ErrClientNotFound = errors.New("client not found")

// Client defines the data access interface for the client domain.
type Client interface {
	Create(ctx context.Context, client models.Client) error
	Get(ctx context.Context, id uint32) (*models.Client, error)
	Update(ctx context.Context, client models.Client) error
	Delete(ctx context.Context, id uint32) error
}

type clientRepository struct {
	db *sqlx.DB
}

// NewClientRepository returns a new sqlx-backed Client repository.
func NewClientRepository(db *sqlx.DB) Client {
	return &clientRepository{db: db}
}

func (r *clientRepository) Create(ctx context.Context, client models.Client) error {
	query := `INSERT INTO clients (name, email, phone, created_at) VALUES (:name, :email, :phone, :created_at)`
	_, err := r.db.NamedExecContext(ctx, query, &client)
	return err
}

func (r *clientRepository) Get(ctx context.Context, id uint32) (*models.Client, error) {
	var client models.Client
	if err := r.db.GetContext(ctx, &client, `SELECT * FROM clients WHERE id = ?`, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) Update(ctx context.Context, client models.Client) error {
	query := `UPDATE clients SET name = :name, email = :email, phone = :phone, updated_at = :updated_at WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, &client)
	return err
}

func (r *clientRepository) Delete(ctx context.Context, id uint32) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM clients WHERE id = ?`, id)
	return err
}
