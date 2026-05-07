package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/example/invoice-service/internal/models"
	"github.com/example/invoice-service/internal/service"
)

// InvoiceHandler handles HTTP requests for invoices.
type InvoiceHandler struct {
	svc service.InvoiceService
}

// NewInvoiceHandler returns a new InvoiceHandler with the given service.
func NewInvoiceHandler(svc service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{svc: svc}
}

// Register mounts all invoice routes on the given Echo group.
func (h *InvoiceHandler) Register(g *echo.Group) {
	g.GET("/invoices", h.ListInvoices)
	g.GET("/invoices/:id", h.GetInvoice)
	g.POST("/invoices", h.CreateInvoice)
}

func (h *InvoiceHandler) ListInvoices(c echo.Context) error {
	invoices, err := h.svc.ListInvoices(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoices)
}

func (h *InvoiceHandler) GetInvoice(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid invoice id")
	}

	inv, err := h.svc.GetInvoice(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, inv)
}

func (h *InvoiceHandler) CreateInvoice(c echo.Context) error {
	var inv models.Invoice
	if err := c.Bind(&inv); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.svc.CreateInvoice(c.Request().Context(), &inv); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return c.JSON(http.StatusCreated, inv)
}
