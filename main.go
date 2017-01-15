package main

import (
    "monitor/daemon"
)

func main() {
    Daemon := &daemon.Daemon{
        LogFile: "/var/log/monitord.log",
    }
    Daemon.Start(0, 0)
}
