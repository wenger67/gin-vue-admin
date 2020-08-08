package websocket

import (
	"encoding/json"
	"gin-vue-admin/global"
	"gin-vue-admin/websocket/model"
	"gin-vue-admin/websocket/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// Client identifier
	serial string
}

type WSEvent struct {
	Target []string `json:"target"`
	What   string   `json:"event"`
	Data   string   `json:"data"`
}

func setReadDeadline(conn *websocket.Conn, d time.Duration) {
	err := conn.SetReadDeadline(time.Now().Add(d))
	if err != nil {
		global.GVA_LOG.Warning("set read deadline to [", conn.RemoteAddr(), "] failed:", err)
	}
}

func setWriteDeadline(conn *websocket.Conn, d time.Duration) {
	err := conn.SetWriteDeadline(time.Now().Add(d))
	if err != nil {
		global.GVA_LOG.Warning("set write deadline to [", conn.RemoteAddr(), "] failed:", err)
	}
}

func (c *Client) read() {
	defer func() {
		global.GVA_LOG.Debug("exit read loop [", c.conn.RemoteAddr(), "]")
		c.hub.unregister <- c
		util.SafeClose(c.conn)
	}()
	c.conn.SetReadLimit(maxMessageSize)
	setReadDeadline(c.conn, pongWait)
	c.conn.SetPongHandler(func(data string) error {
		global.GVA_LOG.Debug("pong handler:", data)
		setReadDeadline(c.conn, pongWait)
		return nil
	})
	event := &WSEvent{}
	for {
		_, buffer, err := c.conn.ReadMessage()
		if err == nil {
			err = json.Unmarshal(buffer, event)
		}
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				global.GVA_LOG.Error("unexpected close error while read msg:", err)
			} else {
				global.GVA_LOG.Warning("read json failed:", err)
			}
			break
		}
		global.GVA_LOG.Debug("event{", event.Target, ",", event.What, "} from ", c.conn.RemoteAddr())
		if event.Target == nil || len(event.Target) == 0 {
			global.GVA_LOG.Debug("broadcast event to all clients")
			c.hub.Broadcast(buffer)
		} else if event.Target[0] == "server" {
			global.GVA_LOG.Debug("device online:", event.Data, "<=>", c.conn.RemoteAddr())
			c.serial = event.Data
			model.UpdateDeviceStatus(c.serial, model.DeviceOnline)
		} else {
			c.hub.Send(event.Target, buffer)
		}
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		global.GVA_LOG.Debug("exit write loop [", c.conn.RemoteAddr(), "]")
		ticker.Stop()
		util.SafeClose(c.conn)
	}()
	for {
		select {
		case msg, ok := <-c.send:
			setWriteDeadline(c.conn, writeWait)
			if !ok {
				// The hub closed the channel.
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				global.GVA_LOG.Warning("write close message failed:", err)
				return
			}

			global.GVA_LOG.Debug("write message to [", c.conn.RemoteAddr(), "]:", string(msg))

			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				global.GVA_LOG.Warning("write message failed:", err)
				return
			}
		case tic := <-ticker.C:
			global.GVA_LOG.Debug("send ping message to [", c.conn.RemoteAddr(), "]")
			setWriteDeadline(c.conn, writeWait)
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte("tik->"+tic.String())); err != nil {
				global.GVA_LOG.Debug("write ping message failed:", err)
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, c *gin.Context) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.GVA_LOG.Debug(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	go client.write()
	go client.read()
}
