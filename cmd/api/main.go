package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fkihai/payflow/internal/delivery/httpx"
	"github.com/fkihai/payflow/internal/delivery/httpx/handler"
	"github.com/fkihai/payflow/internal/infrastructure/config"
	"github.com/fkihai/payflow/internal/infrastructure/db/postgres"
	"github.com/fkihai/payflow/internal/usecase/payment"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	conn := postgres.NewPostgresConnection(&cfg.Database)

	db, err := conn.Connect()
	if err != nil {
		fmt.Printf("cannot connect db, %v\n", err)
		return
	}

	defer db.Close()

	r := postgres.NewPostgresTransactionRepositoy(db)
	u := payment.NewPaymentUsecase(r)
	h := handler.NewPaymentHandler(u, cfg.Peyment)
	router := httpx.Router(h)

	svr := &http.Server{
		Addr:        "127.0.0.1:2001",
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	go func() {
		fmt.Printf("server starting, Addr: %s\n", "127.0.0.1:2001")
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server failed", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = svr.Shutdown(ctx)
	fmt.Println("server stopped")
}
