package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AdDevice struct {
	gorm.Model
	TypeId int `json:"typeId" form:"typeId"`
	Type Category `json:"type" form:"type" gorm:"ForeignKey:TypeId;AssociationForeignKey:ID;comment:'广告机型号'"`
	FactoryId int `json:"factoryId" form:"factoryId"`
	Factory Company `json:"factory" form:"factory" gorm:"ForeignKey:FactoryId;AssociationForeignKey:ID;comment:'广告机生产商'"`
	FactoryTime time.Time `json:"factoryTime" form:"factoryTime"`
	InstallTime time.Time `json:"installTime" form:"installTime"`
	StatusId int `json:"statusId" form:"statusId"`
	Status Category `json:"status" form:"status" gorm:"ForeignKey:StatusId;AssociationForeignKey:ID;comment:'设备状态类型'"`
	Online bool `json:"online" form:"online" gorm:"comment:'是否在线'"`
	LastOfflineTime time.Time `json:"lastOfflineTime" form:"lastOfflineTime" gorm:"comment:'上次离线时间'"`
	LastOnlineTime time.Time `json:"lastOnlineTime" form:"lastOnlineTime" gorm:"comment:'上次上线时间'"`
	Owners []SysUser `json:"owners" form:"owners" gorm:"many2many:ad_device_owners;comment:'设备责任人'"`
	Configs []AdDeviceConfig `json:"configs" form:"configs" gorm:"many2many:ad_device_config_relations;comment:'设备配置'"`
}