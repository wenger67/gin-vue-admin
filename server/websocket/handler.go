package websocket

import (
	"encoding/json"
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	hub = NewHub()
)

func init() {
	go hub.Run()
}

func ServeWS(c *gin.Context) {
	ServeWs(hub, c)
}

func NotifyDevice(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		global.GVA_LOG.Warning("read body failed:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	global.GVA_LOG.Debug("notify device with message:", string(buffer))

	event := &WSEvent{}
	_ = json.Unmarshal(buffer, event)

	if event.Target != nil {
		for i, v := range event.Target {
			event.Target[i] = strings.ToUpper(v)
		}
	}

	global.GVA_LOG.Debug("target:", event.Target)
	if event.Target != nil && len(event.Target) > 0 {
		hub.Send(event.Target, buffer)
	} else {
		hub.Broadcast(buffer)
	}
	w.WriteHeader(http.StatusOK)
}
