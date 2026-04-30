package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"tasm-backend/database"
)

func TestGetConsumablesReturnsServiceUnavailableWithoutDB(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = nil

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/consumables", nil)

	GetConsumables(context)

	if recorder.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status %d, got %d", http.StatusServiceUnavailable, recorder.Code)
	}
}

func TestCreateConsumableRejectsInvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/api/consumables", strings.NewReader("{"))
	context.Request.Header.Set("Content-Type", "application/json")

	CreateConsumable(context)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestCreateConsumableRequiresDatabaseForValidPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.DB = nil

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/api/consumables", strings.NewReader(`{
		"name":"Printer Paper",
		"category":"Office Supplies",
		"location":"Central Stockroom",
		"currentStock":10,
		"reorderLevel":5
	}`))
	context.Request.Header.Set("Content-Type", "application/json")

	CreateConsumable(context)

	if recorder.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status %d, got %d", http.StatusServiceUnavailable, recorder.Code)
	}
}
