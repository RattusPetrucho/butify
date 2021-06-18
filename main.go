package main

import (
	"butify/pkg/adapters"
	"butify/pkg/config"
	"butify/pkg/endpoints"
	"butify/pkg/repositories"
	"butify/pkg/services"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.New()
	db, err := sqlx.Open(cfg.DbDialect, cfg.DbDsn)
	if err != nil {
		log.Println("open db:", err)
		return
	}
	defer db.Close()

	adapter := adapters.New(db)
	repo := repositories.New(adapter)
	service := services.New(repo)
	endpoints := endpoints.New(service)

	log.Fatal(http.ListenAndServe(":80", endpoints))
}
