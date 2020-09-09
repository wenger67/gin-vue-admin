package enum

type HealthDimensionType int32

const (
	_ HealthDimensionType = iota
	HealthTimeDimension
	HealthMaintainDimension
	HealthHumanDimension
	HealthInnerDimension
	HealthSensorDimension
)

