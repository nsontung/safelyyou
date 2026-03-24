package service

import (
	"crypto/rand"
	"safelyyou/store"
	"testing"
	"time"
)

func TestStatsCalculation(t *testing.T) {

	tests := []struct {
		name                   string
		heartbeatsAddedMinutes []int
		uploadTimes            []int
		expectedUptime         float64
		expectedAvgUploadTime  float64
	}{
		{
			name:                   "Multiple heartbeats and stats",
			heartbeatsAddedMinutes: []int{-2, 0, 1, 2, 3, 4},
			uploadTimes:            []int{100, 110, 120, 130, 140},
			expectedUptime:         100.0,
			expectedAvgUploadTime:  120.0,
		},
		{
			name:                   "One heartbeat and one stat",
			heartbeatsAddedMinutes: []int{-2, 0},
			uploadTimes:            []int{100},
			expectedUptime:         100.0,
			expectedAvgUploadTime:  100.0,
		},
		{
			name:                   "No heartbeats, no stats",
			heartbeatsAddedMinutes: []int{},
			uploadTimes:            []int{},
			expectedUptime:         0.0,
			expectedAvgUploadTime:  0.0,
		},
		{
			name:                   "Multiple heartbeats, no stats",
			heartbeatsAddedMinutes: []int{-2, 0, 1, 2, 3, 4},
			uploadTimes:            []int{},
			expectedUptime:         100.0,
			expectedAvgUploadTime:  0.0,
		},
		{
			name:                   "No heartbeats, multiple stats",
			heartbeatsAddedMinutes: []int{},
			uploadTimes:            []int{100, 110, 120, 130, 140},
			expectedUptime:         0.0,
			expectedAvgUploadTime:  120.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := store.NewDataStore()
			svc := NewDeviceService(store)

			// Create a device
			deviceID := rand.Text()
			svc.CreateDevice([]string{deviceID})

			now := time.Now()
			for i := range tt.heartbeatsAddedMinutes {
				svc.RecordHeartbeat(deviceID, now.Add(time.Duration(tt.heartbeatsAddedMinutes[i])*time.Minute))
			}

			for i := range tt.uploadTimes {
				svc.RecordStats(deviceID, now.Add(time.Duration(i)*time.Minute), tt.uploadTimes[i])
			}

			// Get stats and verify calculations
			uptime, avgUploadTime, err := svc.GetStats(deviceID)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if uptime != tt.expectedUptime {
				t.Errorf("Expected uptime %v, got %v", tt.expectedUptime, uptime)
			}
			if avgUploadTime != tt.expectedAvgUploadTime {
				t.Errorf("Expected average upload time %v, got %v", tt.expectedAvgUploadTime, avgUploadTime)
			}
		})
	}
}

func TestOneStatsOnly(t *testing.T) {
	store := store.NewDataStore()
	svc := NewDeviceService(store)

	// Create a device
	deviceID := "test-device-1"
	svc.CreateDevice([]string{deviceID})
	now := time.Now()
	svc.RecordHeartbeat(deviceID, now)
	svc.RecordStats(deviceID, now, 100)

	// Get stats and verify calculations
	uptime, avgUploadTime, err := svc.GetStats(deviceID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if uptime != 100.0 {
		t.Errorf("Expected uptime %v, got %v", 100.0, uptime)
	}

	if avgUploadTime != 100 {
		t.Errorf("Expected average upload time 100, got %v", avgUploadTime)
	}

}
