package main

import (
    "os"
    "fmt"
    "runtime"
    "monitor/cmd/commands"
)

func main() {
    // 调优
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    err := commands.MonitorCmd.Execute()
    
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    
    //Daemon := &daemon.Daemon{
    //    PidFile: "/var/run/monitord.pid",
    //    UnixFile: "/var/run/monitord.sock",
    //    LogFile: "/var/log/monitord.log",
    //}
    //
    //Session, err := mgo.Dial("127.0.0.1:27017")
    //if err != nil {
    //    panic(err)
    //}
    //if Session.DB("vpn").Login("shadowsocks", "mlgR4evB") != nil {
    //    panic(err)
    //}
    //
    //Daemon.Daemon(func(ch chan []byte) {
    //    defer Session.Close()
    //
    //    go (&service.Master{
    //        Addr: ":88",
    //        DBHandler: Session.DB("vpn"),
    //    }).Listen()
    //
    //    for {
    //        time.Sleep(2 * time.Second)
    //        bytes, err := model.Gather{}.Exec()
    //        if err != nil {
    //            log.Println(err)
    //        } else {
    //            log.Println(string(bytes))
    //        }
    //    }
    //})
}
