package model

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"strconv"
	"strings"
)

const (
	DeviceOnline  = "online"
	DeviceOffline = "offline"
)

func UpdateDeviceStatus(serial string, status string) {
	global.GVA_LOG.Debug("update device[", serial, "] status:", status)
	if !strings.HasPrefix(serial, "LIFT_") {
		// TODO handle lift only, handle others later
		return
	}
	//_, err := mysql.Exec("update zx_device set update_at=now(), dev_alive=? where dev_number=?", status, serial)
	//if err != nil {
	//	global.GVA_LOG.Warning("update device status failed:", err)
	//}
	liftId := strings.Split(serial, "_")[1]
	id, _ := strconv.Atoi(liftId)
	err, lift := service.GetLift(uint(id))
	if err != nil {
		global.GVA_LOG.Warning("get lift info failed", err.Error())
		return
	}

	var online bool
	if status == DeviceOnline {
		online = true
	} else {
		online = false
	}
	global.GVA_DB.Debug().Model(&model.AdDevice{}).Where("id = ?", lift.AdDeviceId).Update("online", online)
}

