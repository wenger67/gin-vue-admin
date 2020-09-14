package enum

type SubjectType int32

const (
	_ SubjectType = iota
	// 电梯使用场景
	SubjectLiftType
	// 电梯记录
	SubjectLiftRecordType
	// 电梯故障来源
	SubjectLiftTroubleSourceType
	// 电梯故障解除方式
	SubjectLiftTroubleSolvedType
	// 电梯故障原因
	SubjectLiftTroubleReasonType
	// 电梯厢门状态
	SubjectLiftDoorStatusType
	// 智能设备型号
	SubjectDeviceModalType = 100
	// 智能设备状态
	SubjectDeviceStatusType = 101
	// 智能设备事件
	SubjectDeviceEventType = 102
	// 公司状态
	SubjectCompanyStatusType = 200
	// 公司类别
	SubjectCompanyType = 201
	// 用户电梯关系
	SubjectUserLiftShipType = 300
	// 用户类型
	SubjectUserType = 301
	// 地址标签
	SubjectAddressTagType = 400
	// 消息类型
	SubjectMessageType = 500
	// 消息主体
	SubjectMessageFromTargetType = 501
	// 大数据维度
	SubjectHealthDimensionType = 600
)
