package db_accessor

import (
	"strconv"
	"time"
)

func TimeToString(a time.Time) string {
	var s00 = func(x int) string {
		var result = strconv.Itoa(x)
		for len(result) < 2 {
			result = "0" + result
		}
		return result
	}
	return s00(a.Hour()) + ":" + s00(a.Minute()) + ":" + s00(a.Second())
}
