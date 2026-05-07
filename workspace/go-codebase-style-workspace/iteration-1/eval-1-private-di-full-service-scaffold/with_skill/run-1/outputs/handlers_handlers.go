package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/blkst8/invoice-service/internal/app"
	"github.com/blkst8/invoice-service/internal/log"
	"github.com/blkst8/invoice-service/internal/models"
)

// Handlers holds the service bundle and exposes all handler methods.
type Handlers struct {
	svc *app.Service
}

// New constructs a new Handlers instance.
func New(svc *app.Service) *Handlers {
	return &Handlers{svc: svc}
}

// Healthz returns a simple health check response.
func Healthz(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

// CreateInvoice handles POST /api/v1/invoices.
func (h *Handlers) CreateInvoice(ctx echo.Context) error {
	var req models.Invoice
	if err := ctx.Bind(&req); err != nil {
		log.Logger.Error("failed to bind create invoice request", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := h.svc.Invoice.Create(ctx.Request().Context(), req); err != nil {
		log.Logger.Error("failed to create invoice", zap.Error(err))
		return err
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

// GetInvoice handles GET /api/v1/invoices/:id.
func (h *Handlers) GetInvoice(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid invoice id")
	}

	invoice, err := h.svc.Invoice.Get(ctx.Request().Context(), uint32(id))
	if err != nil {
		log.Logger.Error("failed to get invoice", zap.Uint64("id", id), zap.Error(err))
		return err
	}

	return ctx.JSON(http.StatusOK, invoice)
}

// ListInvoices handles GET /api/v1/invoices.
func (h *Handlers) ListInvoices(ctx echo.Context) error {
	invoices, err := h.svc.Invoice.List(ctx.Request().Context())
	if err != nil {
		log.Logger.Error("failed to list invoices", zap.Error(err))
		return err
	}

	return ctx.JSON(http.StatusOK, invoices)
}
