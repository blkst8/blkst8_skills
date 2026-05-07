package worker

import (
	"context"
	"log"
	"time"

	"github.com/example/invoice-service/internal/service"
)

const reconcileInterval = 5 * time.Minute

// ReconcileWorker runs the invoice reconciliation job on a ticker.
type ReconcileWorker struct {
	svc service.InvoiceService
}

// NewReconcileWorker returns a new ReconcileWorker with the given service.
func NewReconcileWorker(svc service.InvoiceService) *ReconcileWorker {
	return &ReconcileWorker{svc: svc}
}

// Start runs the reconciliation loop until the context is cancelled.
func (w *ReconcileWorker) Start(ctx context.Context) {
	log.Printf("reconcile worker started (interval: %s)", reconcileInterval)
	ticker := time.NewTicker(reconcileInterval)
	defer ticker.Stop()

	// Run once immediately on start
	w.run(ctx)

	for {
		select {
		case <-ticker.C:
			w.run(ctx)
		case <-ctx.Done():
			log.Println("reconcile worker stopped")
			return
		}
	}
}

func (w *ReconcileWorker) run(ctx context.Context) {
	log.Println("starting invoice reconciliation")
	if err := w.svc.ReconcileInvoices(ctx); err != nil {
		log.Printf("reconciliation error: %v", err)
		return
	}
	log.Println("invoice reconciliation complete")
}
