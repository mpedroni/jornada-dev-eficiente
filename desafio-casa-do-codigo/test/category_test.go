package test

import (
	"desafiocdc/internal/category"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategoryHandler(t *testing.T) {
	db, err := openTestDB()
	assert.Nil(t, err)
	defer db.Close()
	defer cleanupDB(t, db)

	r := chi.NewRouter()
	category.Setup(r, db)

	t.Run("should return bad request if body is malformed", func(t *testing.T) {
		w := performRequest(r, "POST", "/categories", `{"name": "Fantasy"},`)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return bad request if name is empty", func(t *testing.T) {
		req := category.CreateCategoryRequest{
			Name: "",
		}
		w := performRequest(r, "POST", "/categories", req)
		assert.Equal(t, http.StatusBadRequest, w.Code)

		assert.Contains(t, w.Body.String(), "name is required")
	})

	t.Run("should create a category successfully", func(t *testing.T) {
		cleanupDB(t, db)

		req := category.CreateCategoryRequest{
			Name: "Fantasy",
		}
		w := performRequest(r, "POST", "/categories", req)
		assert.Equal(t, http.StatusOK, w.Code)

		var resp category.CreateCategoryResponse

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.NotEmpty(t, resp.ID)
		assert.Equal(t, "Fantasy", resp.Name)
	})

	t.Run("should return conflict if already exists a category with the given name", func(t *testing.T) {
		cleanupDB(t, db)

		req := category.CreateCategoryRequest{
			Name: "Fantasy",
		}
		w := performRequest(r, "POST", "/categories", req)
		assert.Equal(t, http.StatusOK, w.Code)

		w = performRequest(r, "POST", "/categories", req)
		assert.Equal(t, http.StatusConflict, w.Code)

		assert.Contains(t, w.Body.String(), "there already exists a category with name Fantasy")
	})

}
