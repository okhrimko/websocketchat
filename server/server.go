package chat

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Conn struct {
	WS   *websocket.Conn
	Send chan string
}

func (c *Conn) SendToHub() {
	defer c.WS.Close()
	for {
		_, msg, err := c.WS.ReadMessage()
		if err != nil {
			return
		}
		DefaultHub.Echo <- string(msg)
	}
}

func (c *Conn) ReceiveFromHub() {
	defer c.WS.Close()
	for {
		c.Write(<-c.Send)
	}
}

func (c *Conn) Write(msg string) error {
	return c.WS.WriteMessage(websocket.TextMessage, []byte(msg))
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	conn := &Conn{
		Send: make(chan string),
		WS:   ws,
	}
	DefaultHub.Join <- conn

	go conn.SendToHub()
	conn.ReceiveFromHub()

}

var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
