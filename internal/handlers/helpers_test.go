package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coderflexx/blog-api/internal/database"
	"github.com/coderflexx/blog-api/internal/models"
	"github.com/coderflexx/blog-api/internal/router"
)

// bootstraps the full app for each test
func setupTestApp(t *testing.T) http.Handler {
	t.Helper()
	database.SetupTestDB()
	return router.Setup()
}

func makeRequest(t *testing.T, app http.Handler, method, url string, body any) *httptest.ResponseRecorder {
	t.Helper()

	var req *http.Request

	if body != nil {
		b, _ := json.Marshal(body)
		req = httptest.NewRequest(method, url, bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req = httptest.NewRequest(method, url, nil)
	}

	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	return rr
}

// seed a category for tests that need one
func seedCategory(t *testing.T) models.Category {
	t.Helper()
	category := models.Category{Name: "Laravel", Slug: "laravel"}
	database.DB.Create(&category)
	return category
}

var postCounter int

func seedPost(t *testing.T, categoryID uint) models.Post {
	t.Helper()
	postCounter++
	post := models.Post{
		Title:      fmt.Sprintf("Test Post %d", postCounter),
		Slug:       fmt.Sprintf("test-post-%d", postCounter),
		Content:    "This is test content for the post.",
		Excerpt:    "Test excerpt",
		CategoryID: categoryID,
	}
	database.DB.Create(&post)
	return post
}
