package model

import (
	"gorm.io/gorm"
	"time"
)

type Lift struct {
	gorm.Model
	NickName string `json:"nickName" form:"nickName" gorm:"comment:'别名'"`
	Code     string `json:"code" form:"code" gorm:"comment:'编号'"`
	// belong to
	InstallerId  int       `json:"installerId" form:"installerId"`
	Installer    Company   `json:"installer" form:"installer" gorm:"foreignKey:InstallerId;comment:'电梯安装公司'"`
	MaintainerId int       `json:"maintainerId" form:"maintainerId"`
	Maintainer   Company   `json:"maintainer" form:"maintainer" gorm:"foreignKey:MaintainerId;comment:'电梯维保公司'"`
	CheckerId    int       `json:"checkerId" form:"checkerId"`
	Checker      Company   `json:"checker" form:"checker" gorm:"ForeignKey:CheckerId;comment:'电梯年检公司'"`
	OwnerId      int       `json:"ownerId" form:"ownerId"`
	Owner        Company   `json:"owner" form:"owner" gorm:"ForeignKey:OwnerId;comment:'电梯使用单位'"`
	FactoryTime  time.Time `json:"factoryTime" form:"factoryTime" gorm:"comment:'电梯出厂时间'"`
	InstallTime  time.Time `json:"installTime" form:"installTime" gorm:"comment:'电梯安装时间'"`
	CheckTime    time.Time `json:"checkTime" form:"checkTime" gorm:"comment:'电梯年检时间'"`
	// belong to
	LiftModelId  int       `json:"liftModelId" form:"liftModelId"`
	LiftModel    LiftModel `json:"liftModel" form:"liftModel" gorm:"ForeignKey:LiftModelId;comment:'电梯型号'"`
	// belong to
	CategoryId int      `json:"categoryId" form:"categoryId"`
	Category   Category `json:"category" form:"category" gorm:"ForeignKey:CategoryId;comment:'电梯用途类别'"`
	// belong to
	AddressId  int     `json:"addressId" form:"addressId"`
	Address    Address `json:"address" form:"address" gorm:"ForeignKey:AddressId;comment:'电梯地址'"`
	FloorCount int     `json:"floorCount" form:"floorCount" gorm:"comment:'总楼层'"`
	Building   string  `json:"building" form:"building" gorm:"comment:'楼栋号'"`
	Cell       int     `json:"cell" form:"cell" gorm:"comment:'单元号'"`
	// has one
	AdDevice Device `json:"adDevice" form:"adDevice" gorm:"comment:'电梯广告机设备'"`
	Location string `json:"location" form:"location"`
}
