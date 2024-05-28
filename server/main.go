package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/gorilla/websocket"
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
	defer ws.Close()
	for {
		//read data from ws
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		//write ws data
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
	fmt.Println("Websocket Server!")

	bindAddress := "localhost:8448"
	r := gin.Default()
	r.GET("/ws", ws)
	r.GET("/", homePage)
	r.Run(bindAddress)
}
