package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Hub struct {
	clients    []*Client
	register   chan *Client
	unRegister chan *Client
	mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unRegister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case c := <-hub.register:
			hub.connect(c)
		case c := <-hub.unRegister:
			hub.disconnect(c)
		}
	}
}

func (hub *Hub) BroadCast(msg interface{}, ignore *Client) {
	data, _ := json.Marshal(msg)
	for _, cl := range hub.clients {
		if cl != ignore {
			cl.outband <- data
		}
	}
}

func (hub *Hub) connect(c *Client) {
	log.Println("client connect : ", c.socket.RemoteAddr())
	hub.mutex.Lock()

	defer hub.mutex.Unlock()

	c.id = c.socket.RemoteAddr().String()
	hub.clients = append(hub.clients, c)

}

func (hub *Hub) disconnect(cl *Client) {
	log.Println("Client Disconnect", cl.socket.RemoteAddr())
	cl.Close()
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	i := -1
	for j, c := range hub.clients {
		if c.id == cl.id {
			i = j
			break
		}
	}

	copy(hub.clients[i:], hub.clients[i+1:])
	hub.clients[len(hub.clients)-1] = nil
	hub.clients = hub.clients[:len(hub.clients)-1]
}

func (hub *Hub) HandlerWs(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c := NewClient(hub, socket)

	hub.register <- c

	go c.Write()

}
