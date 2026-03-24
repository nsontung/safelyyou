package api

import (
	"safelyyou/service"
	"safelyyou/syerror"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.DeviceService
}

func NewHandler(svc *service.DeviceService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) PostHeartbeat(c *gin.Context) {
	var req HeartbeatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	deviceID := c.Param("device_id")
	if err := h.svc.RecordHeartbeat(deviceID, req.SentAt); err != nil {
		switch err {
		case syerror.ErrDeviceNotFound:
			c.JSON(404, gin.H{"msg": err.Error()})
		default:
			c.JSON(500, gin.H{"msg": err.Error()})
		}
		return
	}
	c.Status(204)
}

func (h *Handler) PostStats(c *gin.Context) {
	var req StatsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	deviceID := c.Param("device_id")
	if err := h.svc.RecordStats(deviceID, req.SentAt, req.UploadTime); err != nil {
		switch err {
		case syerror.ErrDeviceNotFound:
			c.JSON(404, gin.H{"msg": err.Error()})
		default:
			c.JSON(500, gin.H{"msg": err.Error()})
		}
		return
	}
	c.Status(204)
}

func (h *Handler) GetStats(c *gin.Context) {
	deviceID := c.Param("device_id")
	uptime, avgUploadTime, err := h.svc.GetStats(deviceID)
	if err != nil {
		switch err {
		case syerror.ErrDeviceNotFound:
			c.JSON(404, gin.H{"msg": err.Error()})
		default:
			c.JSON(500, gin.H{"msg": err.Error()})
		}
		return
	}

	avgUploadTimeDuration := time.Duration(avgUploadTime) * time.Nanosecond

	resp := StatsResponse{
		Uptime:        uptime,
		AvgUploadTime: avgUploadTimeDuration.String(),
	}
	c.JSON(200, resp)
}
