package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub     *Hub
	id      string
	socket  *websocket.Conn
	outband chan []byte
}

func NewClient(h *Hub, socket *websocket.Conn) *Client {
	return &Client{
		Hub:     h,
		socket:  socket,
		outband: make(chan []byte),
	}

}

func (c *Client) Write() {
	for {
		select {
		case msg, ok := <-c.outband:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
			}
			c.socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func (c *Client) Close() {
	c.socket.Close()
	close(c.outband)
}
