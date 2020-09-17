package response

import "panta/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
