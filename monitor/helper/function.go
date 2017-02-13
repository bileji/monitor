package helper

import (
    "time"
    "strconv"
    "crypto/md5"
    "encoding/hex"
)

func UnixTime() int64 {
    return time.Now().Unix()
}

func Md5(Str string) string {
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(Str))
    cipherStr := md5Ctx.Sum(nil)
    return hex.EncodeToString(cipherStr)
}

func RandStr() string {
    UnixTime := int(UnixTime())
    return Md5(strconv.Itoa(UnixTime))
}