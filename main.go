package main

import (
    "monitor/daemon"
)

func main() {
    daemon.Daemon{
        LogFile: "/var/log/monitord.log",
    }.Start(0, 0)
}
