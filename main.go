package main

import (
    "log"
    "time"
    "monitor/daemon"
)

func main() {
    Daemon := &daemon.Daemon{
        LogFile: "/var/log/monitord.log",
    }
    Daemon.Start(0, 0)

    for {
        time.Sleep(1 * time.Second)
        log.Println("hello monitor")
    }
}
