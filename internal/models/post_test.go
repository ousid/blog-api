package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/coderflexx/blog-api/internal/models"
)

func TestPostIsPublished(t *testing.T) {
	now := time.Now()

	// Published post
	published := models.Post{PublishedAt: &now}
	assert.True(t, published.IsPublished())

	// Draft post
	draft := models.Post{PublishedAt: nil}
	assert.False(t, draft.IsPublished())
}
