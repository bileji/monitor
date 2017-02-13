package server

import (
    "log"
    "errors"
    "io/ioutil"
    "encoding/json"
    "monitor/monitor/header"
    "monitor/monitor/helper"
    "monitor/monitor/collector"
    "github.com/noaway/heartbeat"
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

func (n *Node) gather(Spec int) error {
    ht, err := heartbeat.NewTast("gather", Spec)
    if err != nil {
        return err
    }
    ht.Start(func() {
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
        if Answer.Code == header.FAILURE {
            log.Println("gather failure")
        }
        return nil
    })
    return nil
}

func (n *Node) RunForever() error {
    err := n.Verify()
    if err != nil {
        return err
    }
    
    // 收集信息
    go n.gather(5)
    
    return nil
}