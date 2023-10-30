package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHostnamesWithActiveIPs(c *gin.Context) {
	//Get the value of X from envrionment variables
	x := os.Getenv("X")
	if x == "" {
		x = "1" // Set the default value to 1.
	}

	xVal, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(400, gin.H{
			"message": "Invalid value of x",
		})
		return
	}

	// Create a map to count the active IP addresses for each hostname.
	countMap := make(map[string]int)
	for _, ipConfig := range IpConfigData {
		if ipConfig.Active {
			countMap[ipConfig.Hostname]++
		}
	}

	// Filter hostnames that have active IP addresses less than or equal to the threshold.
	var resultHostnames []string
	for hostname, count := range countMap {
		if count <= xVal {
			resultHostnames = append(resultHostnames, hostname)
		}
	}

	c.JSON(200, gin.H{
		"hostnames": resultHostnames,
	})
}
