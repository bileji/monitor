package server

import (
    "errors"
    "io/ioutil"
    "encoding/json"
    "monitor/monitor/header"
    "monitor/monitor/helper"
    "monitor/monitor/server/node"
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
    
    // 收集信息
    go node.Gather(n.Addr, 5)
    
    return nil
}