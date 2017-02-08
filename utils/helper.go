package utils

import (
    "time"
)

func UnixTime() int {
    return time.Now().Second()
}