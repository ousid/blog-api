package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gosimple/slug"

	"github.com/coderflexx/blog-api/internal/database"
	"github.com/coderflexx/blog-api/internal/helpers"
	"github.com/coderflexx/blog-api/internal/models"
)

type PostRequest struct {
	Title      string `json:"title" validate:"required,min=3,max=255"`
	Content    string `json:"content" validate:"required,min=10"`
	Excerpt    string `json:"excerpt" validate:"max=500"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post

	result := database.DB.
		Preload("Category").
		Order("created_at DESC").
		Find(&posts)

	if result.Error != nil {
		helpers.ServerError(w, "Failed to fetch posts")
		return
	}

	helpers.Success(w, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	var post models.Post

	result := database.DB.
		Preload("Category").
		Where("slug = ?", slug).
		First(&post)

	if result.Error != nil {
		helpers.NotFound(w, "Post Not Found")
		return
	}

	helpers.Success(w, post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var req PostRequest

	// Decode + validate in one step
	errs, err := helpers.DecodeAndValidate(r, &req)
	if err != nil {
		helpers.BadRequest(w, err.Error())
		return
	}

	if errs != nil {
		helpers.ValidationFailed(w, errs)
		return
	}

	// Check category exists
	var category models.Category
	if result := database.DB.First(&category, req.CategoryID); result.Error != nil {
		helpers.NotFound(w, "Category not found")
		return
	}

	titleSlug := slug.Make(req.Title)
	if slugExists(titleSlug, 0) {
		helpers.UnprocessableEntity(w, "A post with this title already exists")
		return
	}

	post := models.Post{
		Title:      req.Title,
		Slug:       titleSlug,
		Content:    req.Content,
		Excerpt:    req.Excerpt,
		CategoryID: req.CategoryID,
	}

	result := database.DB.Create(&post)
	if result.Error != nil {
		helpers.ServerError(w, "Failed to create post")
		return
	}

	// load the category relation before returning
	database.DB.Preload("Category").First(&post, post.ID)

	helpers.Created(w, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// find post
	post, ok := findPostByID(w, id)

	if !ok {
		return
	}

	var req PostRequest
	errs, err := helpers.DecodeAndValidate(r, &req)
	if err != nil {
		helpers.BadRequest(w, err.Error())
		return
	}

	if errs != nil {
		helpers.ValidationFailed(w, errs)
		return
	}

	// check category exists
	var category models.Category
	if database.DB.First(&category, req.CategoryID).Error != nil {
		helpers.NotFound(w, "Category not found")
		return
	}

	// check slug uniquness
	newSlug := slug.Make(req.Title)
	if slugExists(newSlug, post.ID) {
		helpers.UnprocessableEntity(w, "A post with this title already exists")
		return
	}

	// udpate fields
	post.Title = req.Title
	post.Slug = newSlug
	post.Content = req.Content
	post.Excerpt = req.Excerpt
	post.CategoryID = req.CategoryID

	if database.DB.Save(&post).Error != nil {
		helpers.ServerError(w, "Failed to update post")
		return
	}

	database.DB.Preload("Category").First(&post, post.ID)

	helpers.Success(w, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post, ok := findPostByID(w, id)
	if !ok {
		return
	}

	// Soft Delete
	if database.DB.Delete(&post).Error != nil {
		helpers.ServerError(w, "Failed to delete post")
		return
	}

	helpers.NotContent(w)
}

func findPostByID(w http.ResponseWriter, id string) (*models.Post, bool) {
	var post models.Post

	if database.DB.First(&post, id).Error != nil {
		helpers.NotFound(w, "Post Not Found")
		return nil, false
	}

	return &post, true
}

func slugExists(newSlug string, excludeID uint) bool {
	var existing models.Post
	query := database.DB.Where("slug = ?", newSlug)

	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	return query.First(&existing).Error == nil
}
