package dispatcher

import (
    "monitor/command/protocol"
    "net"
)

type Dispatcher struct {
    Message protocol.SocketMsg
    Conn    *net.UnixConn
}

func Run(Msg protocol.SocketMsg, Conn *net.UnixConn) {
    &Dispatcher{
        Conn: Conn,
        Message: Msg,
    }
    // todo
}
