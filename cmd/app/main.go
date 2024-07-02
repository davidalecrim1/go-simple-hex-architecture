package main

import (
	"database/sql"
	"go_simple_hex_architecture/internal"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "tax.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	taxRepo := internal.NewTaxSQLiteRepository(db)
	taxSvc := internal.NewTaxService(taxRepo)
	taxWeb := internal.NewTaxWeb(*taxSvc)

	taxRepo.InitDatabaseScript()

	c := chi.NewRouter()
	c.Post("/tax", taxWeb.TaxHandler)
	http.ListenAndServe(":8080", c)
}
