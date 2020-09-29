package enum

type RecordProgress uint

const (
	_ RecordProgress = iota
	RecordCreated
	RecordStarted
	RecordEnded
	RecordReviewed
)
