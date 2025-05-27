package test

import (
	"desafiocdc/internal/author"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuthorHandler(t *testing.T) {
	db, err := openTestDB()
	assert.Nil(t, err)
	defer db.Close()

	r := chi.NewRouter()
	author.Setup(r, db)

	t.Run("should return error when name is empty", func(t *testing.T) {
		cleanupDB(t, db)

		req := author.CreateAuthorRequest{
			Name:        "",
			Email:       "john@doe.com",
			Description: "A test author",
		}

		w := performRequest(r, "POST", "/authors", req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "name is required")
	})

	t.Run("should return error when email is empty", func(t *testing.T) {
		cleanupDB(t, db)

		req := author.CreateAuthorRequest{
			Name:        "John Doe",
			Email:       "",
			Description: "A test author",
		}

		w := performRequest(r, "POST", "/authors", req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "email is required")
	})

	t.Run("should return error when description is empty", func(t *testing.T) {
		cleanupDB(t, db)

		req := author.CreateAuthorRequest{
			Name:        "John Doe",
			Email:       "john@doe.com",
			Description: "",
		}

		w := performRequest(r, "POST", "/authors", req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "description is required")
	})

	t.Run("should return error when description is higher than 400 characters", func(t *testing.T) {
		cleanupDB(t, db)

		req := author.CreateAuthorRequest{
			Name:        "John Doe",
			Email:       "john@doe.com",
			Description: strings.Repeat("a", 401),
		}

		w := performRequest(r, "POST", "/authors", req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "description must be less than 400 characters")
	})

	t.Run("should return error when email already exists", func(t *testing.T) {
		cleanupDB(t, db)

		req := author.CreateAuthorRequest{
			Name:        "John Doe",
			Email:       "john@doe.com",
			Description: "A test author",
		}

		w := performRequest(r, "POST", "/authors", req)
		assert.Equal(t, http.StatusOK, w.Code)

		w = performRequest(r, "POST", "/authors", req)
		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Contains(t, w.Body.String(), "there is already an author registered with the email "+req.Email)
	})

	t.Run("should create author successfully", func(t *testing.T) {
		cleanupDB(t, db)

		req := author.CreateAuthorRequest{
			Name:        "Jane Doe",
			Email:       "john@doe.com",
			Description: strings.Repeat("a", 400),
		}

		w := performRequest(r, "POST", "/authors", req)
		assert.Equal(t, http.StatusOK, w.Code)

		var resp author.CreateAuthorResponse
		err := json.NewDecoder(w.Body).Decode(&resp)
		assert.Nil(t, err)

		assert.Equal(t, req.Name, resp.Name)
		assert.Equal(t, req.Email, resp.Email)
		assert.Equal(t, req.Description, resp.Description)
		assert.NotEmpty(t, resp.ID)
		assert.NotEmpty(t, resp.CreatedAt)
	})
}
