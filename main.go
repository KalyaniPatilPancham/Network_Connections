package main

import (
	"demo/connection"
	"demo/fetcher"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	fetch := fetcher.NewFetcher()

	r.GET("/connections", func(c *gin.Context) {
		url := "https://my-json-server.typicode.com/marcuzh/router_location_test_api/db" // Replace with your actual URL
		data, err := fetch.FetchJSONData(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if data == nil || (len(data.Routers) == 0 && len(data.Locations) == 0) {
			c.JSON(http.StatusOK, gin.H{"message": "No data available"})
			return
		}

		connections := connection.FindConnections(data)
		c.JSON(http.StatusOK, gin.H{"connections": connections})
	})

	r.Run(":8080")
}
