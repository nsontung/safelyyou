package store

import (
	"safelyyou/model"
	"sync"
	"time"
)

type DataStore struct {
	// mu protects access to the data map, avoid race conditions when multiple goroutines access it concurrently
	// RLock for read operations, Lock for write operations
	mu   sync.RWMutex
	data map[string]*model.Device
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: make(map[string]*model.Device),
	}
}

func (ds *DataStore) CreateDevice(deviceID string) *model.Device {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	// If the device already exists, overwrite it with a new instance
	newDevice := model.NewDevice(deviceID)
	ds.data[deviceID] = newDevice
	return newDevice
}

func (ds *DataStore) GetDevice(deviceID string) (*model.Device, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	device, exists := ds.data[deviceID]
	return device, exists
}

func (ds *DataStore) UpdateHeartbeat(deviceID string, sentAt time.Time) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if device, ok := ds.data[deviceID]; ok {
		device.RecordHeartbeat(sentAt)
		return true
	}
	return false
}

func (ds *DataStore) UpdateStats(deviceID string, uploadTime int) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	if device, exists := ds.data[deviceID]; exists {
		device.RecordUpload(uploadTime)
		return true
	}
	return false
}

func (ds *DataStore) GetStats(deviceID string) (uptime, averageUploadTime float64, exists bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	device, ok := ds.data[deviceID]
	if !ok {
		return 0, 0, false
	}
	uptime, averageUploadTime = device.GetStats()
	return uptime, averageUploadTime, true
}
