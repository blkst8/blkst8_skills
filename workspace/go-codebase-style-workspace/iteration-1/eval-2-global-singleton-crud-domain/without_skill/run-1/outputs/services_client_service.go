package services

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/blkst8/client-service/internal/models"
	"github.com/blkst8/client-service/internal/repository"
)

// ClientService defines business logic for the client domain.
type ClientService interface {
	Create(ctx context.Context, client models.Client) error
	Get(ctx context.Context, id uint32) (*models.Client, error)
	Update(ctx context.Context, client models.Client) error
	Delete(ctx context.Context, id uint32) error
}

type clientService struct {
	db   *sqlx.DB
	repo repository.Client
}

// NewClientService constructs a ClientService.
func NewClientService(db *sqlx.DB, repo repository.Client) ClientService {
	return &clientService{db: db, repo: repo}
}

func (s *clientService) Create(ctx context.Context, client models.Client) error {
	client.CreatedAt = time.Now()
	return s.repo.Create(ctx, client)
}

func (s *clientService) Get(ctx context.Context, id uint32) (*models.Client, error) {
	return s.repo.Get(ctx, id)
}

func (s *clientService) Update(ctx context.Context, client models.Client) error {
	now := time.Now()
	client.UpdatedAt = &now
	return s.repo.Update(ctx, client)
}

func (s *clientService) Delete(ctx context.Context, id uint32) error {
	return s.repo.Delete(ctx, id)
}
