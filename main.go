package main

import (
    "os"
    "fmt"
    "runtime"
    "monitor/command"
)

func main() {
    // 调优
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    err := command.MonitorCmd.Command.Execute()
    
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
