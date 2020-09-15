package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/enum"
	"github.com/bxcodec/faker/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
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

type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string
	SysBaseMenuId           uint
}

type SysDataAuthorityId struct {
	SysAuthorityAuthorityId    string
	DataAuthorityIdAuthorityId string
}

type DeviceConfigRelation struct {
	DeviceID       uint
	DeviceConfigID uint
}

type DeviceOwner struct {
	DeviceID  uint
	SysUserID uint
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
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("authority_id > 0").Delete(&model.SysAuthority{})

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
		// super admin
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732732",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "宋江", NickName: "及时雨", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 15, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "888"},

		// platform
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732733",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "卢俊义", NickName: "玉麒麟", Avatar: faker.URL(),
			HasCompId: 15, CompanyId: 15, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "401"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732734",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "吴用", NickName: "智多星", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 15, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "410"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732735",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "公孙胜", NickName: "入云龙", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 15, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "410"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732736",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "关胜", NickName: "大刀", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 15, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "410"},

		// owner
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732737",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "林冲", NickName: "豹子头", Avatar: faker.URL(),
			HasCompId: 10, CompanyId: 10, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "501"},
		{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732738",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "秦明", NickName: "霹雳火", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 10, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "510"},
		{Model: gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732739",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "呼延灼", NickName: "双鞭", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 10, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "510"},
		{Model: gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732740",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "花荣", NickName: "小李广", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 10, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "510"},

		// check
		{Model: gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732741",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "柴进", NickName: "小旋风", Avatar: faker.URL(),
			HasCompId: 4, CompanyId: 4, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "201"},
		{Model: gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732742",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "李应", NickName: "扑天雕", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 4, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},
		{Model: gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732743",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "朱仝", NickName: "美髯公", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 4, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},
		{Model: gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732744",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "花和尚", NickName: "鲁智深", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 4, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},

		// install
		{Model: gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732745",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "武松", NickName: "行者", Avatar: faker.URL(),
			HasCompId: 7, CompanyId: 7, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "201"},
		{Model: gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732746",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "董平", NickName: "双枪将", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 7, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},
		{Model: gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732747",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "张清", NickName: "没羽箭", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 7, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},
		{Model: gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732748",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "杨志", NickName: "青面兽", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 7, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},

		// maintain
		{Model: gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732749",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "徐宁", NickName: "金枪手", Avatar: faker.URL(),
			HasCompId: 1, CompanyId: 1, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "201"},
		{Model: gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732750",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "索超", NickName: "急先锋", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 1, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},
		{Model: gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732751",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "戴宗", NickName: "神行太保", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 1, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},
		{Model: gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732752",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "刘唐", NickName: "赤发鬼", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 1, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "210"},

		// supervise
		{Model: gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732753",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "李逵", NickName: "黑旋风", Avatar: faker.URL(),
			HasCompId: 13, CompanyId: 13, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "301"},
		{Model: gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732754",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "史进", NickName: "九纹龙", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 13, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "310"},
		{Model: gorm.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732755",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "穆弘", NickName: "没遮拦", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 13, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "310"},
		{Model: gorm.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732756",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "雷横", NickName: "插翅虎", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 13, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "310"},

		// common user
		{Model: gorm.Model{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732757",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "李俊", NickName: "混江龙", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 0, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "101"},
		{Model: gorm.Model{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732758",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "阮小二", NickName: "立地太岁", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 0, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "101"},
		{Model: gorm.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732759",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "张横", NickName: "船火儿", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 0, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "101"},
		{Model: gorm.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(),
			PhoneNumber: "17612732760",
			Password:    "96e79218965eb72c92a549dd5a330112", RealName: "阮小五", NickName: "短命二郎", Avatar: faker.URL(),
			HasCompId: 0, CompanyId: 0, Address: "湖北省武汉市东湖高新区软件园A1", AuthorityId: "101"},
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
			TaxCode: "12420100717957600G", Address: "湖北省武汉市江岸区惠济二路11-2号", CategoryID: enum.CompanySupervise},
		{Model: gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "武汉产品质量监督检验所",
			Alias: "武汉产品质量监督检验所", LegalPerson: "林建国", PhoneNumber: "027-87922266", Status: "正常营业",
			RegCode: "0", OrgCode: "0", CreditCode: "12420100441354365Y",
			TaxCode: "12420100441354365Y", Address: "湖北省武汉市东西湖区金银湖东二路5号", CategoryID: enum.CompanySupervise},


		{Model: gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FullName: "华中科技大学",
			Alias: "华中大", LegalPerson: "李元元", PhoneNumber: "027-87928666", Status: "正常营业",
			RegCode: "0", OrgCode: "0", CreditCode: "12100000441626842D", TaxCode: "12100000441626842D",
			Address: "湖北省武汉市洪山区珞瑜路1037号", CategoryID: enum.CompanyPlatform},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitSysDictionaryDetail() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.SysDictionaryDetail{})

	status := new(bool)
	*status = true
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.SysDictionaryDetail{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "smallint", 1, status, 1, 2},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumint", 2, status, 2, 2},
		{gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "int", 3, status, 3, 2},
		{gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "bigint", 4, status, 4, 2},
		{gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "data", 0, status, 0, 3},
		{gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "time", 1, status, 1, 3},
		{gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "year", 2, status, 2, 3},
		{gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "datetime", 3, status, 3, 3},
		{gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "timestamp", 5, status, 5, 3},
		{gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "float", 0, status, 0, 4},
		{gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "double", 1, status, 1, 4},
		{gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "decimal", 2, status, 2, 4},
		{gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "char", 0, status, 0, 5},
		{gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "varchar", 1, status, 1, 5},
		{gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyblob", 2, status, 2, 5},
		{gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinytext", 3, status, 3, 5},
		{gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "text", 4, status, 4, 5},
		{gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "blob", 5, status, 5, 5},
		{gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumblob", 6, status, 6, 5},
		{gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumtext", 7, status, 7, 5},
		{gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longblob", 8, status, 8, 5},
		{gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longtext", 9, status, 9, 5},
		{gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyint", 0, status, 0, 6},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysDictionary() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.SysDictionary{})

	status := new(bool)
	*status = true
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.SysDictionary{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Desc: "性别字典"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Desc: "int类型对应的数据库类型"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Desc: "数据库时间日期类型"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Desc: "数据库浮点型"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Desc: "数据库字符串"},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Desc: "数据库bool类型"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitExaFileUploadAndDownload() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.ExaFileUploadAndDownload{})

	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.ExaFileUploadAndDownload{
		// record 1
		{Model: gorm.Model{ID: 1, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 428, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 428, time.Local)}, Name: "IMG_20200901_152619_428.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/702/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_152619_428.jpg",
			Tag: "jpg", RecordId: 1},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 306, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 306, time.Local)}, Name: "IMG_20200901_152619_306.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/702/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_152619_306.jpg",
			Tag: "jpg", RecordId: 1},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 551, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 551, time.Local)}, Name: "IMG_20200901_152619_551.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/702/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_152619_551.jpg",
			Tag: "jpg", RecordId: 1},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 0, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 0, time.Local)}, Name: "VID_20200901_152433.mp4",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/702/83d67c105d7f3fb976ca1c182a05a28d_VID_20200901_152433.mp4",
			Tag: "mp4", RecordId: 1},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 672, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 672, time.Local)}, Name: "IMG_20200901_152619_672.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/702/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_152619_672.jpg",
			Tag: "jpg", RecordId: 1},

		// record 4
		{Model: gorm.Model{ID: 6, CreatedAt: time.Date(2020, 9, 1, 7, 26, 4, 183, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 30, 9, 428, time.Local)}, Name: "IMG_20200901_152604_183.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/701/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_152604_183.jpg",
			Tag: "jpg", RecordId: 4},
		{Model: gorm.Model{ID: 7, CreatedAt: time.Date(2020, 9, 1, 7, 26, 4, 305, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 306, time.Local)}, Name: "IMG_20200901_152604_305.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/701/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_152604_305.jpg",
			Tag: "jpg", RecordId: 4},
		{Model: gorm.Model{ID: 8, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 551, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 551, time.Local)}, Name: "IMG_20200901_152604_432.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/701/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_152604_432.jpg",
			Tag: "jpg", RecordId: 4},
		{Model: gorm.Model{ID: 9, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 0, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 0, time.Local)}, Name: "VID_20200901_152433.mp4",
			Url: "http://127.0.0." +
				"1:8888/upload/2020_09_01/1/181/701/83d67c105d7f3fb976ca1c182a05a28d_VID_20200901_152433.mp4",
			Tag: "mp4", RecordId: 4},
		{Model: gorm.Model{ID: 10, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 672, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 672, time.Local)}, Name: "IMG_20200901_152604_555.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/701/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_152604_555.jpg",
			Tag: "jpg", RecordId: 4},

		// lift trouble1
		{Model: gorm.Model{ID: 11, CreatedAt: time.Date(2020, 9, 1, 7, 15, 9, 428, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 428, time.Local)}, Name: "IMG_20200901_151623_221.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/694/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_151623_221",
			Tag: "jpg", TroubleId: 1},
		{Model: gorm.Model{ID: 12, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 306, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 306, time.Local)}, Name: "IMG_20200901_151623_632.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/694/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_151623_632",
			Tag: "jpg", TroubleId: 1},
		{Model: gorm.Model{ID: 13, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 551, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 551, time.Local)}, Name: "IMG_20200901_151623_976.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/694/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_151623_976.jpg",
			Tag: "jpg", TroubleId: 1},
		{Model: gorm.Model{ID: 14, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 0, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 0, time.Local)}, Name: "VID_20200901_152433.mp4",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/694/83d67c105d7f3fb976ca1c182a05a28d_VID_20200901_152433.mp4",
			Tag: "mp4", TroubleId: 1},
		{Model: gorm.Model{ID: 15, CreatedAt: time.Date(2020, 9, 1, 7, 25, 9, 672, time.Local),
			UpdatedAt: time.Date(2020, 9, 1, 7, 25, 9, 672, time.Local)}, Name: "IMG_20200901_151624_113.jpg",
			Url: "http://127.0.0.1:8888/upload/2020_09_01/1/181/694/83d67c105d7f3fb976ca1c182a05a28d_IMG_20200901_151624_113.jpg",
			Tag: "jpg", TroubleId: 1},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitSysBaseMenus() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.SysBaseMenu{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.SysBaseMenu{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "仪表盘", Icon: "setting"}},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "about", Name: "about", Hidden: false, Component: "view/about/index.vue", Sort: 100,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "关于我们", Icon: "info"}},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "admin", Name: "superAdmin", Hidden: false, Component: "view/superAdmin/index.vue", Sort: 80,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "超级管理员", Icon: "user-solid"}},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "3",
			Path: "authority", Name: "authority", Hidden: false, Component: "view/superAdmin/authority/authority.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "角色管理", Icon: "s-custom"}},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "3",
			Path: "menu", Name: "menu", Hidden: false, Component: "view/superAdmin/menu/menu.vue", Sort: 2,
			Meta: model.Meta{KeepAlive: true, DefaultMenu: false, Title: "菜单管理", Icon: "s-order"}},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "3",
			Path: "api", Name: "api", Hidden: false, Component: "view/superAdmin/api/api.vue", Sort: 3,
			Meta: model.Meta{KeepAlive: true, DefaultMenu: false, Title: "api管理", Icon: "s-platform"}},
		{Model: gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "3",
			Path: "user", Name: "user", Hidden: false, Component: "view/superAdmin/user/user.vue", Sort: 4,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "用户管理", Icon: "coordinate"}},
		{Model: gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "person", Name: "person", Hidden: false, Component: "view/person/person.vue", Sort: 50,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "个人信息", Icon: "message-solid"}},

		{Model: gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "example", Name: "example", Hidden: false, Component: "view/example/index.vue", Sort: 110,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "示例文件", Icon: "s-management"}},
		{Model: gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "19",
			Path: "table", Name: "table", Hidden: false, Component: "view/example/table/table.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "表格示例", Icon: "s-order"}},
		{Model: gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "19",
			Path: "form", Name: "form", Hidden: false, Component: "view/example/form/form.vue", Sort: 2,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "表单示例", Icon: "document"}},
		{Model: gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "19",
			Path: "rte", Name: "rte", Hidden: false, Component: "view/example/rte/rte.vue", Sort: 3,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "富文本编辑器", Icon: "reading"}},
		{Model: gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "19",
			Path: "excel", Name: "excel", Hidden: false, Component: "view/example/excel/excel.vue", Sort: 4,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "excel导入导出", Icon: "s-marketing"}},
		{Model: gorm.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "19",
			Path: "upload", Name: "upload", Hidden: false, Component: "view/example/upload/upload.vue", Sort: 5,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "上传下载", Icon: "upload"}},
		{Model: gorm.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "19",
			Path: "breakpoint", Name: "breakpoint", Hidden: false, Component: "view/example/breakpoint/breakpoint." +
			"vue", Sort: 6, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "断点续传", Icon: "upload"}},
		{Model: gorm.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "19",
			Path: "customer", Name: "customer", Hidden: false, Component: "view/example/customer/customer.vue", Sort: 7,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "客户列表（资源示例）", Icon: "s-custom"}},

		{Model: gorm.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "systemTools", Name: "systemTools", Hidden: false, Component: "view/systemTools/index.vue", Sort: 90,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "系统工具", Icon: "s-cooperation"}},
		{Model: gorm.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "38",
			Path: "autoCode", Name: "autoCode", Hidden: false, Component: "view/systemTools/autoCode/index.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: true, DefaultMenu: false, Title: "代码生成器", Icon: "cpu"}},
		{Model: gorm.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "38",
			Path: "formCreate", Name: "formCreate", Hidden: false, Component: "view/systemTools/formCreate/index.vue", Sort: 2,
			Meta: model.Meta{KeepAlive: true, DefaultMenu: false, Title: "表单生成器", Icon: "magic-stick"}},
		{Model: gorm.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "38",
			Path: "system", Name: "system", Hidden: false, Component: "view/systemTools/system/system.vue", Sort: 3,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "系统配置", Icon: "s-operation"}},
		{Model: gorm.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "iconList", Name: "iconList", Hidden: false, Component: "view/iconList/index.vue", Sort: 10,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "图标集合", Icon: "star-on"}},
		{Model: gorm.Model{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "3",
			Path: "dictionary", Name: "dictionary", Hidden: false,
			Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "字典管理", Icon: "notebook-2"}},
		{Model: gorm.Model{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "3",
			Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Hidden: true,
			Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "字典详情", Icon: "s-order"}},
		{Model: gorm.Model{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "3",
			Path: "operation", Name: "operation", Hidden: false,
			Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "操作历史", Icon: "time"}},

		{Model: gorm.Model{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "category", Name: "category", Hidden: false, Component: "view/category/index.vue", Sort: 60,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "类别管理", Icon: "s-order"}},
		{Model: gorm.Model{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "54",
			Path: "subject", Name: "categorySubject", Hidden: false, Component: "view/category/subject/subject.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "类别主体", Icon: "s-shop"}},
		{Model: gorm.Model{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "54",
			Path: "categories", Name: "categoryItem", Hidden: false, Component: "view/category/categories/categories." +
			"vue", Sort: 2, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "类别列表", Icon: "money"}},

		{Model: gorm.Model{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "company", Name: "company", Hidden: false, Component: "view/company/index.vue", Sort: 20,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "公司管理", Icon: "school"}},
		{Model: gorm.Model{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "57",
			Path: "list", Name: "list", Hidden: false, Component: "view/company/list/list.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "公司列表", Icon: "s-fold"}},
		{Model: gorm.Model{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "57",
			Path: "employee", Name: "employee", Hidden: false, Component: "view/company/employee/employee.vue", Sort: 2,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "员工管理", Icon: "s-custom"}},

		{Model: gorm.Model{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "address", Name: "address", Hidden: false, Component: "view/address/index.vue", Sort: 70,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "地址管理", Icon: "add-location"}},
		{Model: gorm.Model{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "60",
			Path: "region", Name: "region", Hidden: false, Component: "view/address/region/region.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "区域划分", Icon: "notebook-1"}},
		{Model: gorm.Model{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "60",
			Path: "addressList", Name: "addressList", Hidden: false, Component: "view/address/list/address.vue", Sort: 2,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "常用地址", Icon: "medal-1"}},

		{Model: gorm.Model{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "lift", Name: "lift", Hidden: false, Component: "view/lift/index.vue", Sort: 30,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "电梯管理", Icon: "school"}},
		{Model: gorm.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "63",
			Path: "list", Name: "liftList", Hidden: false, Component: "view/lift/list/list.vue", Sort: 0,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "电梯列表", Icon: "s-fold"}},
		{Model: gorm.Model{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "63",
			Path: "model", Name: "liftModel", Hidden: false, Component: "view/lift/model/model.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "型号管理", Icon: "setting"}},
		{Model: gorm.Model{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "63",
			Path: "change", Name: "liftChange", Hidden: false, Component: "view/lift/change/change.vue", Sort: 2,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "信息变更历史", Icon: "info"}},
		{Model: gorm.Model{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "63",
			Path: "record", Name: "liftRecord", Hidden: false, Component: "view/lift/record/record.vue", Sort: 3,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "维保记录", Icon: "s-cooperation"}},
		{Model: gorm.Model{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "63",
			Path: "trouble", Name: "liftTrouble", Hidden: false, Component: "view/lift/trouble/trouble.vue", Sort: 4,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "故障救援记录", Icon: "s-management"}},

		{Model: gorm.Model{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "user", Name: "userHub", Hidden: false, Component: "view/user/index.vue", Sort: 21,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "用户管理", Icon: "user-solid"}},
		{Model: gorm.Model{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "69",
			Path: "list", Name: "userList", Hidden: false, Component: "view/user/list/list.vue", Sort: 0,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "用户列表", Icon: "s-order"}},

		{Model: gorm.Model{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "sdevice", Name: "sdevice", Hidden: false, Component: "view/sdevice/index.vue", Sort: 40,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "智能设备管理", Icon: "goods"}},
		{Model: gorm.Model{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "71",
			Path: "list", Name: "sdeviceList", Hidden: false, Component: "view/sdevice/list/list.vue", Sort: 0,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "设备列表", Icon: "s-tools"}},
		{Model: gorm.Model{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "71",
			Path: "config", Name: "sdeviceConfig", Hidden: false, Component: "view/sdevice/config/config.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "设备配置", Icon: "s-flag"}},
		{Model: gorm.Model{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "71",
			Path: "event", Name: "sdeviceEvent", Hidden: false, Component: "view/sdevice/event/event.vue", Sort: 2,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "设备事件", Icon: "setting"}},
		{Model: gorm.Model{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "71",
			Path: "data", Name: "sdeviceData", Hidden: false, Component: "view/sdevice/data/data.vue", Sort: 3,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "运行数据", Icon: "picture-outline-round"}},
		{Model: gorm.Model{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "63",
			Path: "liftDetail/:id", Name: "liftDetail", Hidden: true, Component: "view/lift/list/liftDetail.vue",
			Sort: 8, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "电梯详情", Icon: "setting"}},
		{Model: gorm.Model{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "message", Name: "message", Hidden: false, Component: "view/message/index.vue", Sort: 50,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "消息系统", Icon: "message"}},
		{Model: gorm.Model{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0",
			Path: "health", Name: "health", Hidden: false, Component: "view/health/index.vue", Sort: 45,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "健康系统", Icon: "headset"}},
		{Model: gorm.Model{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "78",
			Path: "healthSystem", Name: "healthSystem", Hidden: false, Component: "view/health/system/system.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "大数据预测", Icon: "s-data"}},
		{Model: gorm.Model{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "78",
			Path: "healthChange", Name: "healthChange", Hidden: false, Component: "view/health/change/change.vue",
			Sort: 2, Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "数据变动", Icon: "date"}},
		{Model: gorm.Model{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "77",
			Path: "messageList", Name: "messageList", Hidden: false, Component: "view/message/list/list.vue", Sort: 1,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "消息列表", Icon: "message"}},
		{Model: gorm.Model{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "63",
			Path: "userLift", Name: "userLift", Hidden: false, Component: "view/lift/assign/assign.vue", Sort: 10,
			Meta: model.Meta{KeepAlive: false, DefaultMenu: false, Title: "指派管理", Icon: "user-solid"}},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitExaCustomer() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.ExaCustomer{})
	tx := global.GVA_DB.Begin() // 开始事务
	insert := []model.ExaCustomer{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CustomerName: "测试客户", CustomerPhoneData: "1761111111", SysUserID: 1, SysUserAuthorityID: "888"},
	}
	if tx.Create(&insert).Error != nil { // 遇到错误时回滚事务
		tx.Rollback()
	}
	return tx.Commit().Error
}

func InitCasbinModel() (err error) {
	if global.GVA_DB.Migrator().HasTable("casbin_rule") {
		// first clear, hard mode
		_ = global.GVA_DB.Migrator().DropTable("casbin_rule")
	}

	if err := global.GVA_DB.Migrator().CreateTable(&gormadapter.CasbinRule{}); err != nil {
		return err
	}
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.CasbinModel{
		{"p", "8881", "/base/login", "POST"},
		{"p", "8881", "/base/register", "POST"},
		{"p", "8881", "/api/createApi", "POST"},
		{"p", "8881", "/api/getApiList", "POST"},
		{"p", "8881", "/api/getApiById", "POST"},
		{"p", "8881", "/api/deleteApi", "POST"},
		{"p", "8881", "/api/updateApi", "POST"},
		{"p", "8881", "/api/getAllApis", "POST"},
		{"p", "8881", "/authority/createAuthority", "POST"},
		{"p", "8881", "/authority/deleteAuthority", "POST"},
		{"p", "8881", "/authority/getAuthorityList", "POST"},
		{"p", "8881", "/authority/setDataAuthority", "POST"},
		{"p", "8881", "/menu/getMenu", "POST"},
		{"p", "8881", "/menu/getMenuList", "POST"},
		{"p", "8881", "/menu/addBaseMenu", "POST"},
		{"p", "8881", "/menu/getBaseMenuTree", "POST"},
		{"p", "8881", "/menu/addMenuAuthority", "POST"},
		{"p", "8881", "/menu/getMenuAuthority", "POST"},
		{"p", "8881", "/menu/deleteBaseMenu", "POST"},
		{"p", "8881", "/menu/updateBaseMenu", "POST"},
		{"p", "8881", "/menu/getBaseMenuById", "POST"},
		{"p", "8881", "/user/changePassword", "POST"},
		{"p", "8881", "/user/uploadHeaderImg", "POST"},
		{"p", "8881", "/user/getInfoList", "POST"},
		{"p", "8881", "/user/getUserList", "POST"},
		{"p", "8881", "/user/setUserAuthority", "POST"},
		{"p", "8881", "/fileUploadAndDownload/upload", "POST"},
		{"p", "8881", "/fileUploadAndDownload/getFileList", "POST"},
		{"p", "8881", "/fileUploadAndDownload/deleteFile", "POST"},
		{"p", "8881", "/casbin/updateCasbin", "POST"},
		{"p", "8881", "/casbin/getPolicyPathByAuthorityId", "POST"},
		{"p", "8881", "/jwt/jsonInBlacklist", "POST"},
		{"p", "8881", "/system/getSystemConfig", "POST"},
		{"p", "8881", "/system/setSystemConfig", "POST"},
		{"p", "8881", "/customer/customer", "POST"},
		{"p", "8881", "/customer/customer", "PUT"},
		{"p", "8881", "/customer/customer", "DELETE"},
		{"p", "8881", "/customer/customer", "GET"},
		{"p", "8881", "/customer/customerList", "GET"},
		{"p", "9528", "/base/login", "POST"},
		{"p", "9528", "/base/register", "POST"},
		{"p", "9528", "/api/createApi", "POST"},
		{"p", "9528", "/api/getApiList", "POST"},
		{"p", "9528", "/api/getApiById", "POST"},
		{"p", "9528", "/api/deleteApi", "POST"},
		{"p", "9528", "/api/updateApi", "POST"},
		{"p", "9528", "/api/getAllApis", "POST"},
		{"p", "9528", "/authority/createAuthority", "POST"},
		{"p", "9528", "/authority/deleteAuthority", "POST"},
		{"p", "9528", "/authority/getAuthorityList", "POST"},
		{"p", "9528", "/authority/setDataAuthority", "POST"},
		{"p", "9528", "/menu/getMenu", "POST"},
		{"p", "9528", "/menu/getMenuList", "POST"},
		{"p", "9528", "/menu/addBaseMenu", "POST"},
		{"p", "9528", "/menu/getBaseMenuTree", "POST"},
		{"p", "9528", "/menu/addMenuAuthority", "POST"},
		{"p", "9528", "/menu/getMenuAuthority", "POST"},
		{"p", "9528", "/menu/deleteBaseMenu", "POST"},
		{"p", "9528", "/menu/updateBaseMenu", "POST"},
		{"p", "9528", "/menu/getBaseMenuById", "POST"},
		{"p", "9528", "/user/changePassword", "POST"},
		{"p", "9528", "/user/uploadHeaderImg", "POST"},
		{"p", "9528", "/user/getInfoList", "POST"},
		{"p", "9528", "/user/getUserList", "POST"},
		{"p", "9528", "/user/setUserAuthority", "POST"},
		{"p", "9528", "/fileUploadAndDownload/upload", "POST"},
		{"p", "9528", "/fileUploadAndDownload/getFileList", "POST"},
		{"p", "9528", "/fileUploadAndDownload/deleteFile", "POST"},
		{"p", "9528", "/casbin/updateCasbin", "POST"},
		{"p", "9528", "/casbin/getPolicyPathByAuthorityId", "POST"},
		{"p", "9528", "/jwt/jsonInBlacklist", "POST"},
		{"p", "9528", "/system/getSystemConfig", "POST"},
		{"p", "9528", "/system/setSystemConfig", "POST"},
		{"p", "9528", "/customer/customer", "POST"},
		{"p", "9528", "/customer/customer", "PUT"},
		{"p", "9528", "/customer/customer", "DELETE"},
		{"p", "9528", "/customer/customer", "GET"},
		{"p", "9528", "/customer/customerList", "GET"},
		{"p", "9528", "/autoCode/createTemp", "POST"},
		{"p", "888", "/base/login", "POST"},
		{"p", "888", "/base/register", "POST"},
		{"p", "888", "/api/createApi", "POST"},
		{"p", "888", "/api/getApiList", "POST"},
		{"p", "888", "/api/getApiById", "POST"},
		{"p", "888", "/api/deleteApi", "POST"},
		{"p", "888", "/api/updateApi", "POST"},
		{"p", "888", "/api/getAllApis", "POST"},
		{"p", "888", "/authority/createAuthority", "POST"},
		{"p", "888", "/authority/deleteAuthority", "POST"},
		{"p", "888", "/authority/getAuthorityList", "POST"},
		{"p", "888", "/authority/setDataAuthority", "POST"},
		{"p", "888", "/authority/updateAuthority", "PUT"},
		{"p", "888", "/authority/copyAuthority", "POST"},
		{"p", "888", "/menu/getMenu", "POST"},
		{"p", "888", "/menu/getMenuList", "POST"},
		{"p", "888", "/menu/addBaseMenu", "POST"},
		{"p", "888", "/menu/getBaseMenuTree", "POST"},
		{"p", "888", "/menu/addMenuAuthority", "POST"},
		{"p", "888", "/menu/getMenuAuthority", "POST"},
		{"p", "888", "/menu/deleteBaseMenu", "POST"},
		{"p", "888", "/menu/updateBaseMenu", "POST"},
		{"p", "888", "/menu/getBaseMenuById", "POST"},
		{"p", "888", "/user/changePassword", "POST"},
		{"p", "888", "/user/uploadHeaderImg", "POST"},
		{"p", "888", "/user/getInfoList", "POST"},
		{"p", "888", "/user/getUserList", "POST"},
		{"p", "888", "/user/setUserAuthority", "POST"},
		{"p", "888", "/user/deleteUser", "DELETE"},
		{"p", "888", "/user/findUser", "GET"},
		{"p", "888", "/user/deleteUserList", "DELETE"},
		{"p", "888", "/user/createUser", "POST"},
		{"p", "888", "/fileUploadAndDownload/upload", "POST"},
		{"p", "888", "/fileUploadAndDownload/getFileList", "POST"},
		{"p", "888", "/fileUploadAndDownload/deleteFile", "POST"},
		{"p", "888", "/casbin/updateCasbin", "POST"},
		{"p", "888", "/casbin/getPolicyPathByAuthorityId", "POST"},
		{"p", "888", "/casbin/casbinTest/:pathParam", "GET"},
		{"p", "888", "/jwt/jsonInBlacklist", "POST"},
		{"p", "888", "/system/getSystemConfig", "POST"},
		{"p", "888", "/system/setSystemConfig", "POST"},
		{"p", "888", "/customer/customer", "POST"},
		{"p", "888", "/customer/customer", "PUT"},
		{"p", "888", "/customer/customer", "DELETE"},
		{"p", "888", "/customer/customer", "GET"},
		{"p", "888", "/customer/customerList", "GET"},
		{"p", "888", "/autoCode/createTemp", "POST"},
		{"p", "888", "/autoCode/getTables", "GET"},
		{"p", "888", "/autoCode/getDB", "GET"},
		{"p", "888", "/autoCode/getColume", "GET"},
		{"p", "888", "/sysDictionaryDetail/createSysDictionaryDetail", "POST"},
		{"p", "888", "/sysDictionaryDetail/deleteSysDictionaryDetail", "DELETE"},
		{"p", "888", "/sysDictionaryDetail/updateSysDictionaryDetail", "PUT"},
		{"p", "888", "/sysDictionaryDetail/findSysDictionaryDetail", "GET"},
		{"p", "888", "/sysDictionaryDetail/getSysDictionaryDetailList", "GET"},
		{"p", "888", "/sysDictionary/createSysDictionary", "POST"},
		{"p", "888", "/sysDictionary/deleteSysDictionary", "DELETE"},
		{"p", "888", "/sysDictionary/updateSysDictionary", "PUT"},
		{"p", "888", "/sysDictionary/findSysDictionary", "GET"},
		{"p", "888", "/sysDictionary/getSysDictionaryList", "GET"},
		{"p", "888", "/sysOperationRecord/createSysOperationRecord", "POST"},
		{"p", "888", "/sysOperationRecord/deleteSysOperationRecord", "DELETE"},
		{"p", "888", "/sysOperationRecord/updateSysOperationRecord", "PUT"},
		{"p", "888", "/sysOperationRecord/findSysOperationRecord", "GET"},
		{"p", "888", "/sysOperationRecord/getSysOperationRecordList", "GET"},
		{"p", "888", "/sysOperationRecord/deleteSysOperationRecordByIds", "DELETE"},
		{"p", "888", "/subject/getSubjectById", "POST"},
		{"p", "888", "/subject/getSubjectList", "POST"},
		{"p", "888", "/subject/createSubject", "POST"},
		{"p", "888", "/subject/updateSubject", "POST"},
		{"p", "888", "/subject/deleteSubject", "POST"},
		{"p", "888", "/subject/getAllSubjects", "POST"},
		{"p", "888", "/categories/createCategories", "POST"},
		{"p", "888", "/categories/deleteCategories", "DELETE"},
		{"p", "888", "/categories/deleteCategoriesByIds", "DELETE"},
		{"p", "888", "/categories/updateCategories", "PUT"},
		{"p", "888", "/categories/findCategories", "GET"},
		{"p", "888", "/categories/getCategoriesList", "GET"},
		{"p", "888", "/company/createCompany", "POST"},
		{"p", "888", "/company/deleteCompany", "DELETE"},
		{"p", "888", "/company/deleteCompanyByIds", "DELETE"},
		{"p", "888", "/company/updateCompany", "PUT"},
		{"p", "888", "/company/findCompany", "GET"},
		{"p", "888", "/company/getCompanyList", "GET"},
		{"p", "888", "/region/createRegion", "POST"},
		{"p", "888", "/region/deleteRegion", "DELETE"},
		{"p", "888", "/region/deleteRegionByIds", "DELETE"},
		{"p", "888", "/region/updateRegion", "PUT"},
		{"p", "888", "/region/findRegion", "GET"},
		{"p", "888", "/region/getRegionList", "GET"},
		{"p", "888", "/address/createAddress", "POST"},
		{"p", "888", "/address/deleteAddress", "DELETE"},
		{"p", "888", "/address/deleteAddressByIds", "DELETE"},
		{"p", "888", "/address/updateAddress", "PUT"},
		{"p", "888", "/address/findAddress", "GET"},
		{"p", "888", "/address/getAddressList", "GET"},
		{"p", "888", "/lift/createLift", "POST"},
		{"p", "888", "/lift/deleteLift", "DELETE"},
		{"p", "888", "/lift/deleteLiftByIds", "DELETE"},
		{"p", "888", "/lift/updateLift", "PUT"},
		{"p", "888", "/lift/findLift", "GET"},
		{"p", "888", "/lift/getLiftList", "GET"},
		{"p", "888", "/liftModel/createLiftModel", "POST"},
		{"p", "888", "/liftModel/deleteLiftModel", "DELETE"},
		{"p", "888", "/liftModel/deleteLiftModelByIds", "DELETE"},
		{"p", "888", "/liftModel/updateLiftModel", "PUT"},
		{"p", "888", "/liftModel/findLiftModel", "GET"},
		{"p", "888", "/liftModel/getLiftModelList", "GET"},
		{"p", "888", "/userAdmin/createUserAdmin", "POST"},
		{"p", "888", "/userAdmin/deleteUserAdmin", "DELETE"},
		{"p", "888", "/userAdmin/deleteUserAdminByIds", "DELETE"},
		{"p", "888", "/userAdmin/updateUserAdmin", "PUT"},
		{"p", "888", "/userAdmin/findUserAdmin", "GET"},
		{"p", "888", "/userAdmin/getUserAdminList", "GET"},
		{"p", "888", "/adDeviceData/createAdDeviceData", "POST"},
		{"p", "888", "/adDeviceData/deleteAdDeviceData", "DELETE"},
		{"p", "888", "/adDeviceData/deleteAdDeviceDataByIds", "DELETE"},
		{"p", "888", "/adDeviceData/updateAdDeviceData", "PUT"},
		{"p", "888", "/adDeviceData/findAdDeviceData", "GET"},
		{"p", "888", "/adDeviceData/getAdDeviceDataList", "GET"},
		{"p", "888", "/adDeviceConfig/createAdDeviceConfig", "POST"},
		{"p", "888", "/adDeviceConfig/deleteAdDeviceConfig", "DELETE"},
		{"p", "888", "/adDeviceConfig/deleteAdDeviceConfigByIds", "DELETE"},
		{"p", "888", "/adDeviceConfig/updateAdDeviceConfig", "PUT"},
		{"p", "888", "/adDeviceConfig/findAdDeviceConfig", "GET"},
		{"p", "888", "/adDeviceConfig/getAdDeviceConfigList", "GET"},
		{"p", "888", "/adDeviceEvent/createAdDeviceEvent", "POST"},
		{"p", "888", "/adDeviceEvent/deleteAdDeviceEvent", "DELETE"},
		{"p", "888", "/adDeviceEvent/deleteAdDeviceEventByIds", "DELETE"},
		{"p", "888", "/adDeviceEvent/updateAdDeviceEvent", "PUT"},
		{"p", "888", "/adDeviceEvent/findAdDeviceEvent", "GET"},
		{"p", "888", "/adDeviceEvent/getAdDeviceEventList", "GET"},
		{"p", "888", "/adDevice/createAdDevice", "POST"},
		{"p", "888", "/adDevice/deleteAdDevice", "DELETE"},
		{"p", "888", "/adDevice/deleteAdDeviceByIds", "DELETE"},
		{"p", "888", "/adDevice/updateAdDevice", "PUT"},
		{"p", "888", "/adDevice/findAdDevice", "GET"},
		{"p", "888", "/adDevice/getAdDeviceList", "GET"},
		{"p", "888", "/liftChange/createLiftChange", "POST"},
		{"p", "888", "/liftChange/deleteLiftChange", "DELETE"},
		{"p", "888", "/liftChange/deleteLiftChangeByIds", "DELETE"},
		{"p", "888", "/liftChange/updateLiftChange", "PUT"},
		{"p", "888", "/liftChange/findLiftChange", "GET"},
		{"p", "888", "/liftChange/getLiftChangeList", "GET"},
		{"p", "888", "/liftRecord/createLiftRecord", "POST"},
		{"p", "888", "/liftRecord/deleteLiftRecord", "DELETE"},
		{"p", "888", "/liftRecord/deleteLiftRecordByIds", "DELETE"},
		{"p", "888", "/liftRecord/updateLiftRecord", "PUT"},
		{"p", "888", "/liftRecord/findLiftRecord", "GET"},
		{"p", "888", "/liftRecord/getLiftRecordList", "GET"},
		{"p", "888", "/liftRecord/fillLiftRecord", "POST"},
		{"p", "888", "/liftRecord/reviewLiftRecord", "POST"},
		{"p", "888", "/liftTrouble/createLiftTrouble", "POST"},
		{"p", "888", "/liftTrouble/deleteLiftTrouble", "DELETE"},
		{"p", "888", "/liftTrouble/deleteLiftTroubleByIds", "DELETE"},
		{"p", "888", "/liftTrouble/updateLiftTrouble", "PUT"},
		{"p", "888", "/liftTrouble/findLiftTrouble", "GET"},
		{"p", "888", "/liftTrouble/getLiftTroubleList", "GET"},
		{"p", "888", "/userLift/createUserLift", "POST"},
		{"p", "888", "/userLift/deleteUserLift", "DELETE"},
		{"p", "888", "/userLift/deleteUserLiftByIds", "DELETE"},
		{"p", "888", "/userLift/updateUserLift", "PUT"},
		{"p", "888", "/userLift/findUserLift", "GET"},
		{"p", "888", "/userLift/getUserLiftList", "GET"},
		{"p", "888", "/message/createMessage", "POST"},
		{"p", "888", "/message/deleteMessage", "DELETE"},
		{"p", "888", "/message/deleteMessageByIds", "DELETE"},
		{"p", "888", "/message/updateMessage", "PUT"},
		{"p", "888", "/message/findMessage", "GET"},
		{"p", "888", "/message/getMessageList", "GET"},
		{"p", "888", "/healthSystem/createHealthSystem", "POST"},
		{"p", "888", "/healthSystem/deleteHealthSystem", "DELETE"},
		{"p", "888", "/healthSystem/deleteHealthSystemByIds", "DELETE"},
		{"p", "888", "/healthSystem/updateHealthSystem", "PUT"},
		{"p", "888", "/healthSystem/findHealthSystem", "GET"},
		{"p", "888", "/healthSystem/getHealthSystemList", "GET"},
		{"p", "888", "/healthChange/createHealthChange", "POST"},
		{"p", "888", "/healthChange/deleteHealthChange", "DELETE"},
		{"p", "888", "/healthChange/deleteHealthChangeByIds", "DELETE"},
		{"p", "888", "/healthChange/updateHealthChange", "PUT"},
		{"p", "888", "/healthChange/findHealthChange", "GET"},
		{"p", "888", "/healthChange/getHealthChangeList", "GET"},
	}
	if err = tx.Table("casbin_rule").Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitSysAuthorityMenus() (err error) {
	// first clear, hard mode
	global.GVA_DB.Where("sys_authority_authority_id > 0").Delete(SysAuthorityMenus{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []SysAuthorityMenus{
		{"101", 1},
		{"888", 1},
		{"888", 2},
		{"888", 3},
		{"888", 4},
		{"888", 5},
		{"888", 6},
		{"888", 17},
		{"888", 18},
		{"888", 19},
		{"888", 20},
		{"888", 21},
		{"888", 22},
		{"888", 23},
		{"888", 26},
		{"888", 33},
		{"888", 34},
		{"888", 38},
		{"888", 40},
		{"888", 41},
		{"888", 42},
		{"888", 50},
		{"888", 51},
		{"888", 52},
		{"888", 54},
		{"888", 55},
		{"888", 56},
		{"888", 57},
		{"888", 58},
		{"888", 59},
		{"888", 60},
		{"888", 61},
		{"888", 62},
		{"888", 63},
		{"888", 64},
		{"888", 65},
		{"888", 66},
		{"888", 67},
		{"888", 68},
		{"888", 71},
		{"888", 72},
		{"888", 73},
		{"888", 74},
		{"888", 75},
		{"888", 76},
		{"888", 77},
		{"888", 78},
		{"888", 79},
		{"888", 80},
		{"888", 81},
		{"888", 82},
		{"8881", 1},
		{"8881", 2},
		{"8881", 18},
		{"8881", 38},
		{"8881", 40},
		{"8881", 41},
		{"8881", 42},
		{"9528", 1},
		{"9528", 2},
		{"9528", 3},
		{"9528", 4},
		{"9528", 5},
		{"9528", 6},
		{"9528", 17},
		{"9528", 18},
		{"9528", 19},
		{"9528", 20},
		{"9528", 21},
		{"9528", 22},
		{"9528", 23},
		{"9528", 26},
		{"9528", 33},
		{"9528", 34},
		{"9528", 38},
		{"9528", 40},
		{"9528", 41},
		{"9528", 42},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitDeviceConfig() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.AdDeviceConfig{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.AdDeviceConfig{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "power_on", "06:00:00", "设备默认开机时间,早上6点整"},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "power_off", "22:00:00", "设备默认关机时间,晚上十点整"},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitLiftModel() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.LiftModel{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.LiftModel{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FactoryId: 7, Brand: "奥的斯",
			Modal: "TECE-300VF", Load: 1000},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FactoryId: 7, Brand: "奥的斯",
			Modal: "TECE-60VF", Load: 500},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FactoryId: 7, Brand: "奥的斯",
			Modal: "E311", Load: 800},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, FactoryId: 7, Brand: "奥的斯",
			Modal: "TECE-3VF", Load: 600},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitLift() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.Lift{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.Lift{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, NickName: "天成美景001", Code: "WU0001TCMJ001", InstallerId: 7,
			MaintainerId: 1, OwnerId: 10, CheckerId: 4, FactoryTime: time.Date(2020, 7, 15, 12, 0, 0, 0, time.Local),
			InstallTime: time.Date(2020, 7, 25, 12, 0, 0, 0, time.Local),
			CheckTime:   time.Date(2021, 5, 15, 12, 0, 0, 0, time.Local), LiftModelId: 1,
			CategoryId: uint(enum.HousePassengerLift), AddressId: 1, FloorCount: 32, Building: "1", Cell: 1},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, NickName: "天成美雅001", Code: "WU0001TCMY001", InstallerId: 7,
			MaintainerId: 1, OwnerId: 10, CheckerId: 4, FactoryTime: time.Date(2020, 6, 12, 12, 0, 0, 0, time.Local),
			InstallTime: time.Date(2020, 6, 28, 12, 0, 0, 0, time.Local),
			CheckTime:   time.Date(2021, 5, 15, 12, 0, 0, 0, time.Local), LiftModelId: 2,
			CategoryId: uint(enum.HousePassengerLift), AddressId: 2, FloorCount: 32, Building: "1", Cell: 1},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitDevice() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.Device{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.Device{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, TypeId: enum.DeviceModalV1,
			FactoryId: 15, FactoryTime: time.Date(2020, 5, 28, 12, 0, 0, 0, time.Local),
			InstallTime: time.Date(2020, 6, 28, 12, 0, 0, 0, time.Local), StatusId: enum.DeviceNormal,
			Online: false, LastOfflineTime: time.Date(2020, 6, 30, 12, 0, 0, 0, time.Local),
			LastOnlineTime: time.Date(2020, 6, 30, 8, 0, 0, 0, time.Local), LiftID: 1},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, TypeId: enum.DeviceModalV1,
			FactoryId: 15, FactoryTime: time.Date(2020, 5, 20, 12, 0, 0, 0, time.Local),
			InstallTime: time.Date(2020, 6, 27, 12, 0, 0, 0, time.Local), StatusId: enum.DeviceNormal,
			Online: false, LastOfflineTime: time.Date(2020, 6, 29, 12, 0, 0, 0, time.Local),
			LastOnlineTime: time.Date(2020, 6, 29, 8, 0, 0, 0, time.Local), LiftID: 2},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitDeviceConfigRelation() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("device_id > 0").Delete(&DeviceConfigRelation{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []DeviceConfigRelation{
		{DeviceID: 1, DeviceConfigID: 1},
		{DeviceID: 1, DeviceConfigID: 2},
		{DeviceID: 2, DeviceConfigID: 1},
		{DeviceID: 2, DeviceConfigID: 2},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitDeviceOwner() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("device_id > 0").Delete(&DeviceOwner{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []DeviceOwner{
		{DeviceID: 1, SysUserID: 3},
		{DeviceID: 1, SysUserID: 4},
		{DeviceID: 2, SysUserID: 4},
		{DeviceID: 2, SysUserID: 5},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitLiftRecord() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.LiftRecord{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.LiftRecord{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 1,
			CategoryId: uint(enum.LiftMaintainRecord), Content: "更换电梯们\n更换电梯灯\n更换电梯开关\n更管电梯牵引系统",
			StartTime: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			EndTime:   time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local),
			WorkerId:  18, RecorderId: 2, Progress: 4},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 1,
			CategoryId: uint(enum.LiftMaintainRecord), Progress: 1},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 2,
			CategoryId: uint(enum.LiftMaintainRecord), StartTime: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			WorkerId: 18, Progress: 2},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 2,
			CategoryId: uint(enum.LiftMaintainRecord), Content: "更换电梯们\n更换电梯灯\n更换电梯开关\n更管电梯牵引系统",
			StartTime: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			EndTime:   time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local),
			WorkerId:  18, Progress: 3},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitLiftTrouble() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.LiftTrouble{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.LiftTrouble{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 1,
			FromCategoryId: uint(enum.LiftMaintainTrouble), Content: "更换电梯们\n更换电梯灯\n更换电梯开关\n更管电梯牵引系统",
			StartTime: time.Date(2020, 8, 20, 12, 30, 0, 0, time.Local), StartUserId: 26,
			ResponseTime: time.Date(2020, 8, 20, 12, 35, 0, 0, time.Local), ResponseUserId: 6,
			SceneTime: time.Date(2020, 8, 20, 12, 38, 0, 0, time.Local), SceneUserId: 18,
			FixTime: time.Date(2020, 8, 20, 13, 35, 0, 0, time.Local), FixUserId: 18,
			FixCategoryId: uint(enum.TroubleSolvedByReplaceCell), ReasonCategoryId: uint(enum.TroubleReasonAsAge),
			RecorderId: 2, FeedbackContent: "good", FeedbackRate: 95, Progress: 7},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 1,
			FromCategoryId: uint(enum.LiftMaintainTrouble), StartTime: time.Date(2020, 8, 20, 12, 30, 0, 0, time.Local),
			StartUserId: 26, Progress: 1},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 1,
			FromCategoryId: uint(enum.LiftMaintainTrouble), StartTime: time.Date(2020, 8, 20, 12, 30, 0, 0, time.Local), StartUserId: 26,
			ResponseTime: time.Date(2020, 8, 20, 12, 35, 0, 0, time.Local), ResponseUserId: 6,
			Progress: 2},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 2,
			FromCategoryId: uint(enum.LiftMaintainTrouble), StartTime: time.Date(2020, 8, 20, 12, 30, 0, 0, time.Local), StartUserId: 26,
			ResponseTime: time.Date(2020, 8, 20, 12, 35, 0, 0, time.Local), ResponseUserId: 6,
			SceneTime: time.Date(2020, 8, 20, 12, 38, 0, 0, time.Local), SceneUserId: 18,
			Progress: 3},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 2,
			FromCategoryId: uint(enum.LiftMaintainTrouble), StartTime: time.Date(2020, 8, 20, 12, 30, 0, 0, time.Local), StartUserId: 26,
			ResponseTime: time.Date(2020, 8, 20, 12, 35, 0, 0, time.Local), ResponseUserId: 6,
			SceneTime: time.Date(2020, 8, 20, 12, 38, 0, 0, time.Local), SceneUserId: 18,
			FixTime: time.Date(2020, 8, 20, 13, 35, 0, 0, time.Local), FixUserId: 18,
			FixCategoryId: uint(enum.TroubleSolvedByReplaceCell), ReasonCategoryId: uint(enum.TroubleReasonAsAge),
			Progress: 4},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Date(2020, 8, 20, 12, 23, 0, 0, time.Local),
			UpdatedAt: time.Date(2020, 8, 20, 14, 20, 0, 0, time.Local)}, LiftId: 2,
			FromCategoryId: uint(enum.LiftMaintainTrouble), Content: "更换电梯们\n更换电梯灯\n更换电梯开关\n更管电梯牵引系统",
			StartTime: time.Date(2020, 8, 20, 12, 30, 0, 0, time.Local), StartUserId: 26,
			ResponseTime: time.Date(2020, 8, 20, 12, 35, 0, 0, time.Local), ResponseUserId: 6,
			SceneTime: time.Date(2020, 8, 20, 12, 38, 0, 0, time.Local), SceneUserId: 18,
			FixTime: time.Date(2020, 8, 20, 13, 35, 0, 0, time.Local), FixUserId: 18,
			FixCategoryId: uint(enum.TroubleSolvedByReplaceCell), ReasonCategoryId: uint(enum.TroubleReasonAsAge),
			FeedbackContent: "good", FeedbackRate: 95, Progress: 6},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func InitHealthSystem() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.HealthSystem{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.HealthSystem{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 1, TimeDimension: 200,
			MaintainDimension: 200, HumanDimension: 200, InnerDimension: 200, SensorDimension: 200},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 2, TimeDimension: 200,
			MaintainDimension: 200, HumanDimension: 200, InnerDimension: 200, SensorDimension: 200},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func InitHealthChange() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.HealthChange{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.HealthChange{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 1,
			DimensionId: uint(enum.HealthTimeDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 1,
			DimensionId: uint(enum.HealthHumanDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 1,
			DimensionId: uint(enum.HealthInnerDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 1,
			DimensionId: uint(enum.HealthMaintainDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 1,
			DimensionId: uint(enum.HealthSensorDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 2,
			DimensionId: uint(enum.HealthTimeDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 2,
			DimensionId: uint(enum.HealthHumanDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 2,
			DimensionId: uint(enum.HealthInnerDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 2,
			DimensionId: uint(enum.HealthMaintainDimension), Content: "Initial", Score: 200},
		{Model: gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, LiftId: 2,
			DimensionId: uint(enum.HealthSensorDimension), Content: "Initial", Score: 200},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func InitMessage() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.Message{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.Message{}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitSysApi() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.SysApi{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.SysApi{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/base/login",
			Description: "用户登录", ApiGroup: "base", Method: "POST"},
		{Model: gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/base/register",
			Description: "用户注册", ApiGroup: "base", Method: "POST"},
		{Model: gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/createApi",
			Description: "创建api", ApiGroup: "api", Method: "POST"},
		{Model: gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiList",
			Description: "获取api列表", ApiGroup: "api", Method: "POST"},
		{Model: gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiById",
			Description: "获取api详细信息", ApiGroup: "api", Method: "POST"},
		{Model: gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApi",
			Description: "删除Api", ApiGroup: "api", Method: "POST"},
		{Model: gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/updateApi",
			Description: "更新Api", ApiGroup: "api", Method: "POST"},
		{Model: gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/createAuthority", Description: "创建角色", ApiGroup: "authority", Method: "POST"},
		{Model: gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/deleteAuthority", Description: "删除角色", ApiGroup: "authority", Method: "POST"},
		{Model: gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/getAuthorityList", Description: "获取角色列表", ApiGroup: "authority", Method: "POST"},
		{Model: gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenu", Description: "获取菜单树",
			ApiGroup: "menu", Method: "POST"},
		{Model: gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuList",
			Description: "分页获取基础menu列表",
			ApiGroup:    "menu", Method: "POST"},
		{Model: gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addBaseMenu", Description: "新增菜单",
			ApiGroup: "menu", Method: "POST"},
		{Model: gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuTree",
			Description: "获取用户动态路由",
			ApiGroup:    "menu", Method: "POST"},
		{Model: gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addMenuAuthority",
			Description: "增加menu和角色关联关系", ApiGroup: "menu", Method: "POST"},
		{Model: gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuAuthority",
			Description: "获取指定角色menu", ApiGroup: "menu", Method: "POST"},
		{Model: gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/deleteBaseMenu", Description: "删除菜单",
			ApiGroup: "menu", Method: "POST"},
		{Model: gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/updateBaseMenu", Description: "更新菜单",
			ApiGroup: "menu", Method: "POST"},
		{Model: gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuById",
			Description: "根据id获取菜单",
			ApiGroup:    "menu", Method: "POST"},
		{Model: gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/changePassword", Description: "修改密码",
			ApiGroup: "user", Method: "POST"},
		{Model: gorm.Model{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/uploadHeaderImg", Description: "上传头像",
			ApiGroup: "user", Method: "POST"},
		{Model: gorm.Model{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/getInfoList", Description: "分页获取用户列表",
			ApiGroup: "user", Method: "POST"},
		{Model: gorm.Model{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/getUserList", Description: "获取用户列表",
			ApiGroup: "user", Method: "POST"},
		{Model: gorm.Model{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserAuthority",
			Description: "修改用户角色", ApiGroup: "user", Method: "POST"},
		{Model: gorm.Model{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/upload",
			Description: "文件上传示例", ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: gorm.Model{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileUploadAndDownload/getFileList",
			Description: "获取上传文件列表", ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: gorm.Model{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/updateCasbin",
			Description: "更改角色api权限",
			ApiGroup:    "casbin", Method: "POST"},
		{Model: gorm.Model{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/getPolicyPathByAuthorityId",
			Description: "获取权限列表", ApiGroup: "casbin", Method: "POST"},
		{Model: gorm.Model{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/fileUploadAndDownload/deleteFile", Description: "删除文件", ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: gorm.Model{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/jwt/jsonInBlacklist",
			Description: "jwt加入黑名单", ApiGroup: "jwt", Method: "POST"},
		{Model: gorm.Model{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/setDataAuthority",
			Description: "设置角色资源权限", ApiGroup: "authority", Method: "POST"},
		{Model: gorm.Model{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getSystemConfig",
			Description: "获取配置文件内容", ApiGroup: "system", Method: "POST"},
		{Model: gorm.Model{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/setSystemConfig",
			Description: "设置配置文件内容", ApiGroup: "system", Method: "POST"},
		{Model: gorm.Model{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "创建客户",
			ApiGroup: "customer", Method: "POST"},
		{Model: gorm.Model{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "更新客户",
			ApiGroup: "customer", Method: "PUT"},
		{Model: gorm.Model{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "删除客户",
			ApiGroup: "customer", Method: "DELETE"},
		{Model: gorm.Model{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "获取单一客户",
			ApiGroup: "customer", Method: "GET"},
		{Model: gorm.Model{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customerList",
			Description: "获取客户列表", ApiGroup: "customer", Method: "GET"},
		{Model: gorm.Model{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/casbinTest/:pathParam",
			Description: "RESTFUL模式测试", ApiGroup: "casbin", Method: "GET"},
		{Model: gorm.Model{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/createTemp",
			Description: "自动化代码", ApiGroup: "autoCode", Method: "POST"},
		{Model: gorm.Model{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/updateAuthority",
			Description: "更新角色信息", ApiGroup: "authority", Method: "PUT"},
		{Model: gorm.Model{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/copyAuthority",
			Description: "拷贝角色", ApiGroup: "authority", Method: "POST"},
		{Model: gorm.Model{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/deleteUser", Description: "删除用户",
			ApiGroup: "user", Method: "DELETE"},
		{Model: gorm.Model{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容", ApiGroup: "sysDictionaryDetail", Method: "POST"},
		{Model: gorm.Model{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容", ApiGroup: "sysDictionaryDetail", Method: "DELETE"},
		{Model: gorm.Model{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容", ApiGroup: "sysDictionaryDetail", Method: "PUT"},
		{Model: gorm.Model{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容", ApiGroup: "sysDictionaryDetail", Method: "GET"},
		{Model: gorm.Model{ID: 85, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表", ApiGroup: "sysDictionaryDetail", Method: "GET"},
		{Model: gorm.Model{ID: 86, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/createSysDictionary",
			Description: "新增字典", ApiGroup: "sysDictionary", Method: "POST"},
		{Model: gorm.Model{ID: 87, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/deleteSysDictionary",
			Description: "删除字典", ApiGroup: "sysDictionary", Method: "DELETE"},
		{Model: gorm.Model{ID: 88, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/sysDictionary/updateSysDictionary",
			Description: "更新字典", ApiGroup: "sysDictionary", Method: "PUT"},
		{Model: gorm.Model{ID: 89, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysDictionary/findSysDictionary", Description: "根据ID获取字典", ApiGroup: "sysDictionary", Method: "GET"},
		{Model: gorm.Model{ID: 90, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysDictionary/getSysDictionaryList", Description: "获取字典列表", ApiGroup: "sysDictionary", Method: "GET"},
		{Model: gorm.Model{ID: 91, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录", ApiGroup: "sysOperationRecord", Method: "POST"},
		{Model: gorm.Model{ID: 92, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录", ApiGroup: "sysOperationRecord", Method: "DELETE"},
		{Model: gorm.Model{ID: 93, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysOperationRecord/updateSysOperationRecord", Description: "更新操作记录", ApiGroup: "sysOperationRecord", Method: "PUT"},
		{Model: gorm.Model{ID: 94, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录", ApiGroup: "sysOperationRecord", Method: "GET"},
		{Model: gorm.Model{ID: 95, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表", ApiGroup: "sysOperationRecord", Method: "GET"},
		{Model: gorm.Model{ID: 96, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getTables", Description: "获取数据库表",
			ApiGroup: "autoCode", Method: "GET"},
		{Model: gorm.Model{ID: 97, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getDB", Description: "获取所有数据库",
			ApiGroup: "autoCode", Method: "GET"},
		{Model: gorm.Model{ID: 98, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/autoCode/getColume",
			Description: "获取所选table的所有字段", ApiGroup: "autoCode", Method: "GET"},
		{Model: gorm.Model{ID: 99, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史", ApiGroup: "sysOperationRecord", Method: "DELETE"},
		{Model: gorm.Model{ID: 106, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/subject/getSubjectById",
			Description: "getSubjectById", ApiGroup: "subject", Method: "POST"},
		{Model: gorm.Model{ID: 107, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/subject/getSubjectList",
			Description: "getSubjectList", ApiGroup: "subject", Method: "POST"},
		{Model: gorm.Model{ID: 108, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/subject/createSubject",
			Description: "createSubject", ApiGroup: "subject", Method: "POST"},
		{Model: gorm.Model{ID: 109, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/subject/updateSubject",
			Description: "updateSubject", ApiGroup: "subject", Method: "POST"},
		{Model: gorm.Model{ID: 110, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/subject/deleteSubject",
			Description: "deleteSubject", ApiGroup: "subject", Method: "POST"},
		{Model: gorm.Model{ID: 111, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/subject/getAllSubjects",
			Description: "getAllSubjects", ApiGroup: "subject", Method: "POST"},
		{Model: gorm.Model{ID: 112, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/categories/createCategories",
			Description: "新增categories表", ApiGroup: "categories", Method: "POST"},
		{Model: gorm.Model{ID: 113, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/categories/deleteCategories",
			Description: "删除categories表", ApiGroup: "categories", Method: "DELETE"},
		{Model: gorm.Model{ID: 114, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/categories/deleteCategoriesByIds", Description: "批量删除categories表", ApiGroup: "categories", Method: "DELETE"},
		{Model: gorm.Model{ID: 115, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/categories/updateCategories",
			Description: "更新categories表", ApiGroup: "categories", Method: "PUT"},
		{Model: gorm.Model{ID: 116, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/categories/findCategories",
			Description: "根据ID获取categories表", ApiGroup: "categories", Method: "GET"},
		{Model: gorm.Model{ID: 117, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/categories/getCategoriesList",
			Description: "获取categories表列表", ApiGroup: "categories", Method: "GET"},
		{Model: gorm.Model{ID: 118, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/company/createCompany",
			Description: "新增company表", ApiGroup: "company", Method: "POST"},
		{Model: gorm.Model{ID: 119, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/company/deleteCompany",
			Description: "删除company表", ApiGroup: "company", Method: "DELETE"},
		{Model: gorm.Model{ID: 120, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/company/deleteCompanyByIds",
			Description: "批量删除company表", ApiGroup: "company", Method: "DELETE"},
		{Model: gorm.Model{ID: 121, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/company/updateCompany",
			Description: "更新company表", ApiGroup: "company", Method: "PUT"},
		{Model: gorm.Model{ID: 122, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/company/findCompany",
			Description: "根据ID获取company表", ApiGroup: "company", Method: "GET"},
		{Model: gorm.Model{ID: 123, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/company/getCompanyList",
			Description: "获取company表列表", ApiGroup: "company", Method: "GET"},
		{Model: gorm.Model{ID: 124, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/region/createRegion",
			Description: "新增region表", ApiGroup: "region", Method: "POST"},
		{Model: gorm.Model{ID: 125, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/region/deleteRegion",
			Description: "删除region表", ApiGroup: "region", Method: "DELETE"},
		{Model: gorm.Model{ID: 126, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/region/deleteRegionByIds",
			Description: "批量删除region表", ApiGroup: "region", Method: "DELETE"},
		{Model: gorm.Model{ID: 127, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/region/updateRegion",
			Description: "更新region表", ApiGroup: "region", Method: "PUT"},
		{Model: gorm.Model{ID: 128, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/region/findRegion",
			Description: "根据ID获取region表", ApiGroup: "region", Method: "GET"},
		{Model: gorm.Model{ID: 129, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/region/getRegionList",
			Description: "获取region表列表", ApiGroup: "region", Method: "GET"},
		{Model: gorm.Model{ID: 130, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/address/createAddress",
			Description: "新增address表", ApiGroup: "address", Method: "POST"},
		{Model: gorm.Model{ID: 131, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/address/deleteAddress",
			Description: "删除address表", ApiGroup: "address", Method: "DELETE"},
		{Model: gorm.Model{ID: 132, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/address/deleteAddressByIds",
			Description: "批量删除address表", ApiGroup: "address", Method: "DELETE"},
		{Model: gorm.Model{ID: 133, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/address/updateAddress",
			Description: "更新address表", ApiGroup: "address", Method: "PUT"},
		{Model: gorm.Model{ID: 134, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/address/findAddress",
			Description: "根据ID获取address表", ApiGroup: "address", Method: "GET"},
		{Model: gorm.Model{ID: 135, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/address/getAddressList",
			Description: "获取address表列表", ApiGroup: "address", Method: "GET"},
		{Model: gorm.Model{ID: 136, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/lift/createLift", Description: "新增lift表",
			ApiGroup: "lift", Method: "POST"},
		{Model: gorm.Model{ID: 137, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/lift/deleteLift", Description: "删除lift表",
			ApiGroup: "lift", Method: "DELETE"},
		{Model: gorm.Model{ID: 138, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/lift/deleteLiftByIds",
			Description: "批量删除lift表", ApiGroup: "lift", Method: "DELETE"},
		{Model: gorm.Model{ID: 139, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/lift/updateLift", Description: "更新lift表",
			ApiGroup: "lift", Method: "PUT"},
		{Model: gorm.Model{ID: 140, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/lift/findLift",
			Description: "根据ID获取lift表", ApiGroup: "lift", Method: "GET"},
		{Model: gorm.Model{ID: 141, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/lift/getLiftList",
			Description: "获取lift表列表", ApiGroup: "lift", Method: "GET"},
		{Model: gorm.Model{ID: 142, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/liftModel/createLiftModel",
			Description: "新增liftModel表", ApiGroup: "liftModel", Method: "POST"},
		{Model: gorm.Model{ID: 143, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/liftModel/deleteLiftModel",
			Description: "删除liftModel表", ApiGroup: "liftModel", Method: "DELETE"},
		{Model: gorm.Model{ID: 144, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftModel/deleteLiftModelByIds", Description: "批量删除liftModel表",
			ApiGroup: "liftModel", Method: "DELETE"},
		{Model: gorm.Model{ID: 145, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/liftModel/updateLiftModel",
			Description: "更新liftModel表", ApiGroup: "liftModel", Method: "PUT"},
		{Model: gorm.Model{ID: 146, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/liftModel/findLiftModel",
			Description: "根据ID获取liftModel表", ApiGroup: "liftModel", Method: "GET"},
		{Model: gorm.Model{ID: 147, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/liftModel/getLiftModelList",
			Description: "获取liftModel表列表", ApiGroup: "liftModel", Method: "GET"},
		{Model: gorm.Model{ID: 148, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/userAdmin/createUserAdmin", Description: "新增userAdmin表", ApiGroup: "userAdmin", Method: "POST"},
		{Model: gorm.Model{ID: 149, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/userAdmin/deleteUserAdmin", Description: "删除userAdmin表", ApiGroup: "userAdmin", Method: "DELETE"},
		{Model: gorm.Model{ID: 150, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/userAdmin/deleteUserAdminByIds",
			Description: "批量删除userAdmin表", ApiGroup: "userAdmin", Method: "DELETE"},
		{Model: gorm.Model{ID: 151, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/userAdmin/updateUserAdmin", Description: "更新userAdmin表", ApiGroup: "userAdmin", Method: "PUT"},
		{Model: gorm.Model{ID: 152, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/userAdmin/findUserAdmin", Description: "根据ID获取userAdmin表", ApiGroup: "userAdmin", Method: "GET"},
		{Model: gorm.Model{ID: 153, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/userAdmin/getUserAdminList", Description: "获取userAdmin表列表", ApiGroup: "userAdmin", Method: "GET"},
		{Model: gorm.Model{ID: 154, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userAdmin/createUserAdmin",
			Description: "新增userAdmin表", ApiGroup: "userAdmin", Method: "POST"},
		{Model: gorm.Model{ID: 155, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userAdmin/deleteUserAdmin",
			Description: "删除userAdmin表", ApiGroup: "userAdmin", Method: "DELETE"},
		{Model: gorm.Model{ID: 156, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userAdmin/deleteUserAdminByIds", Description: "批量删除userAdmin表",
			ApiGroup: "userAdmin", Method: "DELETE"},
		{Model: gorm.Model{ID: 157, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userAdmin/updateUserAdmin",
			Description: "更新userAdmin表", ApiGroup: "userAdmin", Method: "PUT"},
		{Model: gorm.Model{ID: 158, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userAdmin/findUserAdmin",
			Description: "根据ID获取userAdmin表", ApiGroup: "userAdmin", Method: "GET"},
		{Model: gorm.Model{ID: 159, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userAdmin/getUserAdminList",
			Description: "获取userAdmin表列表", ApiGroup: "userAdmin", Method: "GET"},
		{Model: gorm.Model{ID: 160, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceData/createAdDeviceData", Description: "新增adDeviceData表", ApiGroup: "adDeviceData", Method: "POST"},
		{Model: gorm.Model{ID: 161, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceData/deleteAdDeviceData", Description: "删除adDeviceData表", ApiGroup: "adDeviceData", Method: "DELETE"},
		{Model: gorm.Model{ID: 162, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceData/deleteAdDeviceDataByIds", Description: "批量删除adDeviceData表", ApiGroup: "adDeviceData",
			Method: "DELETE"},
		{Model: gorm.Model{ID: 163, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceData/updateAdDeviceData", Description: "更新adDeviceData表", ApiGroup: "adDeviceData", Method: "PUT"},
		{Model: gorm.Model{ID: 164, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceData/findAdDeviceData", Description: "根据ID获取adDeviceData表", ApiGroup: "adDeviceData", Method: "GET"},
		{Model: gorm.Model{ID: 165, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceData/getAdDeviceDataList", Description: "获取adDeviceData表列表", ApiGroup: "adDeviceData",
			Method: "GET"},
		{Model: gorm.Model{ID: 166, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceConfig/createAdDeviceConfig", Description: "新增adDeviceConfig表",
			ApiGroup: "adDeviceConfig", Method: "POST"},
		{Model: gorm.Model{ID: 167, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceConfig/deleteAdDeviceConfig", Description: "删除adDeviceConfig表", ApiGroup: "adDeviceConfig",
			Method: "DELETE"},
		{Model: gorm.Model{ID: 168, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceConfig/deleteAdDeviceConfigByIds", Description: "批量删除adDeviceConfig表",
			ApiGroup: "adDeviceConfig", Method: "DELETE"},
		{Model: gorm.Model{ID: 169, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceConfig/updateAdDeviceConfig", Description: "更新adDeviceConfig表",
			ApiGroup: "adDeviceConfig", Method: "PUT"},
		{Model: gorm.Model{ID: 170, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceConfig/findAdDeviceConfig", Description: "根据ID获取adDeviceConfig表",
			ApiGroup: "adDeviceConfig", Method: "GET"},
		{Model: gorm.Model{ID: 171, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceConfig/getAdDeviceConfigList", Description: "获取adDeviceConfig表列表",
			ApiGroup: "adDeviceConfig", Method: "GET"},
		{Model: gorm.Model{ID: 172, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceEvent/createAdDeviceEvent", Description: "新增adDeviceEvent表", ApiGroup: "adDeviceEvent",
			Method: "POST"},
		{Model: gorm.Model{ID: 173, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceEvent/deleteAdDeviceEvent", Description: "删除adDeviceEvent表", ApiGroup: "adDeviceEvent",
			Method: "DELETE"},
		{Model: gorm.Model{ID: 174, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceEvent/deleteAdDeviceEventByIds", Description: "批量删除adDeviceEvent表",
			ApiGroup: "adDeviceEvent", Method: "DELETE"},
		{Model: gorm.Model{ID: 175, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceEvent/updateAdDeviceEvent", Description: "更新adDeviceEvent表", ApiGroup: "adDeviceEvent",
			Method: "PUT"},
		{Model: gorm.Model{ID: 176, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceEvent/findAdDeviceEvent", Description: "根据ID获取adDeviceEvent表", ApiGroup: "adDeviceEvent",
			Method: "GET"},
		{Model: gorm.Model{ID: 177, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDeviceEvent/getAdDeviceEventList", Description: "获取adDeviceEvent表列表", ApiGroup: "adDeviceEvent",
			Method: "GET"},
		{Model: gorm.Model{ID: 178, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/adDevice/createAdDevice",
			Description: "新增adDevice表", ApiGroup: "adDevice", Method: "POST"},
		{Model: gorm.Model{ID: 179, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/adDevice/deleteAdDevice",
			Description: "删除adDevice表", ApiGroup: "adDevice", Method: "DELETE"},
		{Model: gorm.Model{ID: 180, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/adDevice/deleteAdDeviceByIds", Description: "批量删除adDevice表", ApiGroup: "adDevice", Method: "DELETE"},
		{Model: gorm.Model{ID: 181, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/adDevice/updateAdDevice",
			Description: "更新adDevice表", ApiGroup: "adDevice", Method: "PUT"},
		{Model: gorm.Model{ID: 182, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/adDevice/findAdDevice",
			Description: "根据ID获取adDevice表", ApiGroup: "adDevice", Method: "GET"},
		{Model: gorm.Model{ID: 183, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/adDevice/getAdDeviceList",
			Description: "获取adDevice表列表", ApiGroup: "adDevice", Method: "GET"},
		{Model: gorm.Model{ID: 184, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftChange/createLiftChange", Description: "新增liftChange表", ApiGroup: "liftChange", Method: "POST"},
		{Model: gorm.Model{ID: 185, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftChange/deleteLiftChange", Description: "删除liftChange表", ApiGroup: "liftChange", Method: "DELETE"},
		{Model: gorm.Model{ID: 186, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftChange/deleteLiftChangeByIds", Description: "批量删除liftChange表", ApiGroup: "liftChange", Method: "DELETE"},
		{Model: gorm.Model{ID: 187, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftChange/updateLiftChange", Description: "更新liftChange表", ApiGroup: "liftChange", Method: "PUT"},
		{Model: gorm.Model{ID: 188, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/liftChange/findLiftChange",
			Description: "根据ID获取liftChange表", ApiGroup: "liftChange", Method: "GET"},
		{Model: gorm.Model{ID: 189, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftChange/getLiftChangeList", Description: "获取liftChange表列表", ApiGroup: "liftChange", Method: "GET"},
		{Model: gorm.Model{ID: 190, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftRecord/createLiftRecord", Description: "新增liftRecord表", ApiGroup: "liftRecord", Method: "POST"},
		{Model: gorm.Model{ID: 191, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftRecord/deleteLiftRecord", Description: "删除liftRecord表", ApiGroup: "liftRecord", Method: "DELETE"},
		{Model: gorm.Model{ID: 192, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftRecord/deleteLiftRecordByIds", Description: "批量删除liftRecord表", ApiGroup: "liftRecord", Method: "DELETE"},
		{Model: gorm.Model{ID: 193, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftRecord/updateLiftRecord", Description: "更新liftRecord表", ApiGroup: "liftRecord", Method: "PUT"},
		{Model: gorm.Model{ID: 194, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/liftRecord/findLiftRecord",
			Description: "根据ID获取liftRecord表", ApiGroup: "liftRecord", Method: "GET"},
		{Model: gorm.Model{ID: 195, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftRecord/getLiftRecordList", Description: "获取liftRecord表列表", ApiGroup: "liftRecord", Method: "GET"},
		{Model: gorm.Model{ID: 196, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftTrouble/createLiftTrouble", Description: "新增liftTrouble表", ApiGroup: "liftTrouble", Method: "POST"},
		{Model: gorm.Model{ID: 197, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftTrouble/deleteLiftTrouble", Description: "删除liftTrouble表", ApiGroup: "liftTrouble", Method: "DELETE"},
		{Model: gorm.Model{ID: 198, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftTrouble/deleteLiftTroubleByIds", Description: "批量删除liftTrouble表", ApiGroup: "liftTrouble", Method: "DELETE"},
		{Model: gorm.Model{ID: 199, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftTrouble/updateLiftTrouble", Description: "更新liftTrouble表", ApiGroup: "liftTrouble", Method: "PUT"},
		{Model: gorm.Model{ID: 200, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftTrouble/findLiftTrouble", Description: "根据ID获取liftTrouble表", ApiGroup: "liftTrouble", Method: "GET"},
		{Model: gorm.Model{ID: 201, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftTrouble/getLiftTroubleList", Description: "获取liftTrouble表列表", ApiGroup: "liftTrouble", Method: "GET"},
		{Model: gorm.Model{ID: 202, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/findUser",
			Description: "find a user by id", ApiGroup: "user", Method: "GET"},
		{Model: gorm.Model{ID: 203, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/deleteUserList",
			Description: "del user list", ApiGroup: "user", Method: "DELETE"},
		{Model: gorm.Model{ID: 204, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/createUser",
			Description: "create user by admin", ApiGroup: "user", Method: "POST"},
		{Model: gorm.Model{ID: 205, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/createUserLift",
			Description: "新增userLift表", ApiGroup: "userLift", Method: "POST"},
		{Model: gorm.Model{ID: 206, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/deleteUserLift",
			Description: "删除userLift表", ApiGroup: "userLift", Method: "DELETE"},
		{Model: gorm.Model{ID: 207, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/userLift/deleteUserLiftByIds", Description: "批量删除userLift表", ApiGroup: "userLift", Method: "DELETE"},
		{Model: gorm.Model{ID: 208, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/updateUserLift",
			Description: "更新userLift表", ApiGroup: "userLift", Method: "PUT"},
		{Model: gorm.Model{ID: 209, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/findUserLift",
			Description: "根据ID获取userLift表", ApiGroup: "userLift", Method: "GET"},
		{Model: gorm.Model{ID: 210, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/getUserLiftList",
			Description: "获取userLift表列表", ApiGroup: "userLift", Method: "GET"},
		{Model: gorm.Model{ID: 211, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/liftRecord/fillLiftRecord",
			Description: "lift record step2", ApiGroup: "liftRecord", Method: "POST"},
		{Model: gorm.Model{ID: 212, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/liftRecord/reviewLiftRecord", Description: "lift record step3", ApiGroup: "liftRecord", Method: "POST"},
		{Model: gorm.Model{ID: 213, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/message/createMessage",
			Description: "新增message表", ApiGroup: "message", Method: "POST"},
		{Model: gorm.Model{ID: 214, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/message/deleteMessage",
			Description: "删除message表", ApiGroup: "message", Method: "DELETE"},
		{Model: gorm.Model{ID: 215, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/message/deleteMessageByIds",
			Description: "批量删除message表", ApiGroup: "message", Method: "DELETE"},
		{Model: gorm.Model{ID: 216, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/message/updateMessage",
			Description: "更新message表", ApiGroup: "message", Method: "PUT"},
		{Model: gorm.Model{ID: 217, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/message/findMessage",
			Description: "根据ID获取message表", ApiGroup: "message", Method: "GET"},
		{Model: gorm.Model{ID: 218, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/message/getMessageList",
			Description: "获取message表列表", ApiGroup: "message", Method: "GET"},
		{Model: gorm.Model{ID: 219, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthSystem/createHealthSystem", Description: "新增healthSystem表", ApiGroup: "healthSystem",
			Method: "POST"},
		{Model: gorm.Model{ID: 220, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthSystem/deleteHealthSystem", Description: "删除healthSystem表", ApiGroup: "healthSystem",
			Method: "DELETE"},
		{Model: gorm.Model{ID: 221, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthSystem/deleteHealthSystemByIds", Description: "批量删除healthSystem表", ApiGroup: "healthSystem",
			Method: "DELETE"},
		{Model: gorm.Model{ID: 222, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthSystem/updateHealthSystem", Description: "更新healthSystem表", ApiGroup: "healthSystem", Method: "PUT"},
		{Model: gorm.Model{ID: 223, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthSystem/findHealthSystem", Description: "根据ID获取healthSystem表", ApiGroup: "healthSystem",
			Method: "GET"},
		{Model: gorm.Model{ID: 224, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthSystem/getHealthSystemList", Description: "获取healthSystem表列表", ApiGroup: "healthSystem",
			Method: "GET"},
		{Model: gorm.Model{ID: 225, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthChange/createHealthChange", Description: "新增healthChange表", ApiGroup: "healthChange",
			Method: "POST"},
		{Model: gorm.Model{ID: 226, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthChange/deleteHealthChange", Description: "删除healthChange表", ApiGroup: "healthChange",
			Method: "DELETE"},
		{Model: gorm.Model{ID: 227, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthChange/deleteHealthChangeByIds", Description: "批量删除healthChange表", ApiGroup: "healthChange", Method: "DELETE"},
		{Model: gorm.Model{ID: 228, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthChange/updateHealthChange", Description: "更新healthChange表", ApiGroup: "healthChange", Method: "PUT"},
		{Model: gorm.Model{ID: 229, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthChange/findHealthChange", Description: "根据ID获取healthChange表", ApiGroup: "healthChange",
			Method: "GET"},
		{Model: gorm.Model{ID: 230, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path: "/healthChange/getHealthChangeList", Description: "获取healthChange表列表", ApiGroup: "healthChange",
			Method: "GET"},
		{Model: gorm.Model{ID: 231, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/createUserLift",
			Description: "新增userLift表", ApiGroup: "userLift", Method: "POST"},
		{Model: gorm.Model{ID: 232, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/deleteUserLift",
			Description: "删除userLift表", ApiGroup: "userLift", Method: "DELETE"},
		{Model: gorm.Model{ID: 233, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/deleteUserLiftByIds",
			Description: "批量删除userLift表", ApiGroup: "userLift", Method: "DELETE"},
		{Model: gorm.Model{ID: 234, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/updateUserLift",
			Description: "更新userLift表", ApiGroup: "userLift", Method: "PUT"},
		{Model: gorm.Model{ID: 235, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/findUserLift",
			Description: "根据ID获取userLift表", ApiGroup: "userLift", Method: "GET"},
		{Model: gorm.Model{ID: 236, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/userLift/getUserLiftList",
			Description: "获取userLift表列表", ApiGroup: "userLift", Method: "GET"},
		{Model: gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getAllApis", Description: "获取所有api", ApiGroup: "api", Method: "POST"},
	}
	if err := tx.Create(&insert).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InitUserLift() (err error) {
	// first clear, hard mode
	global.GVA_DB.Unscoped().Where("ID > 0").Delete(&model.UserLift{})
	// then add new
	tx := global.GVA_DB.Begin()
	insert := []model.UserLift{
		{Model:gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 18 , LiftId: 1,
			CategoryId: uint(enum.UserLiftMaintain)},
		{Model:gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 19 , LiftId: 1,
			CategoryId: uint(enum.UserLiftMaintain)},
		{Model:gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 20 , LiftId: 2,
			CategoryId: uint(enum.UserLiftMaintain)},
		{Model:gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 21 , LiftId: 2,
			CategoryId: uint(enum.UserLiftMaintain)},

		{Model:gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 14 , LiftId: 1,
			CategoryId: uint(enum.UserLiftInstall)},
		{Model:gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 15 , LiftId: 1,
			CategoryId: uint(enum.UserLiftInstall)},
		{Model:gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 16 , LiftId: 2,
			CategoryId: uint(enum.UserLiftInstall)},
		{Model:gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 17 , LiftId: 2,
			CategoryId: uint(enum.UserLiftInstall)},

		{Model:gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 10 , LiftId: 1,
			CategoryId: uint(enum.UserLiftCheck)},
		{Model:gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 11 , LiftId: 1,
			CategoryId: uint(enum.UserLiftCheck)},
		{Model:gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 12 , LiftId: 2,
			CategoryId: uint(enum.UserLiftCheck)},
		{Model:gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 13 , LiftId: 2,
			CategoryId: uint(enum.UserLiftCheck)},

		{Model:gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 6 , LiftId: 1,
			CategoryId: uint(enum.UserLiftManage)},
		{Model:gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 7 , LiftId: 1,
			CategoryId: uint(enum.UserLiftManage)},
		{Model:gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 8 , LiftId: 2,
			CategoryId: uint(enum.UserLiftManage)},
		{Model:gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UserId: 9 , LiftId: 2,
			CategoryId: uint(enum.UserLiftManage)},
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
	err = InitCompany()
	err = InitUser()
	err = InitSysApi()
	err = InitDevice()
	err = InitDeviceOwner()
	err = InitDeviceConfig()
	err = InitDeviceConfigRelation()
	err = InitLift()
	err = InitLiftModel()
	err = InitHealthSystem()
	err = InitHealthChange()
	err = InitSysAuthorityMenus()
	err = InitCasbinModel()
	err = InitExaCustomer()
	err = InitSysBaseMenus()
	err = InitSysDictionary()
	err = InitSysDictionaryDetail()
	err = InitExaFileUploadAndDownload()
	err = InitLiftRecord()
	err = InitLiftTrouble()
	err = InitUserLift()
	if err != nil {
		global.GVA_LOG.Error("initialize data failed", err)
	} else {
		global.GVA_LOG.Debug("initialize data success")
	}
}
