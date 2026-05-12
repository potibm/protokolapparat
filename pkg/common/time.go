package common

import "time"

// NowUnix returns the current time as a Unix timestamp in seconds.
func NowUnix() int64 {
	return time.Now().Unix()
}
