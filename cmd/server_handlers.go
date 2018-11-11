package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

var uniqueIPs map[string]int

func init() {
	startTime = time.Now()
}

// Uptime contains information about a single Uptime
type Uptime struct {
	Message string `json:"message" binding:"required"`
	Uptime  string `json:"uptime" binding:"required"`
}

func newUptime() *Uptime {
	precisionLevel := time.Second
	return &Uptime{"This site has been up for the following uptime", time.Since(startTime).Truncate(precisionLevel).String()}
}

// UptimeHandler returns (in memory)
func UptimeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, newUptime())
}

// StartTime contains information about a single StartTime
type StartTime struct {
	Message   string `json:"message" binding:"required"`
	StartTime string `json:"starttime" binding:"required"`
}

func getStartTime() StartTime {
	return StartTime{"Start Time", startTime.Truncate(time.Second).String()}
}

// StartTimeHandler returns (in memory)
func StartTimeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, getStartTime())
}

// ClientIP contains information about a single ClientIP
type ClientIP struct {
	Message  string `json:"message" binding:"required"`
	ClientIP string `json:"clientip" binding:"required"`
}

func getClientIP(c *gin.Context) ClientIP {
	return ClientIP{"Current ClientIP", c.ClientIP()}
}

// ClientIPHandler returns (in memory)
func ClientIPHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, getClientIP(c))
}

// UniqueIP contains information about a single ClientIP
type UniqueIP struct {
	Message  string `json:"message" binding:"required"`
	UniqueIP string `json:"uniqueip" binding:"required"`
}

func getUniqueIP(uniqueIPs map[string]int) UniqueIP {
	var finalString string
	for k, v := range uniqueIPs {
		finalString += fmt.Sprintf("IP Address [%s] | Appearances [%d] \n", k, v)
	}
	return UniqueIP{"Unique IP Address's", finalString}
}

// IPAddressHandler returns (in memory)
func IPAddressHandler(c *gin.Context) {
	targetFiles := []string{"gin.log"}
	uniqueIPs = findUniqueIPs(targetFiles)
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, getUniqueIP(uniqueIPs))
}
