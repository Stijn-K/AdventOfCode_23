package common

import "strconv"

func StringToInt(in string) int {
	out, _ := strconv.ParseInt(in, 10, 32)
	return int(out)
}

func StringToInt64(in string) int64 {
	out, _ := strconv.ParseInt(in, 10, 64)
	return out
}
