package model

import (
	"gin-vue-admin/global"
)

const (
	DeviceOnline  = "online"
	DeviceOffline = "offline"
)

func UpdateDeviceStatus(serial string, status string) {
	global.GVA_LOG.Debug("update device[", serial, "] status:", status)
	//_, err := mysql.Exec("update zx_device set update_at=now(), dev_alive=? where dev_number=?", status, serial)
	//if err != nil {
	//	global.GVA_LOG.Warning("update device status failed:", err)
	//}
}

func ResetDeviceStatus() {
	global.GVA_LOG.Debug("reset device status")
	//noinspection SqlWithoutWhere
	//_, err := mysql.Exec("update zx_device set update_at=now(), dev_alive=?", DeviceOffline)
	//if err != nil {
	//	global.GVA_LOG.Warning("reset device status failed:", err)
	//}
}

