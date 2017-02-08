package utils

import (
    "time"
)

func UnixTime() int64 {
    return int64(time.Now().Second())
}