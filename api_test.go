package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"safelyyou/api"
	"safelyyou/service"
	"safelyyou/store"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestPostHeartbeatSuccess(t *testing.T) {
	store := store.NewDataStore()
	svc := service.NewDeviceService(store)
	handler := api.NewHandler(svc)

	// Create a test device
	deviceID := "test-device-123"
	svc.CreateDevice([]string{deviceID})

	// Create a test router and register the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/v1/devices/:device_id/heartbeat", handler.PostHeartbeat)

	// Create a test request
	reqBody := `{"sent_at": "2024-06-01T12:00:00Z"}`
	req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/devices/%s/heartbeat", deviceID), bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	// Verify that the heartbeat was recorded in the store
	device, _ := store.GetDevice(deviceID)
	if device.LastHeartbeat.IsZero() || device.LastHeartbeat.Format(time.RFC3339) != "2024-06-01T12:00:00Z" {
		t.Errorf("Expected LastHeartbeat to be recorded correctly, got %v", device.LastHeartbeat)
	}
}

func TestPostHeartbeatFailure(t *testing.T) {
	store := store.NewDataStore()
	svc := service.NewDeviceService(store)
	handler := api.NewHandler(svc)

	// Create a test router and register the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/v1/devices/:device_id/heartbeat", handler.PostHeartbeat)

	// Create a test request for a non-existent device
	reqBody := `{"sent_at": "2024-06-01T12:00:00Z"}`
	req, _ := http.NewRequest("POST", "/api/v1/devices/non-existent-device/heartbeat", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestPostStatsSuccess(t *testing.T) {
	store := store.NewDataStore()
	svc := service.NewDeviceService(store)
	handler := api.NewHandler(svc)

	// Create a test device
	deviceID := "test-device-456"
	svc.CreateDevice([]string{deviceID})

	// Create a test router and register the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/v1/devices/:device_id/stats", handler.PostStats)

	// Create a test request
	reqBody := `{"sent_at": "2024-06-01T12:00:00Z", "upload_time": 150}`
	req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/devices/%s/stats", deviceID), bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	// Verify that the stats were recorded in the store
	device, _ := store.GetDevice(deviceID)
	if device.UploadCount != 1 || device.UploadSum != 150 {
		t.Errorf("Expected UploadCount and UploadSum to be recorded correctly, got %d and %d", device.UploadCount, device.UploadSum)
	}
}

func TestPostStatsFailure(t *testing.T) {
	store := store.NewDataStore()
	svc := service.NewDeviceService(store)
	handler := api.NewHandler(svc)

	// Create a test router and register the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/v1/devices/:device_id/stats", handler.PostStats)

	// Create a test request for a non-existent device
	reqBody := `{"sent_at": "2024-06-01T12:00:00Z", "upload_time": 150}`
	req, _ := http.NewRequest("POST", "/api/v1/devices/non-existent-device/stats", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestGetStatsSuccess(t *testing.T) {
	store := store.NewDataStore()
	svc := service.NewDeviceService(store)
	handler := api.NewHandler(svc)

	// Create a test device and record some stats
	deviceID := "test-device-789"
	svc.CreateDevice([]string{deviceID})
	now := time.Now()
	svc.RecordHeartbeat(deviceID, now)
	svc.RecordHeartbeat(deviceID, now.Add(2*time.Minute))

	svc.RecordStats(deviceID, now, 20000000000)
	svc.RecordStats(deviceID, now.Add(1*time.Minute), 40000000000)

	// Create a test router and register the handler
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/api/v1/devices/:device_id/stats", handler.GetStats)

	// Create a test request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/devices/%s/stats", deviceID), nil)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expectedBody := `{"uptime":100,"avg_upload_time":"30s"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}
