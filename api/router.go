package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	r.POST("/devices/:device_id/heartbeat", h.PostHeartbeat)
	r.POST("/devices/:device_id/stats", h.PostStats)
	r.GET("/devices/:device_id/stats", h.GetStats)
}
