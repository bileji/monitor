package main

import (
    "os"
    "fmt"
    "net"
    "encoding/json"
    "monitor/cmd/protocols"
    "github.com/spf13/cobra"
)

func main() {
    UnixL, err := net.ListenUnix("unix", &net.UnixAddr{Name: "/var/run/monitord.sock", Net: "unix"})
    if err != nil {
        fmt.Printf("%v", err)
    }
    
    RootCmd := &cobra.Command{
        Use: "monitord",
        Short: "Linux server status monitor",
        Long: "",
        Run: func(cmd *cobra.Command, args []string) {
            // TODO ...
            for {
                if Fd, err := UnixL.AcceptUnix(); err != nil {
                    fmt.Printf("%v", err)
                } else {
                    Message, _ := json.Marshal(protocols.Socket{Method: "test", Body: []byte(""), Timestamp: 1234567890})
                    
                    Fd.Write(Message)
                }
            }
        },
        RunE: func(cmd *cobra.Command, args []string) error {
            // TODO
            return nil
        },
    }
    
    if err := RootCmd.Execute(); err != nil {
        fmt.Printf("%v", err)
        os.Exit(-1)
    }
}
