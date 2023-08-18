package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate"
	"github.com/puoklam/shipper-core/db"
	"github.com/puoklam/shipper-core/http/api"
)

func init() {
	if os.Getenv("ENV") == "dev" {
		os.Setenv("PORT", "3000")
		os.Setenv("DB_URL", "postgres://admin:password@host.docker.internal/shipper-core?sslmode=disable")
		os.Setenv("MIGRATION_SRC_URL", "./db/migrations")
	}

	err := db.Migrate()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalln(err)
	}

	if err = db.Init(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("exit")
		os.Exit(0)
	}()

	r := chi.NewRouter()
	SetupMiddlewares(r)

	handlers := api.New()
	for _, h := range handlers {
		h.SetupRoutes(r)
	}

	srv := NewServer(r)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
