package digest

import (
	"github.com/prometheus/common/log"
	"testing"
)

func TestSumByHmacSHA256(t *testing.T) {
	appId := "APP_AF100AAE66EA4CE6A4F0B5F20F06"
	deviceId := "DEV_DE4826EBC32746A795F4CF45E33F"
	securityKey := "7GaIDctCkMHtaIskQF6xWnUFcWTe"
	version := "1.0"
	filePath := "/tmp/1.txt"
	others := map[string]string{}
	others["user"] = "userlabel"
	others["group"] = "grouplabel"
	others["expire"] = "360000"
	sha256, _ := CheckSumByHmacSHA256(appId, deviceId, securityKey, version, filePath, others)
	log.Info(sha256)
}
