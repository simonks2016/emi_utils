package Is

import "time"

func IsMillisecondTimestamp(ts int64) bool {
	// 当前时间
	now := time.Now().Unix()
	// 判断是否明显比当前时间大了三位（毫秒一般比秒多3位）
	return ts > now*100
}
