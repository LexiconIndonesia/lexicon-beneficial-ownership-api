package main

import (
	"context"
	bo "lexicon/bo-api/beneficiary_ownership"
	bo_v1 "lexicon/bo-api/beneficiary_ownership/v1"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// INITIATE CONFIGURATION
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}
	cfg := defaultConfig()
	cfg.loadFromEnv()

	log.Debug().Any("config", cfg).Msg("config loaded")

	ctx := context.Background()
	// INITIATE DATABASES

	// PGSQL
	pgsqlClient, err := pgxpool.New(ctx, cfg.PgSql.ConnStr())

	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to PGSQL Database")
	}
	defer pgsqlClient.Close()

	bo.SetDatabase(pgsqlClient)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// // Set a timeout value on the request context (ctx), that will signal
	// // through ctx.Done() that the request has timed out and further
	// // processing should be stopped.
	// r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/beneficiary-ownership", bo_v1.Router())
	})

	log.Info().Msg("Starting up server...")

	if err := http.ListenAndServe(cfg.Listen.Addr(), r); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
		return
	}

	log.Info().Msg("Server Stopped")
}
