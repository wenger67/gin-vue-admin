package response

type LiftCountByAddress struct {
	Address     int    `json:"address"`
	AddressName string `json:"addressName"`
	Total       int    `json:"total"`
}

type LiftTroubleCountByProgress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type LiftRecordCountByType struct {
	Type  int `json:"type"`
	Total int `json:"total"`
}
