package utils

import (
    "net"
    "fmt"
    "time"
    "strconv"
    "crypto/md5"
    "encoding/hex"
    "monitor/cmd/protocols"
)

func UnixTime() int64 {
    return time.Now().Unix()
}

func ParseOutPut(Conn *net.UnixConn) {
    Buffer := make([]byte, protocols.READ_LENGTH)
    Len, err := Conn.Read(Buffer)
    if err == nil {
        fmt.Println(string(Buffer[0:Len]))
    }
    Conn.Close()
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