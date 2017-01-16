package main

import (
    "os"
    "log"
    "fmt"
    "monitor/daemon"
    //"monitor/collector"
)

func main() {
    
    File, err := os.Create("/var/log/monitord.log")
    if err != nil {
        fmt.Println("创建日志文件错误", err)
        return -1, err
    }
    log.SetOutput(File)
    
    Daemon := &daemon.Daemon{
        PidFile: "/var/run/monitord.pid",
        LogFile: "/var/log/monitord.log",
    }
    _, err = Daemon.Start(0, 0)
    if err != nil {
        log.Println(err)
        os.Exit(-1)
    }
    
    //Daemon.Signal()
    
    //for {
    //    time.Sleep(1 * time.Second)
    //    log.Println("hello monitor")
    //bytes, err := collector.Start()
    //if err != nil {
    //    log.Println(err)
    //} else {
    //    log.Println(string(bytes))
    //}
    //}
}
