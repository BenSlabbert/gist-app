package util

import (
	"strconv"
	"time"
)

func UnixTimestampStringToTime(s string) (time.Time, error) {
	unixTimestamp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(unixTimestamp, 0), nil
}

func UnixTimestampIntToTime(i int) time.Time {
	return time.Unix(int64(i), 0)
}
