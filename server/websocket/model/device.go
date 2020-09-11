package model

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"strconv"
	"strings"
	"time"
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
	liftId := strings.Split(serial, "_")[1]
	id, _ := strconv.Atoi(liftId)
	err, lift := service.GetLift(uint(id))
	if err != nil {
		global.GVA_LOG.Warning("get lift info failed", err.Error())
		return
	}

	db := global.GVA_DB.Model(&lift.AdDevice);
	var eventTypeId uint
	var str string
	// update state
	if status == DeviceOnline {
		eventTypeId = 165
		str = "上线"
		db.Updates(map[string]interface{}{"online": true, "last_online_time": time.Now()})
	} else {
		eventTypeId = 164
		str = "离线"
		db.Updates(map[string]interface{}{"online": false, "last_online_time": time.Now()})
	}

	// create event
	// 164:off 165:on
	deviceEvent := model.AdDeviceEvent{DeviceID: lift.AdDevice.ID, TypeId: eventTypeId,
		Content: "设备于 " + time.Now().Format("2006-01-02 15:04:05") + str}
	service.CreateAdDeviceEvent(deviceEvent)
}

