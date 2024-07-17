package main

import (
	"demo/connection"
	"demo/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockFetcher struct{}

func (m *mockFetcher) FetchJSONData(url string) (*models.Data, error) {
	return &models.Data{
		Routers: []models.Router{
			{ID: 1, Name: "RouterA", LocationID: 1, RouterLinks: []int{2}},
			{ID: 2, Name: "RouterB", LocationID: 2, RouterLinks: []int{1}},
		},
		Locations: []models.Location{
			{ID: 1, Name: "LocationA"},
			{ID: 2, Name: "LocationB"},
		},
	}, nil
}

func TestConnectionsRoute(t *testing.T) {
	fetcher := &mockFetcher{}

	r := gin.Default()
	r.GET("/connections", func(c *gin.Context) {
		data, err := fetcher.FetchJSONData("https://api.example.com/data")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		connections := connection.FindConnections(data)
		c.JSON(http.StatusOK, gin.H{"connections": connections})
	})

	req, _ := http.NewRequest("GET", "/connections", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"connections":["LocationA <-> LocationB"]}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
