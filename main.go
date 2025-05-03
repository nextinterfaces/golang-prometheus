package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Define Prometheus metrics
var requestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP Requests",
	},
	[]string{"path", "method"},
)

func init() {
	// Register custom metrics
	prometheus.MustRegister(requestCount)
}

func main() {
	router := gin.Default()

	// Middleware to collect metrics
	router.Use(func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method
		c.Next()
		requestCount.WithLabelValues(path, method).Inc()
	})

	// Basic endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Run(":8080")
}
