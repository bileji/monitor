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
    
    Unix, err := net.DialUnix("unix", &net.UnixAddr{Name: "/var/run/monitord.sock", Net: "unix"}, &net.UnixAddr{Name: "/var/run/monitor.sock", Net: "unix"})
    
    if err != nil {
        fmt.Printf("%v\r\n", err)
    }
    
    RootCmd := &cobra.Command{
        Use: "monitord",
        Short: "Linux server status monitor",
        Long: "",
        Run: func(cmd *cobra.Command, args []string) {
            // TODO ...
            Message, _ := json.Marshal(protocols.Socket{Method: "test", Body: []byte(""), Timestamp: 1234567890})
            Unix.Write(Message)
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
