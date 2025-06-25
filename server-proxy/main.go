package main

import (
	"log"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getProxyWebSocket(c *gin.Context, host string) {
	urlTarget := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   "/ws",
	}

	proxy := NewProxy(&urlTarget)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func homePage(c *gin.Context) {
	c.File("public/index.html")
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file", err.Error())
	}

	proxy := os.Getenv("PROXY")
	if proxy == "" {
		log.Fatal("PROXY environment variable is not set")
	}

	bindAddress := os.Getenv("SERVER_URL")
	if bindAddress == "" {
		log.Fatal("SERVER_URL environment variable is not set")
	}

	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		getProxyWebSocket(c, proxy)
	})
	r.GET("/", homePage)
	if err := r.Run(bindAddress); err != nil {
		panic(err)
	}
}
