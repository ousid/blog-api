package handlers_test

import(
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coderflexx/blog-api/internal/database"
	"github.com/coderflexx/blog-api/internal/models"
	"github.com/coderflexx/blog-api/internal/router"
	"github.com/stretchr/testify/assert"
)

// bootstrap the full app for each test
func setupTestApp(t *testing.T) http.Handler {
	t.Helper()
	database.SetupTestDB()
	return router.Setup()
}

