package enum

type UserLiftShipType int32

const (
	_ UserLiftShipType = iota
	UserUseLift
	UserMaintainLift
	UserManageLift
	UserCheckLift

)

