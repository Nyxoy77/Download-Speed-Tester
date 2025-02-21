package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func testDownloadSpeed() float64 {
	url := "http://proof.ovh.net/files/10Mb.dat" // Sample file
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return 0
	}

	elapsed := time.Since(start).Seconds()
	return float64(n) / (1024 * 1024) / elapsed * 8 // Convert to Mbps
}

func speedTestHandler(c *gin.Context) {
	speed := testDownloadSpeed()
	c.JSON(http.StatusOK, gin.H{"download_speed_mbps": fmt.Sprintf("%.2f", speed)})
}

func main() {
	r := gin.Default()
	r.GET("/speedtest", speedTestHandler)
	r.Run(":8081") // Start server on port 8080
}
