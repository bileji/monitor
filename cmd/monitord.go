package main

import (
    "os"
    "fmt"
    "net"
    "runtime"
    "monitor/daemon"
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
    
    Daemon.Daemon(func(ch chan []byte, wr *net.UnixListener) {
        
    })
}