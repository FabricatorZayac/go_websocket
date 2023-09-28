package main

import (
	"fmt"
	"net/http"

    "gosocket/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request, isPanel bool) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		IsPanel: isPanel,
		Conn:    conn,
		Pool:    pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r, false)
	})
	http.HandleFunc("/panel", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r, true)
	})
}

func main() {
	fmt.Println("Websocket server")
	setupRoutes()

	http.ListenAndServe(":4444", nil)
}
