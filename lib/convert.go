package lib

import "strconv"

func FloatToString[T float32 | float64](n T) string {
	switch nt := any(n).(type) {
	case float32:
		return strconv.FormatFloat(float64(nt), 'f', -1, 32)

	case float64:
		return strconv.FormatFloat(nt, 'f', -1, 64)

	default:
		return ""
	}
}
