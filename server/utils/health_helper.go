package utils

import "time"

func CalTimeDimensionInitVal(t time.Time) (value int) {
	now := time.Now()
	hourDuration := now.Sub(t) / time.Hour
	dayDuration := hourDuration / 24

	// TODO use algorithm to replace
	if dayDuration < 180 {
		return HealthDimensionInitialValue
	} else if dayDuration < 365 {
		return int(HealthDimensionInitialValue - dayDuration/180)
	} else if dayDuration < 1000 {
		return int(HealthDimensionInitialValue - dayDuration/150)
	} else if dayDuration < 3000 {
		return int(HealthDimensionInitialValue - dayDuration/100)
	} else {
		temp := HealthDimensionInitialValue - dayDuration/80
		if temp > 0 {
			return int(temp)
		} else {
			return 0
		}
	}
}
