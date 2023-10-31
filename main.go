package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AddUrls(r *gin.Engine) {
	r.GET("/hostnames", GetHostnamesWithActiveIPs)
}

func InitializeRouter() (*gin.Engine, string) {
	godotenv.Load()
	r := gin.Default()
	AddUrls(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Set a default port if not specified in the environment.
	}

	// fmt.Printf("Server is running on port %s\n", port)
	return r, port
}
func main() {
	r, port := InitializeRouter()
	log.Fatal(r.Run(":" + port))
}
