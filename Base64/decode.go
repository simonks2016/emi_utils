package Base64

import "encoding/base64"

// Base64Decode 将 Base64 字符串解码为普通字符串
func Base64DecodeWithError(encoded string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

func Base64Decode(encoded string) string {

	if b, err := Base64DecodeWithError(encoded); err != nil {
		return encoded
	} else {
		return b
	}
}
