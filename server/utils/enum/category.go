package enum

type Category int

const (
	_ Category = iota
	// SubjectLiftType
	// 住宅客梯
	HousePassengerLift
	// 住宅货梯
	HouseFreightLift
	// 商场客梯
	MarketPassengerLift
	// 商场货梯
	MarketFreightLift
	// 医院客梯
	HospitalPassengerLift
	// 医院货梯
	HospitalFreightLift
	// 写字楼客梯
	OfficePassengerLift
	// 写字楼货梯
	OfficeFreightLift
	// 政府客梯
	GovPassengerLift
	// 政府货梯
	GovFreightLift

	// SubjectLiftRecordType
	// 维保记录
	LiftMaintainRecord = 100
	// 年检记录
	LiftUniCheckRecord = 101
	// 投诉建议
	LiftComplainRecord = 102

	// SubjectLiftTroubleSourceType
	LiftUserReportTrouble   = 200
	LiftAIDetectTrouble     = 201
	LiftMaintainTrouble     = 202
	LiftCheckTrouble        = 203
	LiftOwnerTrouble        = 204
	LiftUrgentButtonTrouble = 205

	// SubjectLiftTroubleSolvedType
	TroubleSolvedByAuto = 300
	TroubleSolvedByReplaceCell = 301
	TroubleSolvedByNone = 302

	// SubjectLiftTroubleReasonType
	TroubleReasonAsHuman = 400
	TroubleReasonAsMaintain = 401
	TroubleReasonAsCheck = 402
	TroubleReasonAsAge = 403
	TroubleReasonAsElectric = 404

	// SubjectLiftDoorStatusType
	DoorCloseFailed = 500
	DoorOpenFailed = 501
	DoorCloseSuccess = 502
	DoorOpenSuccess = 503
	DoorOpenAtUnderfloor = 504

	// SubjectDeviceModalType
	DeviceModalV1 = 600
	DeviceModalV2 = 601
	DeviceModalV3 = 602

	// SubjectDeviceStatusType
	DeviceUnderTest = 700
	DeviceUnderCheck = 701
	DeviceNormal = 702
	DeviceAbnormal = 703

	// SubjectDeviceEventType
	DeviceOnlineEvent = 800
	DeviceOfflineEvent = 801
	DevicePersonDetectEvent = 802
	DeviceMotorDetectEvent = 803

	// SubjectCompanyStatusType
	CompanyNormal = 900
	CompanyIrregular = 901

	// SubjectCompanyType
	CompanyMaintain = 1000
	CompanyCheck = 1001
	CompanyOwner = 1002
	CompanyInstall = 1003
	CompanySupervise = 1004
	CompanyPlatform = 1005

	// SubjectUserLiftShipType
	UserLiftUse = 1100
	UserLiftMaintain = 1101
	UserLiftManage = 1102
	UserLiftCheck = 1103
	UserLiftInstall = 1104
	UserLiftSupervise = 1105

	// SubjectUserType
	UserCommon = 1200
	UserPresent = 1201
	UserCompanyManager = 1202
	UserCompanyEmployee = 1203
	UserAdmin = 1204
	UserSuperAdmin = 1205

	// SubjectAddressTagType
	AddressTagBaby = 1300
	AddressTagJunior = 1301
	AddressTagSenior = 1302
	AddressTagCollege = 1303
	AddressTagWorker = 1304
	AddressTagDoctor = 1305
	AddressTagTeacher = 1306
	AddressTagFarmer = 1307
	AddressTagDesigner = 1308
	AddressTagEngineer = 1309
	AddressTagDriver = 1310
	AddressTagCook = 1311
	AddressTagBusinessman = 1312

	// SubjectMessageType
	MessageUrgent = 1400
	MessageLiftChange = 1401
	MessageAssign = 1402
	MessageDeviceEvent = 1403
	MessageNewLift = 1404
	MessageNewTrouble = 1405
	MessageNewRecord = 1406
	MessageRecordDone = 1407
	MessageTroubleFixed = 1408
	MessageNewMaintainWorker = 1409
	MessageNewCheckWorker = 1410
	MessageNewInstallWorker = 1411
	MessageNewManageWorker = 1412
	MessageNewSuperviseWorker = 1413

	// SubjectMessageFromTargetType
	MessageSubjectLift = 1500
	MessageSubjectLiftRecord = 1501
	MessageSubjectLiftTrouble = 1502
	MessageSubjectDeviceEvent = 1503
	MessageSubjectUser = 1504
	MessageSubjectCompany = 1505

	// SubjectHealthDimensionType
	HealthTimeDimension = 1600
	HealthMaintainDimension = 1601
	HealthHumanDimension = 1602
	HealthInnerDimension = 1603
	HealthSensorDimension = 1604
)

