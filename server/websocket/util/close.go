package util

import (
	"gin-vue-admin/global"
	"io"
)

func SafeClose(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		global.GVA_LOG.Warning("close closer failed")
	}
}
