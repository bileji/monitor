package server

import (
    "log"
    "errors"
    "io/ioutil"
    "encoding/json"
    "monitor/monitor/header"
    "monitor/monitor/helper"
    "github.com/noaway/heartbeat"
    "monitor/monitor/collector/model"
    "monitor/monitor/collector"
)

type Gather header.Gather

func (g *Gather) Exec([]byte, error) {
    return json.Marshal(Gather{
        Cpu:     collector.Cpu{}.Gather(),
        Docker:  collector.Docker{}.Gather(),
        Memory:  collector.Memory{}.Gather(),
        Disk:    collector.Disk{}.Gather(),
        Network: collector.Network{}.Gather(),
    })
}

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
    Ht, err := heartbeat.NewTast("gather", Spec)
    if err != nil {
        return err
    }
    Ht.Start(func() error {
        Gather := header.Gather{}
        
        Buffer, _ := Gather.Exec()
        
        R, err := helper.Request(header.METHOD, header.SCHEMA + n.Addr + "/gather", string(Buffer))
        if err != nil {
            log.Printf("%v", err)
        }
        defer R.Body.Close()
        Body, _ := ioutil.ReadAll(R.Body)
        var Answer header.Answer
        json.Unmarshal(Body, &Answer)
        if Answer.Code == header.FAILURE {
            log.Println(Answer)
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