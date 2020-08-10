package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Lift struct {
	gorm.Model
	NickName     string `json:"nickName" form:"nickName" gorm:"comment:'别名'"`
	Code         string `json:"code" form:"code" gorm:"comment:'编号'"`
	InstallerId  int `json:"installer_id" form:"installer_id"`
	MaintainerId int `json:"maintainer_id" form:"maintainer_id"`
	CheckerId    int `json:"checker_id" form:"checker_id"`
	OwnerId      int `json:"owner_id" form:"owner_id"`
	Installer    Company `json:"installer" form:"installer" gorm:"ForiegnKey:InstallerId;AssociationForiegnKey:ID;comment:'电梯安装公司'"`
	Maintainer   Company `json:"maintainer" form:"maintainer" gorm:"ForiegnKey:MaintainerId;AssociationForiegnKey:ID;comment:'电梯维保公司'"`
	Checker      Company `json:"checker" form:"checker" gorm:"ForiegnKey:CheckerId;AssociationForiegnKey:ID;comment:'电梯年检公司'"`
	Owner        Company `json:"owner" form:"owner" gorm:"ForiegnKey:OwnerId;AssociationForiegnKey:ID;comment:'电梯使用单位'"`
	FactoryTime  time.Time `json:"factory_time" form:"factory_time" gorm:"comment:'电梯出厂时间'"`
	InstallTime  time.Time `json:"install_time" form:"install_time" gorm:"comment:'电梯安装时间'"`
	CheckTime    time.Time `json:"check_time" form:"check_time" gorm:"comment:'电梯年检时间'"`
	LiftModelId  int `json:"lift_model_id" form:"lift_model_id"`
	LiftModel    LiftModel `json:"lift_model" form:"lift_model" gorm:"ForiegnKey:LiftModelId;AssociationForiegnKey:ID;comment:'电梯型号'"`
	CategoryId   int `json:"category_id" form:"category_id"`
	Category     Category `json:"category" form:"category" gorm:"ForiegnKey:CategoryId;AssociationForiegnKey:ID;comment:'电梯用途类别'"`
	FloorCount   int `json:"floor_count" form:"floor_count" gorm:"comment:'总楼层'"`
	Latitude     float32 `json:"latitude" form:"latitude" sql:"type:decimal(10,7)"`
	Longitude    float32 `json:"longitude" form:"longitude" sql:"type:decimal(10,7)"`
	AddressId    int `json:"address_id" form:"address_id"`
	Address      Address `json:"address" form:"address" gorm:"ForiegnKey:AddressId;AssociationForiegnKey:ID;comment:'电梯地址'"`
	RegionId     int `json:"region_id" form:"region_id"`
	Region       Region `json:"region" form:"region" gorm:"ForiegnKey:RegionId;AssociationForiegnKey:ID;comment:'电梯区域划分'"`
	Building     string `json:"building" form:"building" gorm:"comment:'楼栋号'"`
	Cell         int `json:"cell" form:"cell" gorm:"comment:'单元号'"`
	AdDeviceId   int `json:"ad_device_id" form:"ad_device_id"`
	AdDevice AdDevice `json:"ad_device" form:"ad_device" gorm:"ForiegnKey:AdDeviceId;AssociationForiegnKey:ID;comment:'电梯广告机设备'"`
}
