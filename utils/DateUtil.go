package util

import "time"

const (
	YYYYMMDDHHMMSS = "20060102150405"
	DateOnly       = "2006-01-02"
	TimeOnly       = "15:04:05"
	DateTime       = "2006-01-02 15:04:05"
)

// StrToTime 字符串转日期
func StrToTime(timeStr, formatStr string) (time.Time, error) {
	var cstZone = time.FixedZone("CST", 8*3600)
	return time.ParseInLocation(formatStr, timeStr, cstZone)
}

func TimeToStr(t time.Time, format string) string {
	return t.Format(format)
}
