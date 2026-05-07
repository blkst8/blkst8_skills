package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/example/invoice-service/internal/handler"
	"github.com/example/invoice-service/internal/repository"
	"github.com/example/invoice-service/internal/service"
	"github.com/example/invoice-service/internal/worker"
	"github.com/example/invoice-service/server"
)

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN environment variable is required")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Wire up dependencies using private DI pattern
	invoiceRepo := repository.NewInvoiceRepository(db)
	invoiceService := service.NewInvoiceService(invoiceRepo)
	invoiceHandler := handler.NewInvoiceHandler(invoiceService)
	reconcileWorker := worker.NewReconcileWorker(invoiceService)

	// Start HTTP server
	srv := server.New(invoiceHandler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start background worker
	go reconcileWorker.Start(ctx)

	// Start server in background
	go func() {
		addr := os.Getenv("SERVER_ADDR")
		if addr == "" {
			addr = ":8080"
		}
		log.Printf("starting HTTP server on %s", addr)
		if err := srv.Start(addr); err != nil {
			log.Printf("server stopped: %v", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down...")
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

	log.Println("shutdown complete")
}
