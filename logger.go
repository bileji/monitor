package main

import (
    log "github.com/Sirupsen/logrus"
    "runtime"
)

func main() {
    
    line := func() int {
        _, _, line, _ := runtime.Caller(1)
        return line
    }
    
    fname := func() uintptr{
        fname, _, _, _ := runtime.Caller(2)
        return fname
    }
    
    log.SetFormatter(&log.TextFormatter{
        FullTimestamp: true,
        TimestampFormat: "2006-01-02 15:04:05",
    })
    
    log.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("A group of walrus emerges from the ocean")
    
    log.WithFields(log.Fields{
        "omg":    true,
        "number": line(),
    }).Warn("The group's number increased tremendously!")
    
    //log.SetFormatter(&log.JSONFormatter{
    //    TimestampFormat: "2006-01-02 15:04:05",
    //})
    log.Warn("xxxxx")
    
    // A common pattern is to re-use fields between logging statements by re-using
    // the logrus.Entry returned from WithFields()
    contextLogger := log.WithFields(log.Fields{
        "line": line(),
        "func": runtime.FuncForPC(fname()).Name(),
    })
    
    contextLogger.Info("I'll be logged with common and other field")
    contextLogger.Info("Me too")

    log.WithFields(log.Fields{
        "omg":    true,
        "number": 100,
    }).Fatal("The ice breaks!")
}