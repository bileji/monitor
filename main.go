package main

import (
    "log"
    "time"
    "monitor/daemon"
    "monitor/collector"
    "monitor/service"
)

func main() {
    
    go (&service.Master{
        Addr: ":88",
    }).Listen()
    
    Daemon := &daemon.Daemon{
        PidFile: "/var/run/monitord.pid",
        LogFile: "/var/log/monitord.log",
    }
    
    Daemon.Daemon(func() {
        for {
            time.Sleep(2 * time.Second)
            bytes, err := collector.Start()
            if err != nil {
                log.Println(err)
            } else {
                log.Println(string(bytes))
            }
        }
    })
}
