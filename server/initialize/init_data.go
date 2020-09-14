package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/enum"
	"github.com/bxcodec/faker/v3"
	uuid "github.com/satori/go.uuid"
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
		{Model: gorm.Model{ID: uint(enum.HousePassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "住宅客梯"},
		{Model: gorm.Model{ID: uint(enum.HouseFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "住宅货梯"},
		{Model: gorm.Model{ID: uint(enum.MarketPassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "商场客梯"},
		{Model: gorm.Model{ID: uint(enum.MarketFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "商场货梯"},
		{Model: gorm.Model{ID: uint(enum.HospitalPassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "医院客梯"},
		{Model: gorm.Model{ID: uint(enum.HospitalFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "医院货梯"},
		{Model: gorm.Model{ID: uint(enum.OfficePassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "写字楼客梯"},
		{Model: gorm.Model{ID: uint(enum.OfficeFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "写字楼货梯"},
		{Model: gorm.Model{ID: uint(enum.GovPassengerLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "政府客梯"},
		{Model: gorm.Model{ID: uint(enum.GovFreightLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftType), CategoryName: "政府货梯"},

		{Model: gorm.Model{ID: uint(enum.LiftMaintainRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftRecordType), CategoryName: "维保记录"},
		{Model: gorm.Model{ID: uint(enum.LiftUniCheckRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftRecordType), CategoryName: "年检记录"},
		{Model: gorm.Model{ID: uint(enum.LiftComplainRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftRecordType), CategoryName: "投诉建议"},

		{Model: gorm.Model{ID: uint(enum.LiftUserReportTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: "电梯乘客"},
		{Model: gorm.Model{ID: uint(enum.LiftAIDetectTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: "AI自检"},
		{Model: gorm.Model{ID: uint(enum.LiftMaintainTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: "维保单位"},
		{Model: gorm.Model{ID: uint(enum.LiftCheckTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: "年检单位"},
		{Model: gorm.Model{ID: uint(enum.LiftOwnerTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSourceType), CategoryName: "使用单位"},

		{Model: gorm.Model{ID: uint(enum.TroubleSolvedByAuto), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSolvedType), CategoryName: "自动修复"},
		{Model: gorm.Model{ID: uint(enum.TroubleSolvedByReplaceCell), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSolvedType), CategoryName: "替换元件"},
		{Model: gorm.Model{ID: uint(enum.TroubleSolvedByNone), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleSolvedType), CategoryName: "误报故障"},

		{Model: gorm.Model{ID: uint(enum.TroubleReasonAsHuman), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: "人为原因"},
		{Model: gorm.Model{ID: uint(enum.TroubleReasonAsMaintain), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: "维保不当"},
		{Model: gorm.Model{ID: uint(enum.TroubleReasonAsCheck), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: "年检上报"},
		{Model: gorm.Model{ID: uint(enum.TroubleReasonAsAge), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: "电梯老化"},
		{Model: gorm.Model{ID: uint(enum.TroubleReasonAsElectric), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftTroubleReasonType), CategoryName: "元件损坏"},

		{Model: gorm.Model{ID: uint(enum.DoorCloseFailed), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: "关门失败"},
		{Model: gorm.Model{ID: uint(enum.DoorOpenFailed), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: "开门失败"},
		{Model: gorm.Model{ID: uint(enum.DoorCloseSuccess), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: "关门正常"},
		{Model: gorm.Model{ID: uint(enum.DoorOpenSuccess), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: "开门正常"},
		{Model: gorm.Model{ID: uint(enum.DoorOpenAtUnderfloor), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectLiftDoorStatusType), CategoryName: "非平层开门"},

		{Model: gorm.Model{ID: uint(enum.DeviceModalV1), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceModalType), CategoryName: "智能设备V1"},
		{Model: gorm.Model{ID: uint(enum.DeviceModalV2), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceModalType), CategoryName: "智能设备V2"},
		{Model: gorm.Model{ID: uint(enum.DeviceModalV3), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceModalType), CategoryName: "智能设备V3"},

		{Model: gorm.Model{ID: uint(enum.DeviceUnderTest), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceStatusType), CategoryName: "测试中"},
		{Model: gorm.Model{ID: uint(enum.DeviceUnderCheck), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceStatusType), CategoryName: "检修中"},
		{Model: gorm.Model{ID: uint(enum.DeviceNormal), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceStatusType), CategoryName: "正常运行"},
		{Model: gorm.Model{ID: uint(enum.DeviceAbnormal), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceStatusType), CategoryName: "异常"},

		{Model: gorm.Model{ID: uint(enum.DeviceOnlineEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceEventType), CategoryName: "上线"},
		{Model: gorm.Model{ID: uint(enum.DeviceOfflineEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceEventType), CategoryName: "离线"},
		{Model: gorm.Model{ID: uint(enum.DevicePersonDetectEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceEventType), CategoryName: "人形检测"},
		{Model: gorm.Model{ID: uint(enum.DeviceMotorDetectEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectDeviceEventType), CategoryName: "电动车"},

		{Model: gorm.Model{ID: uint(enum.CompanyNormal), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyStatusType), CategoryName: "正常开业"},
		{Model: gorm.Model{ID: uint(enum.CompanyIrregular), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyStatusType), CategoryName: "状态异常"},

		{Model: gorm.Model{ID: uint(enum.CompanyMaintain), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: "维保公司"},
		{Model: gorm.Model{ID: uint(enum.CompanyCheck), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: "年检公司"},
		{Model: gorm.Model{ID: uint(enum.CompanyOwner), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: "使用单位"},
		{Model: gorm.Model{ID: uint(enum.CompanyInstall), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: "安装公司"},
		{Model: gorm.Model{ID: uint(enum.CompanySupervise), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: "监督机构"},
		{Model: gorm.Model{ID: uint(enum.CompanyPlatform), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectCompanyType), CategoryName: "运营平台"},

		{Model: gorm.Model{ID: uint(enum.UserLiftUse), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: "乘坐"},
		{Model: gorm.Model{ID: uint(enum.UserLiftMaintain), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: "维保"},
		{Model: gorm.Model{ID: uint(enum.UserLiftManage), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: "管理"},
		{Model: gorm.Model{ID: uint(enum.UserLiftCheck), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: "年检"},
		{Model: gorm.Model{ID: uint(enum.UserLiftInstall), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserLiftShipType), CategoryName: "安装"},

		{Model: gorm.Model{ID: uint(enum.UserCommon), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: "普通用户"},
		{Model: gorm.Model{ID: uint(enum.UserPresent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: "展示账户"},
		{Model: gorm.Model{ID: uint(enum.UserCompanyManager), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: "公司管理员"},
		{Model: gorm.Model{ID: uint(enum.UserCompanyEmployee), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: "公司员工"},
		{Model: gorm.Model{ID: uint(enum.UserAdmin), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: "管理员"},
		{Model: gorm.Model{ID: uint(enum.UserSuperAdmin), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectUserType), CategoryName: "超级管理员"},

		{Model: gorm.Model{ID: uint(enum.AddressTagBaby), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "婴儿"},
		{Model: gorm.Model{ID: uint(enum.AddressTagJunior), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "初中生"},
		{Model: gorm.Model{ID: uint(enum.AddressTagSenior), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "高中生"},
		{Model: gorm.Model{ID: uint(enum.AddressTagCollege), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "大学生"},
		{Model: gorm.Model{ID: uint(enum.AddressTagWorker), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "工人"},
		{Model: gorm.Model{ID: uint(enum.AddressTagDoctor), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "医生"},
		{Model: gorm.Model{ID: uint(enum.AddressTagTeacher), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "教师"},
		{Model: gorm.Model{ID: uint(enum.AddressTagFarmer), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "农名"},
		{Model: gorm.Model{ID: uint(enum.AddressTagDesigner), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "设计师"},
		{Model: gorm.Model{ID: uint(enum.AddressTagEngineer), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "工程师"},
		{Model: gorm.Model{ID: uint(enum.AddressTagDriver), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "司机"},
		{Model: gorm.Model{ID: uint(enum.AddressTagCook), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "厨师"},
		{Model: gorm.Model{ID: uint(enum.AddressTagBusinessman), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectAddressTagType), CategoryName: "商人"},

		{Model: gorm.Model{ID: uint(enum.MessageUrgent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "紧急消息"},
		{Model: gorm.Model{ID: uint(enum.MessageLiftChange), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "电梯变更"},
		{Model: gorm.Model{ID: uint(enum.MessageAssign), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "指派消息"},
		{Model: gorm.Model{ID: uint(enum.MessageDeviceEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "设备事件"},
		{Model: gorm.Model{ID: uint(enum.MessageNewLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "新增电梯"},
		{Model: gorm.Model{ID: uint(enum.MessageNewTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "新增故障"},
		{Model: gorm.Model{ID: uint(enum.MessageNewRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "新增记录"},
		{Model: gorm.Model{ID: uint(enum.MessageRecordDone), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "完成记录"},
		{Model: gorm.Model{ID: uint(enum.MessageTroubleFixed), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageType), CategoryName: "故障修复"},

		{Model: gorm.Model{ID: uint(enum.MessageSubjectLift), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: "电梯"},
		{Model: gorm.Model{ID: uint(enum.MessageSubjectLiftRecord), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: "电梯记录"},
		{Model: gorm.Model{ID: uint(enum.MessageSubjectLiftTrouble), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: "电梯故障"},
		{Model: gorm.Model{ID: uint(enum.MessageSubjectDeviceEvent), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: "设备事件"},
		{Model: gorm.Model{ID: uint(enum.MessageSubjectUser), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: "用户"},
		{Model: gorm.Model{ID: uint(enum.MessageSubjectCompany), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectMessageFromTargetType), CategoryName: "公司"},

		{Model: gorm.Model{ID: uint(enum.HealthTimeDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: "时间维度"},
		{Model: gorm.Model{ID: uint(enum.HealthMaintainDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: "维保维度"},
		{Model: gorm.Model{ID: uint(enum.HealthHumanDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: "人为维度"},
		{Model: gorm.Model{ID: uint(enum.HealthInnerDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: "内部维度"},
		{Model: gorm.Model{ID: uint(enum.HealthSensorDimension), CreatedAt: time.Now(), UpdatedAt: time.Now()},
			CategorySubjectId: uint(enum.SubjectHealthDimensionType), CategoryName: "数据维度"},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitRegion() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.Region{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.Region{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "东湖高新区"},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "江夏区"},
		{gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "蔡甸区"},
		{gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "东西湖区"},
		{gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "青山区"},
		{gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "汉南区"},
		{gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "新洲区"},
		{gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "武昌区"},
		{gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "洪山区"},
		{gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "江汉区"},
		{gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "江岸区"},
		{gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "汉阳区"},
		{gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "武汉市", "桥口区"},
		{gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄石市", "下陆区"},
		{gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄石市", "铁山区"},
		{gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄石市", "黄石港区"},
		{gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄石市", "西塞山区区"},
		{gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄石市", "阳新县"},
		{gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄石市", "大冶市"},
		{gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "十堰市", "茅箭区"},
		{gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "十堰市", "张湾区"},
		{gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "十堰市", "郧阳区"},
		{gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "十堰市", "郧西县"},
		{gorm.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "十堰市", "竹山县"},
		{gorm.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "十堰市", "竹溪县"},
		{gorm.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "十堰市", "房县"},
		{gorm.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "十堰市", "丹江口市"},
		{gorm.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "襄城区"},
		{gorm.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "樊城区"},
		{gorm.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "襄州区"},
		{gorm.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "南漳县"},
		{gorm.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "谷城县"},
		{gorm.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "保康县"},
		{gorm.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "枣阳市"},
		{gorm.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "宜城市"},
		{gorm.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "襄阳市", "老河口市"},
		{gorm.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆州市", "沙市区"},
		{gorm.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆州市", "荆州区"},
		{gorm.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆州市", "江陵县"},
		{gorm.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆州市", "公安县"},
		{gorm.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆州市", "松滋市"},
		{gorm.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆州市", "石首市"},
		{gorm.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆州市", "洪湖市"},
		{gorm.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆州市", "监利市"},
		{gorm.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "鄂州市", "华容区"},
		{gorm.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "鄂州市", "鄂城区"},
		{gorm.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "鄂州市", "梁子湖区"},
		{gorm.Model{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆门市", "东宝区"},
		{gorm.Model{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆门市", "掇刀区"},
		{gorm.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆门市", "沙洋县"},
		{gorm.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆门市", "钟祥市"},
		{gorm.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "荆门市", "京山市"},
		{gorm.Model{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "黄州区"},
		{gorm.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "团风县"},
		{gorm.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "红安县"},
		{gorm.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "罗田县"},
		{gorm.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "英山县"},
		{gorm.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "黄梅县"},
		{gorm.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "浠水县"},
		{gorm.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "蕲春县"},
		{gorm.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "武穴市"},
		{gorm.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "黄冈市", "麻城市"},
		{gorm.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "咸宁市", "咸安区"},
		{gorm.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "咸宁市", "嘉鱼县"},
		{gorm.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "咸宁市", "崇阳县"},
		{gorm.Model{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "咸宁市", "通城县"},
		{gorm.Model{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "咸宁市", "通山县"},
		{gorm.Model{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "咸宁市", "赤壁市"},
		{gorm.Model{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "孝感市", "孝南区"},
		{gorm.Model{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "孝感市", "孝昌县"},
		{gorm.Model{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "孝感市", "大悟县"},
		{gorm.Model{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "孝感市", "云梦县"},
		{gorm.Model{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "孝感市", "应城市"},
		{gorm.Model{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "孝感市", "汉川市"},
		{gorm.Model{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "孝感市", "安陆市"},
		{gorm.Model{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "随州市", "曾都区"},
		{gorm.Model{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "随州市", "随县"},
		{gorm.Model{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "随州市", "广水市"},
		{gorm.Model{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "恩施州", "恩施市"},
		{gorm.Model{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "恩施州", "利川市"},
		{gorm.Model{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "恩施州", "建始县"},
		{gorm.Model{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "恩施州", "巴东县"},
		{gorm.Model{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "恩施州", "宣恩县"},
		{gorm.Model{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "恩施州", "咸丰县"},
		{gorm.Model{ID: 85, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "恩施州", "来凤县"},
		{gorm.Model{ID: 86, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "恩施州", "鹤峰县"},
		{gorm.Model{ID: 87, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "西陵区"},
		{gorm.Model{ID: 88, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "伍家岗区"},
		{gorm.Model{ID: 89, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "点军区"},
		{gorm.Model{ID: 90, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "猇亭区"},
		{gorm.Model{ID: 91, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "夷陵区"},
		{gorm.Model{ID: 92, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "远安县"},
		{gorm.Model{ID: 93, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "兴山县"},
		{gorm.Model{ID: 94, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "秭归县"},
		{gorm.Model{ID: 95, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "长阳土家族"},
		{gorm.Model{ID: 96, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "五峰土家族"},
		{gorm.Model{ID: 97, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "宜都市"},
		{gorm.Model{ID: 98, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "当阳市"},
		{gorm.Model{ID: 99, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "宜昌市", "枝江市"},
		{gorm.Model{ID: 100, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "潜江市", ""},
		{gorm.Model{ID: 101, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "仙桃市", ""},
		{gorm.Model{ID: 102, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "神龙架林区", ""},
		{gorm.Model{ID: 103, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "湖北省", "天门市", ""},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

/**
many2many struct that auto generate
*/
type AddressTag struct {
	AddressID  uint
	CategoryID uint
}

func InitAddress() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.Address{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.Address{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 1, AddressName: "光谷一路与湖口一路交汇处统建天成美景",
			Location: "114.443241,30.490429", UserAmount: 11 * 120 * 3},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 1, AddressName: "光谷一路与湖口一路交汇处统建天成美雅",
			Location: "114.443702,30.49277", UserAmount: 9 * 100 * 3},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 1, AddressName: "高新大道与教育中路交汇处国创光谷上城",
			Location: "114.454216,30.491996", UserAmount: 15 * 120 * 3},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 1, AddressName: "关山大道221号保利花园一二期",
			Location: "114.42036,30.505016", UserAmount: 25 * 60 * 3},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 1, AddressName: "关山大道216号武汉三医院光谷院区",
			Location: "114.417766,30.504676", UserAmount: 5000},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 1, AddressName: "高新大道777号东湖高新政务中心",
			Location: "114.503733,30.496582", UserAmount: 3000},
		{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 1, AddressName: "关山大道332号保利时代北区",
			Location: "114.414381,30.497011", UserAmount: 11 * 120 * 3},
		{Model: gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 1, AddressName: "关山大道355号光谷K11购物艺术中心",
			Location: "114.419444,30.499009", UserAmount: 5000},
		{Model: gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 9, AddressName: "雄楚大道693号武汉工程大学武昌校区",
			Location: "114.397046,30.508165", UserAmount: 15000},
		{Model: gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 9, AddressName: "武珞路442号中南国际城",
			Location: "114.335713,30.540093", UserAmount: 3000},
		{Model: gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 9, AddressName: "野芷湖西路10号保利心语6期",
			Location: "114.330518,30.484954", UserAmount: 8 * 120 * 3},
		{Model: gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 9, AddressName: "黄家湖西路16号湖北中医药大学",
			Location: "114.276969,30.453788", UserAmount: 20000},
		{Model: gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 11, AddressName: "芦沟桥路28号武汉天地",
			Location: "114.317984,30.613122", UserAmount: 20 * 120 * 3},
		{Model: gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 11, AddressName: "中山大道1515号武汉天地壹方购物中心",
			Location: "114.316584,30.615446", UserAmount: 10000},
		{Model: gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 11, AddressName: "一元路汉景村1号武汉十六中学",
			Location: "114.307593,30.59924", UserAmount: 1200},
		{Model: gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 11, AddressName: "京汉大道1238号中城国际",
			Location: "114.304968,30.600858", UserAmount: 3000},
		{Model: gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 12, AddressName: "四新北路606号观澜国际",
			Location: "114.228067,30.531804", UserAmount: 11 * 120 * 3},
		{Model: gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 12, AddressName: "玉龙路168号中国铁建国际城",
			Location: "114.205507,30.565817", UserAmount: 15 * 120 * 3},
		{Model: gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 12, AddressName: "鹦鹉大道619号武汉国际博览中心",
			Location: "114.248387,30.513182", UserAmount: 5000},
		{Model: gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 12, AddressName: "琴断口街办事处汉阳大道528号赫山路十里华府",
			Location: "114.226765,30.566739", UserAmount: 17 * 120 * 3},
		{Model: gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, RegionId: 12, AddressName: "芳草六街与四新南路交叉路口往西约100米金地澜菲溪岸",
			Location: "114.217781,30.520061", UserAmount: 23 * 120 * 3},
	}

	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitAddressTags() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("address_id > 0").Delete(&AddressTag{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []AddressTag{
		{AddressID: 1, CategoryID: uint(enum.AddressTagBaby)},
		{AddressID: 1, CategoryID: uint(enum.AddressTagEngineer)},
		{AddressID: 1, CategoryID: uint(enum.AddressTagCollege)},

		{AddressID: 2, CategoryID: uint(enum.AddressTagJunior)},
		{AddressID: 2, CategoryID: uint(enum.AddressTagWorker)},

		{AddressID: 3, CategoryID: uint(enum.AddressTagBusinessman)},
		{AddressID: 3, CategoryID: uint(enum.AddressTagCook)},
		{AddressID: 3, CategoryID: uint(enum.AddressTagDoctor)},
		{AddressID: 3, CategoryID: uint(enum.AddressTagTeacher)},


		{AddressID: 4, CategoryID: uint(enum.AddressTagDesigner)},
		{AddressID: 4, CategoryID: uint(enum.AddressTagTeacher)},
		{AddressID: 4, CategoryID: uint(enum.AddressTagBaby)},

		{AddressID: 5, CategoryID: uint(enum.AddressTagWorker)},
		{AddressID: 5, CategoryID: uint(enum.AddressTagDesigner)},
		{AddressID: 5, CategoryID: uint(enum.AddressTagSenior)},
		{AddressID: 5, CategoryID: uint(enum.AddressTagJunior)},
		{AddressID: 5, CategoryID: uint(enum.AddressTagFarmer)},
		{AddressID: 5, CategoryID: uint(enum.AddressTagCook)},

		{AddressID: 6, CategoryID: uint(enum.AddressTagBaby)},
		{AddressID: 6, CategoryID: uint(enum.AddressTagEngineer)},

		{AddressID: 7, CategoryID: uint(enum.AddressTagBaby)},

		{AddressID: 8, CategoryID: uint(enum.AddressTagEngineer)},

		{AddressID: 9, CategoryID: uint(enum.AddressTagCollege)},

		{AddressID: 10, CategoryID: uint(enum.AddressTagTeacher)},
		{AddressID: 10, CategoryID: uint(enum.AddressTagDoctor)},
		{AddressID: 10, CategoryID: uint(enum.AddressTagDesigner)},

		{AddressID: 11, CategoryID: uint(enum.AddressTagCook)},
		{AddressID: 11, CategoryID: uint(enum.AddressTagEngineer)},

		{AddressID: 12, CategoryID: uint(enum.AddressTagFarmer)},
		{AddressID: 12, CategoryID: uint(enum.AddressTagBaby)},

		{AddressID: 13, CategoryID: uint(enum.AddressTagWorker)},
		{AddressID: 13, CategoryID: uint(enum.AddressTagEngineer)},

		{AddressID: 14, CategoryID: uint(enum.AddressTagDoctor)},
		{AddressID: 14, CategoryID: uint(enum.AddressTagTeacher)},

		{AddressID: 15, CategoryID: uint(enum.AddressTagCook)},
		{AddressID: 15, CategoryID: uint(enum.AddressTagWorker)},

		{AddressID: 16, CategoryID: uint(enum.AddressTagBaby)},
		{AddressID: 16, CategoryID: uint(enum.AddressTagJunior)},
		{AddressID: 16, CategoryID: uint(enum.AddressTagSenior)},
		{AddressID: 16, CategoryID: uint(enum.AddressTagCollege)},

		{AddressID: 17, CategoryID: uint(enum.AddressTagCook)},
		{AddressID: 17, CategoryID: uint(enum.AddressTagTeacher)},

		{AddressID: 18, CategoryID: uint(enum.AddressTagJunior)},
		{AddressID: 18, CategoryID: uint(enum.AddressTagCollege)},
		{AddressID: 18, CategoryID: uint(enum.AddressTagTeacher)},

		{AddressID: 19, CategoryID: uint(enum.AddressTagSenior)},
		{AddressID: 19, CategoryID: uint(enum.AddressTagWorker)},

		{AddressID: 20, CategoryID: uint(enum.AddressTagJunior)},
		{AddressID: 20, CategoryID: uint(enum.AddressTagSenior)},
		{AddressID: 20, CategoryID: uint(enum.AddressTagWorker)},
		{AddressID: 20, CategoryID: uint(enum.AddressTagEngineer)},

		{AddressID: 21, CategoryID: uint(enum.AddressTagFarmer)},
		{AddressID: 21, CategoryID: uint(enum.AddressTagCook)},
		{AddressID: 21, CategoryID: uint(enum.AddressTagDoctor)},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitSysAuthority() (err error) {
	tx := global.GVA_DB.Begin()
	insert := []model.SysAuthority{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "101", AuthorityName: "普通用户", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "201", AuthorityName: "公司管理账户", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "210", AuthorityName: "员工账户", ParentId: "201"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "301", AuthorityName: "监督单位管理账号", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "310", AuthorityName: "监督单位员工", ParentId: "301"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "401", AuthorityName: "运营平台管理账号", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "410", AuthorityName: "运营平台员工", ParentId: "401"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "501", AuthorityName: "物业管理账号", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "510", AuthorityName: "物业员工", ParentId: "501"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "超级管理员", ParentId: "0"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: "普通管理员", ParentId: "888"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "10000", AuthorityName: "演示账号", ParentId: "0"},
	}
	if tx.Create(&insert).Error != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitUser() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.SysUser{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.SysUser{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), PhoneNumber: "17612732732", RealName: faker.FirstName() + faker.LastName(),
			NickName: faker.LastName(), Avatar: faker.URL(), CompId:, CompanyId:, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "888"}
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitCompany() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.Company{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.Company{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉智能电梯有限公司",
			Alias: "武汉智能电梯有限公司", LegalPerson: "陈纯星", PhoneNumber: "027-87928666", Status: "正常营业",
			RegCode: "420111000041704", OrgCode: "178092678", CreditCode: "91420111178092678B",
			TaxCode: "91420111178092678B", Address: "湖北省武汉市洪山区武珞路718号", CategoryID: enum.CompanyMaintain},
	{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉华美东菱电梯有限公司",
			Alias: "武汉华美东菱电梯有限公司", LegalPerson: "戴俭勤", PhoneNumber: "027-88082935", Status: "正常营业",
			RegCode: "420100000230772", OrgCode: "71195427-3", CreditCode: "91420100711954273A",
			TaxCode: "91420100711954273A", Address: "湖北省武汉市洪山区花山镇附5号", CategoryID: enum.CompanyMaintain},
	{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉博瑞通达电梯有限公司",
			Alias: "博瑞通达", LegalPerson: "宋保山", PhoneNumber: "027-88229959", Status: "正常营业",
			RegCode: "420111000163221", OrgCode: "56835045X", CreditCode: "9142011156835045XL",
			TaxCode: "9142011156835045XL", Address: "湖北省武汉市洪山区和平街金地自在城(K2地块6期)A单元15层4-6号", CategoryID: enum.CompanyMaintain},

	{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉星奥电梯有限公司",
			Alias: "星奥电梯", LegalPerson: "黄亮平", PhoneNumber: "027-83536667", Status: "正常营业",
			RegCode: "420111000363899", OrgCode: "303529227", CreditCode: "91420111303529227Y",
			TaxCode: "91420111303529227Y", Address: "湖北省武汉市江汉区经济开发区江兴路25号A栋新金科技企业孵化器A704,A705室", CategoryID: enum.CompanyCheck},
	{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉利科电梯有限公司",
			Alias: "利科电梯", LegalPerson: "曾怀池", PhoneNumber: "13135672966", Status: "正常营业",
			RegCode: "420102000132629", OrgCode: "555021198", CreditCode: "91420116555021198J",
			TaxCode: "91420116555021198J", Address: "湖北省武汉市黄陂区三里桥街发展大道特一号9栋1单元", CategoryID: enum.CompanyCheck},
	{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉奥星电梯有限公司",
			Alias: "奥星电梯", LegalPerson: "邓友运", PhoneNumber: "027-83551522", Status: "正常营业",
			RegCode: "420102000016383", OrgCode: "764645052", CreditCode: "91420102764645052B",
			TaxCode: "91420102764645052B", Address: "湖北省武汉市江岸区二七路89号东立国际二期19栋1层29室", CategoryID: enum.CompanyCheck},

	{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉九井电梯有限公司",
			Alias: "九井电梯", LegalPerson: "崔锐", PhoneNumber: "13657223390", Status: "正常营业",
			RegCode: "420000000011766", OrgCode: "728301279", CreditCode: "914200007283012799",
			TaxCode: "914200007283012799", Address: "湖北省武汉市洪山区岳家嘴公务员小区14栋801室", CategoryID: enum.CompanyInstall},
	{Model: gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉惠之美物业服务有限公司",
			Alias: "惠之美物业", LegalPerson: "胡国平", PhoneNumber: "027-59318888", Status: "正常营业",
			RegCode: "914201127246641280", OrgCode: "724664128", CreditCode: "914201127246641280",
			TaxCode: "91420111178092678B", Address: "湖北省武汉市东西湖区银湖科技产业开发园18号", CategoryID: enum.CompanyInstall},
	{Model: gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉中汇日安电梯维保工程有限公司",
			Alias: "中汇日安电梯", LegalPerson: "王法玲", PhoneNumber: "027-87928644", Status: "正常营业",
			RegCode: "MA49G5A11", OrgCode: "MA49G5AB7", CreditCode: "91420100MA49G5AB7Q",
			TaxCode: "91420100MA49G5AB7Q", Address: "湖北省武汉市东湖高新技术开发区华师园北路18号博瀚科技光电产业园", CategoryID: enum.CompanyInstall},

	{Model: gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "湖北中楚物业股份有限公司",
			Alias: "中楚物业", LegalPerson: "成学荣", PhoneNumber: "027-51817371", Status: "正常营业",
			RegCode: "420106000048107", OrgCode: "679115081", CreditCode: "91420106679115081N",
			TaxCode: "91420106679115081N", Address: "湖北省武汉市武昌区中南路3号领秀中南26层5号", CategoryID: enum.CompanyOwner},
	{Model: gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "泛海物业管理武汉有限公司",
			Alias: "泛海物业", LegalPerson: "郑翼龙", PhoneNumber: "027-83871999", Status: "正常营业",
			RegCode: "420100000192187", OrgCode: "55501107X", CreditCode: "9142010055501107XX",
			TaxCode: "9142010055501107XX", Address: "湖北省武汉市江汉区云彩路198号泛海城市广场写字楼8层", CategoryID: enum.CompanyOwner},
	{Model: gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "开世艺商业管理武汉有限公司",
			Alias: "K11", LegalPerson: "朱滨", PhoneNumber: "027-87676765", Status: "正常营业",
			RegCode: "420100400101364", OrgCode: "MA4KU8MA0", CreditCode: "91420100MA4KU8MQ0K",
			TaxCode: "91420100MA4KU8MQ0K", Address: "湖北省武汉市东湖高新技术开发区关山大道特一号", CategoryID: enum.CompanyOwner},

	{Model: gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉市特种设备监督检验所",
			Alias: "武汉市特种设备监督检验所", LegalPerson: "邹少俊", PhoneNumber: "027-83872299", Status: "正常营业",
			RegCode: "0", OrgCode: "0", CreditCode: "12420100717957600G",
			TaxCode: "12420100717957600G", Address: "湖北省武汉市江岸区惠济二路11-2号", CategoryID: enum.CompanyOwner},
	{Model: gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉产品质量监督检验所",
			Alias: "武汉产品质量监督检验所", LegalPerson: "林建国", PhoneNumber: "027-87922266", Status: "正常营业",
			RegCode: "0", OrgCode: "0", CreditCode: "12420100441354365Y",
			TaxCode: "12420100441354365Y", Address: "湖北省武汉市东西湖区金银湖东二路5号", CategoryID: enum.CompanyOwner},


	{Model: gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "华中科技大学",
			Alias: "华中大", LegalPerson: "李元元", PhoneNumber: "027-87928666", Status: "正常营业",
			RegCode: "0", OrgCode: "0", CreditCode: "12100000441626842D",
			TaxCode: "12100000441626842D", Address: "湖北省武汉市洪山区珞瑜路1037号", CategoryID: enum.CompanyPlatform},
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
	err = InitRegion()
	err = InitAddress()
	err = InitAddressTags()
	err = InitSysAuthority()
	if err != nil {
		global.GVA_LOG.Error("initialize data failed", err)
	} else {
		global.GVA_LOG.Debug("initialize data success")
	}
}
