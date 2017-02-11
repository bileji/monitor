package dispatcher

import (
    "net"
    "encoding/json"
    "monitor/command/common"
    "monitor/command/protocol"
)

type Dispatcher struct {
    Message protocol.SocketMsg
    Conn    *net.UnixConn
}

func (dis *Dispatcher) Res(Code int, Msg string) (string, error) {
    Res := protocol.Response{
        Code: Code,
        Body: []byte(Msg),
    }
    return json.Marshal(Res)
}

func Run(Msg protocol.SocketMsg, Conn *net.UnixConn) {
    Dis := &Dispatcher{
        Conn: Conn,
        Message: Msg,
    }
    
    switch Dis.Message.Command {
    
    case common.CMD_ROLE:
    
    case common.CMD_SERVER_INIT:
    
    case common.CMD_SERVER_TOKEN:
    
    case common.CMD_JOIN:
    
    default:
        
    }
}
