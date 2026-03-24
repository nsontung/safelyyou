package main

import (
	"fmt"
	"os"
	"safelyyou/api"
	"safelyyou/service"
	"safelyyou/store"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting the server at :6733")
	// Start the Gin-Gonic server
	router := gin.Default()
	store := store.NewDataStore()
	svc := service.NewDeviceService(store)
	handler := api.NewHandler(svc)
	group := router.Group("/api/v1")
	// Uncomment to add this middleware for logging, debugging purposes
	// group.Use(middleware.RequestLogger())
	api.RegisterRoutes(group, handler)

	// Load device IDs from CSV and create devices in the store
	deviceIDs, err := api.LoadDeviceIDsFromCSV("devices.csv")
	if err != nil {
		fmt.Println("Error loading device IDs:", err)
		os.Exit(1)
	}
	svc.CreateDevice(deviceIDs)

	router.Run(":6733")
}
