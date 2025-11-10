package Base64

import "encoding/base64"

// Base64Encode 将原字符串转码成Base64格式
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
