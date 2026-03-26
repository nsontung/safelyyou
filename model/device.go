package model

import "time"

type Device struct {
	ID string

	HeartbeatCount int
	FirstHeartbeat time.Time
	LastHeartbeat  time.Time

	UploadCount int
	UploadSum   int
}

func NewDevice(id string) *Device {
	return &Device{ID: id}
}

func (d *Device) RecordHeartbeat(t time.Time) {
	if d.HeartbeatCount == 0 {
		d.FirstHeartbeat = t
	}

	d.LastHeartbeat = t
	d.HeartbeatCount++
}

func (d *Device) RecordUpload(uploadTime int) {
	d.UploadCount++
	d.UploadSum += uploadTime
}

func (d *Device) GetStats() (uptime, averageUploadTime float64) {
	// Calculate uptime as the percentage of time the device has been active based on heartbeats
	if d.HeartbeatCount > 0 {
		if d.HeartbeatCount == 1 {
			uptime = 100.0 // If there's only one heartbeat, we can consider the device as fully active
		} else {
			// Be aware of division by zero if the first and last heartbeat are the same
			// *** NOTE ***
			// I think the formula should be like this
			// uptime = (number of heartbeats) / (time between first and last heartbeat in minutes + 1) * 100
			// We should add 1 minute. This way we avoid division by zero and also account for the fact that a single heartbeat should indicate some uptime.
			uptime = float64(d.HeartbeatCount) / (d.LastHeartbeat.Sub(d.FirstHeartbeat).Minutes()) * 100
		}
	}

	if d.UploadCount > 0 {
		averageUploadTime = float64(d.UploadSum) / float64(d.UploadCount)
	}

	return uptime, averageUploadTime
}
