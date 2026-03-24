package model

import "time"

type Device struct {
	ID string

	heartbeatCount int
	firstHeartbeat time.Time
	lastHeartbeat  time.Time

	uploadCount int
	uploadSum   int
}

func NewDevice(id string) *Device {
	return &Device{ID: id}
}

func (d *Device) RecordHeartbeat(t time.Time) {
	if d.heartbeatCount == 0 {
		d.firstHeartbeat = t
	}

	d.lastHeartbeat = t
	d.heartbeatCount++
}

func (d *Device) RecordUpload(uploadTime int) {
	d.uploadCount++
	d.uploadSum += uploadTime
}

func (d *Device) GetStats() (uptime, averageUploadTime float64) {
	// Calculate uptime as the percentage of time the device has been active based on heartbeats
	if d.heartbeatCount > 0 {
		if d.heartbeatCount == 1 {
			uptime = 100.0 // If there's only one heartbeat, we can consider the device as fully active
		} else {
			// Be aware of division by zero if the first and last heartbeat are the same
			// *** NOTE ***
			// I think the formula should be like this
			// uptime = (number of heartbeats) / (time between first and last heartbeat in minutes + 1) * 100
			// We should add 1 minute. This way we avoid division by zero and also account for the fact that a single heartbeat should indicate some uptime.
			uptime = float64(d.heartbeatCount) / (d.lastHeartbeat.Sub(d.firstHeartbeat).Minutes()) * 100
		}
	}

	if d.uploadCount > 0 {
		averageUploadTime = float64(d.uploadSum) / float64(d.uploadCount)
	}

	return uptime, averageUploadTime
}
