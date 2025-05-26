package main

import (
	"desafiocdc/internal/author"
	"desafiocdc/pkg/database/sqlite"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Default().Println("opening db")
	db, err := sqlite.OpenDB("file:tmp/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Default().Println("running migrations")
	if err := sqlite.Migrate(db, "db/schema.sql"); err != nil {
		log.Fatal(err)
	}

	log.Default().Println("configuring router")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	author.Setup(r, db)

	log.Default().Println("staring server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
