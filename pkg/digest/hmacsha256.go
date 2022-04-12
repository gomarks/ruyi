package digest

import (
	"encoding/base64"
	"fmt"
	"github.com/bingoohuang/gg/pkg/mapp"
	"github.com/chmike/hmacsha256"
	"github.com/codingsince1985/checksum"
	"log"
	"strings"
)

// CheckSumByHmacSHA256 基础报文做 HmacSHA256 校验码
func CheckSumByHmacSHA256(appId, deviceId, securityKey, version, filePath string, others map[string]string) (string, error) {
	allParams := mapp.Clone(others)
	if filePath != "" {
		md5, err := checksum.MD5sum(filePath)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		allParams["md5"] = md5
	}
	allParams["appId"] = appId
	allParams["deviceId"] = deviceId
	allParams["signAlgo"] = "HmacSHA256"
	allParams["version"] = version

	for k, v := range allParams {
		if v == "" {
			delete(allParams, k)
		}
	}

	sortedParams := make([]string, 0)
	for _, k := range mapp.KeysSorted(allParams) {
		sortedParams = append(sortedParams, fmt.Sprintf("%s=%s", k, allParams[k]))
	}

	toDigest := strings.Join(sortedParams, "&")

	h := hmacsha256.Obj{}
	h.Init([]byte(securityKey))
	digest := h.Digest(nil, []byte(toDigest))
	return toDigest + fmt.Sprintf("&signature=%s", base64.StdEncoding.EncodeToString(digest)), nil
}
