package enum

type HealthDimensionType int

const (
	_ HealthDimensionType = iota
	HealthTimeDimension
	HealthMaintainDimension
	HealthHumanDimension
	HealthInnerDimension
	HealthSensorDimension
)

