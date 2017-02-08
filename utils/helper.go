package utils

import (
    "net"
    "fmt"
    "time"
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