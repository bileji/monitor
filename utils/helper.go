package utils

import (
    "net"
    "time"
    "monitor/cmd/protocols"
)

func UnixTime() int64 {
    return time.Now().Unix()
}

func ParseOutPut(Conn *net.UnixConn) {
    Buffer := make([]byte, protocols.READ_LENGTH)
    Conn.Read(Buffer)
}