package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ws(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer func() {
		if err := ws.Close(); err != nil {
			log.Println("close:", err)
		}
	}()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func homePage(c *gin.Context) {
	c.File("public/index.html")
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	bindAddress := os.Getenv("SERVER_URL")
	if bindAddress == "" {
		log.Fatal("SERVER_URL is not set in the environment")
	}

	r := gin.Default()
	r.GET("/ws", ws)
	r.GET("/", homePage)

	err := r.Run(bindAddress)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
