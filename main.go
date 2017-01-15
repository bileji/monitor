package main

import (
    "log"
    "time"
    "monitor/daemon"
    "monitor/collector"
)

func main() {
    Daemon := &daemon.Daemon{
        LogFile: "/var/log/monitord.log",
    }
    Daemon.Start(0, 0)
    for {
        log.Println("hello monitor")
        time.Sleep(10 * time.Second)
        bytes, _ := collector.Start()
        log.Println(string(bytes))
        
    }
}
