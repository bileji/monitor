package common

import (
    "net"
    "fmt"
    "encoding/json"
    "monitor/monitor"
    "monitor/monitor/header"
)

type Dispatcher struct {
    Message header.UnixMsg
    Conn    *net.UnixConn
}

func (d *Dispatcher) Res(Code int, Msg string) (int, error) {
    Byte, _ := json.Marshal(Response{
        Code: Code,
        Body: []byte(Msg),
    })
    return d.Conn.Write(Byte)
}

func Run(Msg header.UnixMsg, Conn *net.UnixConn, Monitor *monitor.Monitor) {
    Dis := &Dispatcher{
        Conn: Conn,
        Message: Msg,
    }
    
    switch Dis.Message.Command {
    case CMD_ROLE:
        Dis.Res(SUCCESS, Monitor.Role())
    case CMD_SERVER_INIT:
        err := Monitor.ManagerInit(Dis.Message.Body);
        if err == nil {
            Dis.Res(SUCCESS, "success")
            return
        }
        Dis.Res(FAILURE, fmt.Sprintf("%v", err))
        return
    case CMD_SERVER_TOKEN:
        Token, err := Monitor.ManagerToken()
        fmt.Println("++++++")
        fmt.Println(Token)
        fmt.Println("------")
        if err == nil {
            Dis.Res(SUCCESS, Token)
        }
        Dis.Res(FAILURE, fmt.Sprintf("%v", err))
        return
    case CMD_JOIN:
        err := Monitor.Join(Dis.Message.Body)
        if err == nil {
            Dis.Res(SUCCESS, "success")
        }
        Dis.Res(FAILURE, fmt.Sprintf("%v", err))
        return
    default:
        
    }
}
