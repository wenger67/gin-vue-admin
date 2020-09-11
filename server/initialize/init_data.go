package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/enum"
	"gorm.io/gorm"
	"time"
)

func Temp() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.CategorySubject{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.CategorySubject{}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}


func InitSubject() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.CategorySubject{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.CategorySubject{
		{gorm.Model{ID: uint(enum.SubjectLiftType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"电梯使用场景"},
		{gorm.Model{ID: uint(enum.SubjectLiftRecordType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"电梯记录"},
		{gorm.Model{ID: uint(enum.SubjectLiftTroubleSourceType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"电梯故障来源"},
		{gorm.Model{ID: uint(enum.SubjectLiftTroubleSolvedType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"电梯故障解除方式"},
		{gorm.Model{ID: uint(enum.SubjectLiftTroubleReasonType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"电梯故障原因"},
		{gorm.Model{ID: uint(enum.SubjectLiftDoorStatusType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"电梯厢门状态"},
		{gorm.Model{ID: uint(enum.SubjectDeviceModalType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"智能设备型号"},
		{gorm.Model{ID: uint(enum.SubjectDeviceStatusType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"智能设备状态"},
		{gorm.Model{ID: uint(enum.SubjectDeviceEventType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"智能设备事件"},
		{gorm.Model{ID: uint(enum.SubjectCompanyStatusType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"公司状态"},
		{gorm.Model{ID: uint(enum.SubjectCompanyType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"公司类别"},
		{gorm.Model{ID: uint(enum.SubjectUserLiftShipType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"用户电梯关系"},
		{gorm.Model{ID: uint(enum.SubjectUserType), CreatedAt: time.Now(), UpdatedAt: time.Now()}, "用户类型"},
		{gorm.Model{ID: uint(enum.SubjectAddressTagType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"地址标签"},
		{gorm.Model{ID: uint(enum.SubjectMessageType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"消息类型"},
		{gorm.Model{ID: uint(enum.SubjectMessageFromTargetType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"消息主体"},
		{gorm.Model{ID: uint(enum.SubjectHealthDimensionType), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			"大数据维度"},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitCategory() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.Category{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.Category{
		{Model:gorm.Model{ID: uint(enum.HousePassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.HouseFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MarketPassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MarketFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.HospitalPassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.HospitalFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.OfficePassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.OfficeFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.GovPassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.GovFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.LiftMaintainRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftRecordType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.LiftUniCheckRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftRecordType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.LiftComplainRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftRecordType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.LiftUserReportTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.LiftAIDetectTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.LiftMaintainTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.LiftCheckTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.LiftOwnerTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.TroubleSolvedByAuto), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSolvedType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.TroubleSolvedByReplaceCell), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSolvedType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.TroubleSolvedByNone), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSolvedType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.TroubleReasonAsHuman), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.TroubleReasonAsMaintain), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.TroubleReasonAsCheck), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.TroubleReasonAsAge), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.TroubleReasonAsElectric), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.DoorCloseFailed), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DoorOpenFailed), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DoorCloseSuccess), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DoorOpenSuccess), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DoorOpenAtUnderfloor), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.DeviceModalV1), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceModalType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DeviceModalV2), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceModalType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DeviceModalV3), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceModalType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.DeviceUnderTest), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceStatusType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DeviceUnderCheck), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceStatusType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DeviceNormal), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceStatusType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DeviceAbnormal), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceStatusType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.DeviceOnlineEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceEventType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DeviceOfflineEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceEventType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DevicePersonDetectEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceEventType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.DeviceMotorDetectEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceEventType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.CompanyNormal), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyStatusType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.CompanyIrregular), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyStatusType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.CompanyMaintain), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.CompanyCheck), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.CompanyOwner), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.CompanyInstall), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.CompanySupervise), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.CompanyPlatform), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.UserLiftUse), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserLiftMaintain), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserLiftManage), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserLiftCheck), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserLiftInstall), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.UserCommon), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserPresent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserCompanyManager), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserCompanyEmployee), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserAdmin), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.UserSuperAdmin), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.AddressTagBaby), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagJunior), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagSenior), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagCollege), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagWorker), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagDoctor), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagTeacher), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagFarmer), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagDesigner), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagEngineer), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagDriver), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagCook), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.AddressTagBusinessman), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.MessageUrgent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageLiftChange), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageAssign), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageDeviceEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageNewLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageNewTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageNewRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageRecordDone), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageTroubleFixed), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.MessageSubjectLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageSubjectLiftRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageSubjectLiftTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageSubjectDeviceEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageSubjectUser), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.MessageSubjectCompany), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: ""},

		{Model:gorm.Model{ID: uint(enum.HealthTimeDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.HealthMaintainDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.HealthHumanDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.HealthInnerDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: ""},
		{Model:gorm.Model{ID: uint(enum.HealthSensorDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: ""},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}


func InitData() {
	var err error
	err = InitSubject()
	err = InitCategory()
	if err != nil {
		global.GVA_LOG.Error("initialize data failed", err)
	} else {
		global.GVA_LOG.Debug("initialize data success")
	}
}
