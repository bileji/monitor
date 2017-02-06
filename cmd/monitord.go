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

const (
    READ_LENGTH = 2048
)

type DBAuth struct {
    Host     string `json:"host"`
    Port     int `json:"port"`
    Database string `json:"database"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func Scheduler(Con *net.UnixConn) {
    for {
        Buffer := make([]byte, READ_LENGTH)
        Len, err := Con.Read(Buffer);
        if err != nil {
            Con.Close()
            return
        }
        var Message protocols.Socket
        json.Unmarshal(Buffer[0: Len], &Message)
        // todo 接收到cli信息,然后处理
        log.Println(Message)
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    RootCmd := &cobra.Command{
        Use: "monitord",
        Short: "Linux server status monitor daemon",
        Long: "######to do...",
        RunE:func(cmd *cobra.Command, args []string) error {
            // TODO run ...
            Daemon := &daemon.Daemon{
                PidFile: "/var/run/monitord.pid",
                UnixFile: "/var/run/monitord.sock",
                LogFile: "/var/log/monitord.log",
            }
            
            Daemon.Daemon(func(Unix *net.UnixListener) {
                defer Unix.Close()
                
                for {
                    if Fd, err := Unix.AcceptUnix(); err != nil {
                        log.Printf("%v\r\n", err)
                    } else {
                        go Scheduler(Fd)
                    }
                }
            })
            return nil
        },
    }
    
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}