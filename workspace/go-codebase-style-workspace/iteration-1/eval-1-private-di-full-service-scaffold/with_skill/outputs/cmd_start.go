package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/blkst8/invoice-service/internal/app"
	"github.com/blkst8/invoice-service/internal/config"
	httpserver "github.com/blkst8/invoice-service/internal/http"
	"github.com/blkst8/invoice-service/internal/worker"
	workerhandlers "github.com/blkst8/invoice-service/internal/worker/handlers"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the invoice-service HTTP server and background workers",
	Run:   startFunc,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func startFunc(_ *cobra.Command, _ []string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	db := app.WithDatabase()
	defer db.Close()

	repo := app.WithRepository(db)
	svc := app.WithServices(db, repo)

	if config.C.Worker.Enabled {
		reconcileJob := worker.NewWorker(config.C.Worker.JobsIntervals.ReconcileInvoices)
		reconcileJob.RunAsync(workerhandlers.NewReconcileInvoices(svc).Handle)
		defer reconcileJob.Close()
	}

	srv := httpserver.NewServer(svc)
	srv.Serve()
	defer srv.Shutdown(context.Background())

	<-ctx.Done()
}
