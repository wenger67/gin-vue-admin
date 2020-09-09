package enum

type SubjectType int32

const (
	_ SubjectType = iota
	SubjectLiftType
	SubjectLiftRecordType
	SubjectLiftTroubleSourceType
	SubjectLiftTroubleSolvedType
	SubjectLiftTroubleReasonType
	SubjectLiftTroubleType
	SubjectLiftDoorStatusType

	SubjectDeviceModalType
	SubjectDeviceStatusType
	SubjectDeviceEventType

	SubjectCompanyStatusType
	SubjectCompanyType

	SubjectUserLiftShipType
	SubjectUserType

	SubjectAddressTagType

	SubjectMessageType
	SubjectMessageFromTargetType

	SubjectHealthDimensionType
)
