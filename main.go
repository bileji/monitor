package main

import (
    "log"
    "time"
    "monitor/daemon"
    "monitor/collector"
    "monitor/service"
    "gopkg.in/mgo.v2"
)

func main() {
    
    Daemon := &daemon.Daemon{
        PidFile: "/var/run/monitord.pid",
        LogFile: "/var/log/monitord.log",
    }
    
    session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }
    if session.DB("vpn").Login("shadowsocks", "mlgR4evB") != nil {
        panic(err)
    }
    
    Daemon.Daemon(func() {
        defer session.Close()
    
        go (&service.Master{
            Addr: ":88",
            DBHandler: session.DB("vpn"),
        }).Listen()
        
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
