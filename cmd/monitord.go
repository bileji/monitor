package main

import (
    "os"
    "fmt"
    "log"
    "net"
    "runtime"
    "encoding/json"
    "monitor/daemon"
    "monitor/cmd/protocols"
    "github.com/spf13/cobra"
)

type DBAuth struct {
    Host     string `json:"host"`
    Port     int `json:"port"`
    Database string `json:"database"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    RootCmd := &cobra.Command{
        Use: "monitord",
        Short: "monitor daemon",
        // todo ...
        Long: "to do...",
        RunE:func(cmd *cobra.Command, args []string) error {
            
            return nil
        },
    }
    
    if err := RootCmd.Execute(); err != nil {
        // todo
        fmt.Printf("%v", err)
        os.Exit(-1)
    }
    
    Daemon := &daemon.Daemon{
        PidFile: "/var/run/monitord.pid",
        UnixFile: "/var/run/monitord.sock",
        LogFile: "/var/log/monitord.log",
    }
    
    Daemon.Daemon(func(Unix *net.UnixListener) {
        for {
            if Fd, err := Unix.AcceptUnix(); err != nil {
                log.Printf("accect error: %v", err)
            } else {
                for {
                    Buffer := make([]byte, 512)
                    if Len, err := Fd.Read(Buffer); err != nil {
                        log.Printf("accect error: %v", err)
                    } else {
                        var Message protocols.Socket
                        json.Unmarshal(Buffer[0: Len], &Message)
                        // todo 接收到cli信息,然后处理
    
                    }
                }
            }
        }
    })
}