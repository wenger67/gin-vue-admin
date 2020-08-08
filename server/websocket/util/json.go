package util

import (
	"encoding/json"
	"gin-vue-admin/global"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadJson(input io.ReadCloser, obj interface{}) bool {
	buffer, err := ioutil.ReadAll(input)
	if err != nil {
		global.GVA_LOG.Warning("read json failed:", err)
		return false
	}
	err = json.Unmarshal(buffer, obj)
	if err != nil {
		global.GVA_LOG.Warning("parse json failed:", err)
		global.GVA_LOG.Warning("original body:", string(buffer))
		return false
	}
	return true
}

func SendJson(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		global.GVA_LOG.Warning("send json failed:", err)
	}
}
