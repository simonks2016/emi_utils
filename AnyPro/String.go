package AnyPro

import "fmt"

func Any2String(a any) string {
	if a == nil {
		return ""
	}
	return fmt.Sprintf("%v", a)
}
