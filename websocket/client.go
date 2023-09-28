package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	IsPanel bool
	Conn    *websocket.Conn
	Pool    *Pool
}

type Message struct {
	IsFromPanel bool
	Type        int
	Body        string
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{
			IsFromPanel: c.IsPanel,
			Type:        messageType,
			Body:        string(p),
		}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
