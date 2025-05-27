package test

import (
	"bytes"
	"database/sql"
	"desafiocdc/pkg/database/sqlite"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func openTestDB() (*sql.DB, error) {
	db, err := sqlite.OpenDB(":memory:")
	if err != nil {
		return nil, err
	}

	if err := sqlite.Migrate(db, "../db/schema.sql"); err != nil {
		return nil, err
	}

	return db, nil
}

func cleanupDB(t *testing.T, db *sql.DB) {
	t.Helper()

	_, err := db.Exec(`
		BEGIN TRANSACTION;

		DELETE FROM authors;
		DELETE FROM categories;

		COMMIT;
	`)
	if err != nil {
		// stop everything if we can't clean up the database
		panic(fmt.Errorf("failed to clean up test database: %w", err))
	}
}

func performRequest(r *chi.Mux, method, path string, body any) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		req.Body = io.NopCloser(bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
