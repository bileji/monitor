package common

import (
    "net"
    "encoding/json"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "monitor/command/protocol"
    "monitor/command/dispatcher"
)

type MonitorFlags struct {
    Config string
    Daemon bool
    Pid    string
    Log    string
    Debug  string
}

type ServerFlags struct {
    
}

type JoinFlags struct {
    Addr  string
    Token string
}

type Flags struct {
    Main   MonitorFlags
    Server ServerFlags
    Join   JoinFlags
}

type Command struct {
    Flags     Flags
    Command   *cobra.Command
    Children  []*Command
    Viper     *viper.Viper
    Configure *Configure
}

func (c *Command) NewChildren() {
    for _, Child := range c.Children {
        c.Command.AddCommand(Child.Command)
    }
}

func (c *Command) Scheduler(Listener *net.UnixListener) {
    defer Listener.Close()
    
    var socket = func(Con *net.UnixConn) {
        for {
            Buffer := make([]byte, SOCKET_REC_LENGTH)
            Len, err := Con.Read(Buffer);
            if err != nil {
                Con.Close()
                return
            }
            var Message protocol.SocketMsg
            json.Unmarshal(Buffer[0: Len], &Message)
            
            dispatcher.Run(Message, Con)
        }
    }
    
    for {
        if UnixConn, err := Listener.AcceptUnix(); err == nil {
            go socket(UnixConn)
        }
    }
}

