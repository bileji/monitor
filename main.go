package main

import (
    "os"
    "log"
    "time"
    "monitor/daemon"
)

func main() {
    
    File, err := os.Create("log")
    if err != nil {
        log.Println("创建日志文件错误", err)
        return
    }
    log.SetOutput(File)

    Daemon := &daemon.Daemon{
        LogFile: "/var/log/monitord.log",
    }
    Daemon.Start(0, 0)
    for {
        time.Sleep(1 * time.Second)
        log.Println("hello monitor")
    }
}
