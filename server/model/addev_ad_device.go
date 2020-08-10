package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AdDevice struct {
	gorm.Model
	TypeId int `json:"type_id" form:"type_id"`
	Type Category `json:"type" form:"type" gorm:"ForiegnKey:TypeId;AssociationForiegnKey:ID;comment:'广告机型号'"`
	FactoryId int `json:"factory_id" form:"factory_id"`
	Factory Company `json:"factory" form:"factory" gorm:"ForiegnKey:FactoryId;AssociationForiegnKey:ID;comment:'广告机生产商'"`
	FactoryTime time.Time `json:"factory_time" form:"factory_time"`
	InstallTime time.Time `json:"install_time" form:"install_time"`
	StatusId int `json:"status_id" form:"status_id"`
	Status Category `json:"status" form:"status" gorm:"ForiegnKey:StatusId;AssociationForiegnKey:ID;comment:'设备状态类型'"`
	Online bool `json:"online" form:"online" gorm:"comment:'是否在线'"`
	LastOfflineTime time.Time `json:"last_offline_time" form:"last_offline_time" gorm:"上次离线时间"`
	LastOnlineTime time.Time `json:"last_online_time" form:"last_online_time" gorm:"上次上线时间"`
	Owners []UserAdmin `json:"owners" form:"owners" gorm:"many2many:ad_device_owners;comment:'设备责任人'"`
	Configs []AdDeviceConfig `json:"configs" form:"configs" gorm:"many2many:ad_device_config_relations;comment:'设备配置'"`
}