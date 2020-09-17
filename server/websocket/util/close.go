package util

import (
	"gin-vue-admin/global"
	"io"
)

func SafeClose(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		global.PantaLog.Warning("close closer failed")
	}
}
