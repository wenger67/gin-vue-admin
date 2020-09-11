package enum

type SubjectType int32

const (
	_ SubjectType = iota
	SubjectLiftType
	SubjectLiftRecordType
	SubjectLiftTroubleSourceType
	SubjectLiftTroubleSolvedType
	SubjectLiftTroubleReasonType
	SubjectLiftDoorStatusType

	SubjectDeviceModalType = 100
	SubjectDeviceStatusType = 101
	SubjectDeviceEventType = 102

	SubjectCompanyStatusType = 200
	SubjectCompanyType = 201

	SubjectUserLiftShipType = 300
	SubjectUserType = 301

	SubjectAddressTagType = 400

	SubjectMessageType = 500
	SubjectMessageFromTargetType = 501

	SubjectHealthDimensionType = 600
)
