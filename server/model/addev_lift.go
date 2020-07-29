package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Lift struct {
	gorm.Model
	NickName     string
	Code         string
	InstallerId  int
	MaintainerId int
	CheckerId    int
	OwnerId      int
	Installer    Company
	Maintainer   Company
	Checker      Company
	Owner        Company
	FactoryTime  time.Time
	InstallTime  time.Time
	CheckTime    time.Time
	LiftModelId  int
	LiftModel    LiftModel
	CategoryId   int
	Category     Category
	FloorCount   int
	Latitude     float32 `sql:"type:decimal(10,7)"`
	Longitude    float32 `sql:"type:decimal(10,7)"`
	AddressId    int
	Address      Address
	RegionId     int
	Region       Region
	Building     string
	Cell         int
	AdDeviceId   int
}
