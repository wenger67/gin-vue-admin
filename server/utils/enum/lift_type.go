package enum

type LiftType int32

const (
	_ LiftType = iota
	HousePassengerLift
	HouseFreightLift
	MarketPassengerLift
	MarketFreightLift
	HospitalPassengerLift
	HospitalFreightLift
	OfficePassengerLift
	OfficeFreightLift
	GovPassengerLift
	GovFreightLift
)
