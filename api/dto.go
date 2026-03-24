package api

import "time"

type HeartbeatRequest struct {
	SentAt time.Time `json:"sent_at"`
}

type StatsRequest struct {
	SentAt     time.Time `json:"sent_at"`
	UploadTime int       `json:"upload_time"`
}

type StatsResponse struct {
	Uptime        float64 `json:"uptime"`
	AvgUploadTime string  `json:"avg_upload_time"`
}
