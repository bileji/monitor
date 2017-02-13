package server

import (
    "errors"
    "io/ioutil"
    "encoding/json"
    "monitor/monitor/header"
    "monitor/monitor/helper"
)

type Node header.Node

func (n *Node) Verify() error {
    R, err := helper.Request("POST", n.Addr + "/verify", n.Token)
    if err != nil {
        return err
    }
    defer R.Body.Close()
    Body, _ := ioutil.ReadAll(R.Body)
    var Answer header.Answer
    json.Unmarshal(Body, &Answer)
    if Answer.Code == 0 {
        return nil
    } else {
        return errors.New("verify token failure")
    }
}