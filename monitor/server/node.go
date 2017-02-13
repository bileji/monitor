package server

import (
    "log"
    "errors"
    "io/ioutil"
    "encoding/json"
    "monitor/monitor/header"
    "monitor/monitor/helper"
    "monitor/monitor/collector"
)

type Node header.Node

func (n *Node) Verify() error {
    R, err := helper.Request(header.METHOD, header.SCHEMA + n.Addr + "/verify", n.Token)
    if err != nil {
        return err
    }
    defer R.Body.Close()
    Body, _ := ioutil.ReadAll(R.Body)
    var Answer header.Answer
    json.Unmarshal(Body, &Answer)
    if Answer.Code == header.SUCCESS {
        return nil
    } else {
        return errors.New("verify token failure")
    }
}

func (n *Node) RunForever() error {
    err := n.Verify()
    if err != nil {
        return err
    }
    
    go func() {
        Gather, _ := json.Marshal(header.Gather{
            Cpu: collector.Cpu{}.Exec(),
            Docker: collector.Docker{}.Exec(),
            Memory: collector.Memory{}.Exec(),
            Network: collector.Network{}.Exec(),
        })
        
        R, err := helper.Request(header.METHOD, header.SCHEMA + n.Addr + "/gather", Gather)
        if err != nil {
            log.Printf("%v", err)
        }
        defer R.Body.Close()
        Body, _ := ioutil.ReadAll(R.Body)
        var Answer header.Answer
        json.Unmarshal(Body, &Answer)
        if Answer.Code != header.SUCCESS {
            log.Println("gather failure")
        }
    }()
    
    return nil
}