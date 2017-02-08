package main

import (
    "os"
    "fmt"
    "runtime"
    "monitor/cmd/commands"
)

func main() {
    // 调优
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    err := commands.MonitorCmd.Execute()
    
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
