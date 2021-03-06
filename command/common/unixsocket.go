package common

import (
    "os"
    "net"
    "fmt"
    "time"
    "encoding/json"
    "monitor/monitor/header"
)

type Response struct {
    Code int
    Body []byte
}

type Socket struct {
    CUnix string
    SUnix string
    Conn  *net.UnixConn
}

func (s *Socket) Close() error {
    if s.Conn != nil {
        return s.Conn.Close()
    }
    return nil
}

func (s *Socket) UnixSocket() error {
    os.Remove(s.CUnix)
    
    lAddr, err := net.ResolveUnixAddr("unix", s.CUnix)
    if err != nil {
        return err
    }
    rAddr, err := net.ResolveUnixAddr("unix", s.SUnix)
    if err != nil {
        return err
    }
    s.Conn, err = net.DialUnix("unix", lAddr, rAddr)
    if err != nil {
        return err
    }
    return nil
}

func (s *Socket)SendMessage(Message header.UnixMsg) {
    if Message.Timestamp == 0 {
        Message.Timestamp = time.Now().Unix()
    }
    MsgBuffer, _ := json.Marshal(Message)
    s.Conn.Write(MsgBuffer)
}

func (s *Socket)EchoReceive() {
    defer s.Conn.Close()
    Buffer := make([]byte, SOCKET_REC_LENGTH)
    if Len, err := s.Conn.Read(Buffer); err == nil {
        Message := Response{}
        json.Unmarshal(Buffer[0:Len], &Message)
        fmt.Println(string(Message.Body))
    }
}