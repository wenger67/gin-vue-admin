package enum

type LiftRecordType int32

const (
	_ LiftRecordType = iota
	LiftMaintainRecord
	LiftUniCheckRecord
	LiftComplainRecord
)

