package common

import (
    "net"
    "fmt"
    "encoding/json"
    "monitor/command/protocol"
    "monitor/monitor"
    "monitor/monitor/header"
)

type Dispatcher struct {
    Message header.UnixMsg
    Conn    *net.UnixConn
}

func (dis *Dispatcher) Res(Code int, Msg string) (int, error) {
    Res := protocol.Response{
        Code: Code,
        Body: []byte(Msg),
    }
    Byte, _ := json.Marshal(Res)
    return dis.Conn.Write(Byte)
}

func Run(Msg header.UnixMsg, Conn *net.UnixConn, Monitor *monitor.Monitor) {
    Dis := &Dispatcher{
        Conn: Conn,
        Message: Msg,
    }
    
    switch Dis.Message.Command {
    case CMD_ROLE:
    
    case CMD_SERVER_INIT:
        err := Monitor.SInit(Dis.Message.Body);
        if err != nil {
            Dis.Res(-1, fmt.Sprintf("%v", err))
            return
        }
        Dis.Res(-1, "success")
        return
    case CMD_SERVER_TOKEN:
    
    case CMD_JOIN:
    
    default:
        
    }
}
