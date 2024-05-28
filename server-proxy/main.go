package main

import (
	"fmt"

	"net/url"

	"github.com/gin-gonic/gin"
)

func getProxyWebSocket(c *gin.Context) {
	urlTarget := url.URL{
		Scheme: "ws",
		Host:   "localhost:8448",
		Path:   "/ws",
	}

	proxy := NewProxy(&urlTarget)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func homePage(c *gin.Context) {
	c.File("public/index.html")
}

func main() {
	fmt.Println("Proxy Server!")

	bindAddress := "localhost:8488"
	r := gin.Default()

	r.GET("/ws", getProxyWebSocket)
	r.GET("/", homePage)
	r.Run(bindAddress)
}
