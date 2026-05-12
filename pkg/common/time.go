package common

import "time"

func NowUnix() int64 {
    return time.Now().Unix()
}