package main

import (
    "log"
    "time"
    "monitor/daemon"
    //"monitor/collector"
)

func main() {
    
    Daemon := &daemon.Daemon{
        PidFile: "/var/run/monitord.pid",
        LogFile: "/var/log/monitord.log",
    }
    Daemon.Daemon()
    
    for {
        time.Sleep(1 * time.Second)
        log.Println("hello monitor")
        //bytes, err := collector.Start()
        //if err != nil {
        //    log.Println(err)
        //} else {
        //    log.Println(string(bytes))
        //}
    }
}
