package main

import (
    "log"
    "time"
    "monitor/daemon"
    "monitor/service"
    "gopkg.in/mgo.v2"
    "monitor/collector/model"
)

func main() {
    
    Daemon := &daemon.Daemon{
        PidFile: "/var/run/monitord.pid",
        LogFile: "/var/log/monitord.log",
    }
    
    Session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }
    if Session.DB("vpn").Login("shadowsocks", "mlgR4evB") != nil {
        panic(err)
    }
    
    Daemon.Daemon(func() {
        defer Session.Close()
        
        go (&service.Master{
            Addr: ":88",
            DBHandler: Session.DB("vpn"),
        }).Listen()
        
        for {
            time.Sleep(2 * time.Second)
            bytes, err := model.Gather{}.Exec()
            if err != nil {
                log.Println(err)
            } else {
                log.Println(string(bytes))
            }
        }
    })
}
