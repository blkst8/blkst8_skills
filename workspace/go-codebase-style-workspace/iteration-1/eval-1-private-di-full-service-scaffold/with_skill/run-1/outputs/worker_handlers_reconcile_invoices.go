package handlers

import (
	"context"

	"go.uber.org/zap"

	"github.com/blkst8/invoice-service/internal/app"
	"github.com/blkst8/invoice-service/internal/log"
)

type reconcileInvoices struct {
	svc *app.Service
}

// NewReconcileInvoices constructs the reconcile invoices job handler.
func NewReconcileInvoices(svc *app.Service) *reconcileInvoices {
	return &reconcileInvoices{svc: svc}
}

// Handle runs the invoice reconciliation job. It is called by the worker on
// each tick and marks overdue invoices accordingly.
func (h *reconcileInvoices) Handle(ctx context.Context) {
	log.Logger.Info("starting invoice reconciliation job")

	if err := h.svc.Invoice.Reconcile(ctx); err != nil {
		log.Logger.Error("invoice reconciliation failed", zap.Error(err))
		return
	}

	log.Logger.Info("invoice reconciliation completed")
}
