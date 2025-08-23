package Is

import (
	"encoding/base64"
	"regexp"
	"strings"
)

func IsLegalString(str string, min, max int) bool {

	if len([]rune(str)) > max || len([]rune(str)) <= min {
		return false
	}

	// 允许字符：中英文、数字、下划线、中划线、感叹号、问号、#、$、&
	pattern := `^[a-zA-Z0-9_\-\!\?\#\$&\p{Han}]+$`
	// 匹配汉字、英文、数字
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	return matched
}

func IsBase64String(str string) bool {
	// 先试用正则来判断
	pattern := "^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$"
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	if !(len(str)%4 == 0 && matched) {
		return false
	}
	// 在使用自身带的解码和转码 在查看他么的值是否相等,如果相等说明是base64 否则不是
	deStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return false
	}
	bs64Str := base64.StdEncoding.EncodeToString(deStr)
	//如果解码后再转码和原来的字符一样说明是base64
	return str == bs64Str
}

func IsLegalPlatform(p string) bool {
	switch strings.ToLower(p) {
	case "ios", "android", "web", "macos", "windows", "liunx", "tvos":
		return true
	default:
		return false
	}
}

func IsStringsArray(str string) bool {

	// 去除前后空格
	str = strings.TrimSpace(str)
	if str == "" {
		return false
	}
	// 定义正则表达式，匹配任意一个分隔符
	// 注意 `&&` 要写在 `&` 之前，否则会先匹配到单个 `&`
	re := regexp.MustCompile(`&&|&|、|；|;|,|，`)
	// 使用正则拆分
	parts := re.Split(str, -1)
	// 统计非空项个数
	count := 0
	for _, part := range parts {
		if strings.TrimSpace(part) != "" {
			count++
		}
	}
	// 如果分割后有超过1项，就认为是“数组”
	return count > 1
}

func IsHttpsURL(s string) bool {

	p := `^(https?:\/\/)([0-9a-z.]+)(:[0-9]+)?([/0-9a-z.]+)?(\?[0-9a-z&=]+)?(#[0-9-a-z]+)?`

	if m, err := regexp.Match(p, []byte(s)); err != nil {
		return false
	} else {
		return m
	}
}

func IsMobilePhone(m string) bool {
	p := `^\+?[0-9]{6,15}$`
	match, err := regexp.MatchString(p, m)
	if err != nil {
		return false
	}
	return match
}

func IsLegalDataId(str string) bool {

	p := `^[a-zA-Z0-9_\-]{1,40}$`

	matched, err := regexp.MatchString(p, str)
	if err != nil {
		return false
	}
	return matched
}
