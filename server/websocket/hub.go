package websocket

import (
	"gin-vue-admin/global"
	"gin-vue-admin/websocket/model"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// send message to all clients.
	broadcast chan []byte

	// send message to target clients.
	send chan message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type message struct {
	target  []string
	content []byte
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		send:       make(chan message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			global.PantaLog.Info("client register")
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.removeClient(client)
			}
		case message := <-h.broadcast:
			global.PantaLog.Debug("broadcast msg to all clients")
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					h.removeClient(client)
				}
			}
		case message := <-h.send:
			global.PantaLog.Debug("send msg to target:", message.target)
			for client := range h.clients {
				//log.D("check client [", client.conn.RemoteAddr(), "/", client.serial, "]")
				if !contains(message.target, client.serial) {
					continue
				}
				global.PantaLog.Debug("send msg to [", client.conn.RemoteAddr(), "/", client.serial, "]")
				select {
				case client.send <- message.content:
				default:
					h.removeClient(client)
				}
			}
		}
	}
}

func (h *Hub) Send(target []string, msg []byte) {
	//log.D("send msg(", string(msg), ") to (", target, ")")
	h.send <- message{target: target, content: msg}
}

func (h *Hub) Broadcast(msg []byte) {
	//log.D("broadcast msg(", string(msg), ")")
	h.broadcast <- []byte(msg)
}

func (h *Hub) removeClient(client *Client) {
	if len(client.serial) > 0 {
		model.UpdateDeviceStatus(client.serial, model.DeviceOffline)
	}
	close(client.send)
	delete(h.clients, client)
}

func contains(c []string, t string) bool {
	for _, v := range c {
		if v == t {
			return true
		}
	}
	return false
}

