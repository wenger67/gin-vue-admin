package enum

type TroubleProgress uint

const (
	_ TroubleProgress = iota
	TroubleCreated
	TroubleResponded
	TroubleScened
	TroubleFixed
	TroubleFeedback
	TroubleReviewed
)
