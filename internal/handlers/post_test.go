package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coderflexx/blog-api/internal/database"
)

func TestListPosts(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	category := seedCategory(t)
	seedPost(t, category.ID)
	seedPost(t, category.ID)

	rr := makeRequest(t, app, http.MethodGet, "/api/posts", nil)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]any
	_ = json.Unmarshal(rr.Body.Bytes(), &response)

	data := response["data"].([]any)
	assert.Len(t, data, 2)
}

func TestGetPost(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	category := seedCategory(t)
	post := seedPost(t, category.ID)

	rr := makeRequest(t, app, http.MethodGet, "/api/posts/"+post.Slug, nil)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]any
	_ = json.Unmarshal(rr.Body.Bytes(), &response)

	data := response["data"].(map[string]any)
	assert.Equal(t, post.Title, data["title"])
	assert.Equal(t, post.Slug, data["slug"])
}

func TestGetPost_NotFound(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	rr := makeRequest(t, app, http.MethodGet, "/api/posts/non-existent", nil)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestCreatePost(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	category := seedCategory(t)

	body := map[string]any{
		"title":       "my new post",
		"content":     "this is the full content of my new post.",
		"excerpt":     "short summary.",
		"category_id": category.ID,
	}

	rr := makeRequest(t, app, http.MethodPost, "/api/posts", body)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response map[string]any
	_ = json.Unmarshal(rr.Body.Bytes(), &response)

	data := response["data"].(map[string]any)
	assert.Equal(t, "my new post", data["title"])
	assert.Equal(t, "my-new-post", data["slug"])
}

func TestCreatePost_ValidationFails(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	body := map[string]any{
		"title": "Hi",
	}

	rr := makeRequest(t, app, http.MethodPost, "/api/posts", body)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

	var response map[string]any
	_ = json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, "Validation failed", response["message"])
	assert.NotEmpty(t, response["errors"])
}

func TestCreatePost_DuplicatingSlug(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	category := seedCategory(t)
	post := seedPost(t, category.ID)

	body := map[string]any{
		"title":       post.Title,
		"content":     "Different content but same title slug.",
		"category_id": category.ID,
	}

	rr := makeRequest(t, app, http.MethodPost, "/api/posts", body)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestUpdatePost(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	category := seedCategory(t)
	post := seedPost(t, category.ID)

	body := map[string]any{
		"title":       "Updated Post Title",
		"content":     "Updated content that is long enough to pass validation.",
		"excerpt":     "Updated excerpt",
		"category_id": category.ID,
	}

	rr := makeRequest(t, app, http.MethodPut, fmt.Sprintf("/api/posts/%d", post.ID), body)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]any
	_ = json.Unmarshal(rr.Body.Bytes(), &response)

	data := response["data"].(map[string]any)
	assert.Equal(t, "Updated Post Title", data["title"])
	assert.Equal(t, "updated-post-title", data["slug"])
}

func TestUpdatePost_NotFound(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	body := map[string]any{
		"title":       "Updated Post Title",
		"content":     "Updated content that is long enough.",
		"category_id": 1,
	}

	rr := makeRequest(t, app, http.MethodPut, "/api/posts/999", body)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestDeletePost(t *testing.T) {
	app := setupTestApp(t)
	defer database.CleanupTestDB()

	category := seedCategory(t)
	post := seedPost(t, category.ID)

	rr := makeRequest(t, app, http.MethodDelete, fmt.Sprintf("/api/posts/%d", post.ID), nil)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
