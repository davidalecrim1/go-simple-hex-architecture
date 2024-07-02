package internal

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type TaxSQLiteRepository struct {
	DB *sql.DB
}

func NewTaxSQLiteRepository(db *sql.DB) *TaxSQLiteRepository {
	return &TaxSQLiteRepository{DB: db}
}

func (r *TaxSQLiteRepository) Save(id string, tax float64) error {
	_, err := r.DB.Exec("INSERT INTO tax (id, tax) VALUES (?, ?)", id, tax)
	return err
}

func (r *TaxSQLiteRepository) InitDatabaseScript() {
	sqlFile, err := os.ReadFile("./sql/schema.sql")

	if err != nil {
		log.Fatal(err)
	}
	_, err = r.DB.Exec(string(sqlFile))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("SQL script executed successfully!")
}
