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
    os.Remove("/var/run/monitor.sock")
    lAddr, err := net.ResolveUnixAddr("unix", "/var/run/monitor.sock")
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    
    rAddr, err := net.ResolveUnixAddr("unix", "/var/run/monitord.sock")
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    
    Unix, err := net.DialUnix("unix", lAddr, rAddr)
    
    if err != nil {
        fmt.Println(err)
    }
    
    // TODO remove this example
    Message, _ := json.Marshal(protocols.Socket{Method: "test", Body: []byte(""), Timestamp: 1234567890})
    fmt.Printf(string(Message))
    Unix.Write(Message)
    
    RootCmd := &cobra.Command{
        Use: "monitor",
        Short: "Linux server status monitor",
        Long: "",
        Run: func(cmd *cobra.Command, args []string) {
            // TODO ...
            Message, _ := json.Marshal(protocols.Socket{Method: "test", Body: []byte(""), Timestamp: 1234567890})
            fmt.Printf(string(Message))
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
