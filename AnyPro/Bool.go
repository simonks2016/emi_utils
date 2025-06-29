package AnyPro

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Any2Bool(data interface{}) (bool, error) {
	if data == nil {
		return false, errors.New("invalid data")
	}

	switch v := data.(type) {
	case bool:
		return v, nil
	case int:
		return v == 1, nil
	case int64:
		return v == 1, nil
	case int32:
		return v == 1, nil
	case int16:
		return v == 1, nil
	case int8:
		return v == 1, nil
	case float64:
		return v == 1.0, nil
	case float32:
		return float64(v) == 1.0, nil
	case string:
		s := strings.TrimSpace(strings.ToLower(v))
		if s == "" {
			return false, nil
		}
		if s == "1" || s == "true" || s == "yes" {
			return true, nil
		}
		if s == "0" || s == "false" || s == "no" {
			return false, nil
		}
		if num, err := strconv.ParseFloat(s, 64); err == nil {
			return num == 1.0, nil
		}
		return false, fmt.Errorf("cannot convert string to bool: %s", v)
	default:
		return false, fmt.Errorf("unsupported type: %T", data)
	}
}

func Any2NoErrBool(data interface{}) bool {
	r, _ := Any2Bool(data)
	return r
}
