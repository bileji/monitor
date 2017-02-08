package utils

import (
    "time"
)

func UnixTime() int64 {
    return time.Now().Unix()
}