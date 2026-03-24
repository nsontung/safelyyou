package service

import (
	"safelyyou/model"
	"safelyyou/store"
	"safelyyou/syerror"
	"time"
)

type DeviceService struct {
	store *store.DataStore
}

func NewDeviceService(store *store.DataStore) *DeviceService {
	return &DeviceService{store: store}
}

func (ds *DeviceService) CreateDevice(deviceIDs []string) []*model.Device {
	var devices []*model.Device
	for _, deviceID := range deviceIDs {
		device := ds.store.CreateDevice(deviceID)
		devices = append(devices, device)
	}
	return devices
}

func (ds *DeviceService) RecordHeartbeat(deviceID string, sentAt time.Time) error {
	if exists := ds.store.UpdateHeartbeat(deviceID, sentAt); !exists {
		return syerror.ErrDeviceNotFound
	}
	return nil
}

func (ds *DeviceService) RecordStats(deviceID string, sentAt time.Time, uploadTime int) error {
	if exists := ds.store.UpdateStats(deviceID, uploadTime); !exists {
		return syerror.ErrDeviceNotFound
	}
	return nil
}

func (ds *DeviceService) GetStats(deviceID string) (uptime, averageUploadTime float64, err error) {
	uptime, averageUploadTime, exists := ds.store.GetStats(deviceID)
	if !exists {
		return 0, 0, syerror.ErrDeviceNotFound
	}
	return uptime, averageUploadTime, nil
}
