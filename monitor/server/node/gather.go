package node

import (
    "monitor/monitor/header"
    "encoding/json"
    "monitor/monitor/collector"
    "github.com/noaway/heartbeat"
    "monitor/monitor/helper"
    "log"
    "io/ioutil"
)

type Gather header.Gather

func (g *Gather) Exec() ([]byte, error) {
    return json.Marshal(Gather{
        Cpu:     collector.Cpu{}.Gather(),
        Docker:  collector.Docker{}.Gather(),
        Memory:  collector.Memory{}.Gather(),
        Disk:    collector.Disk{}.Gather(),
        Network: collector.Network{}.Gather(),
        Process: collector.Process{}.Gather("'node|nginx|php|mongod|mysql|redis|memcache'"),
    })
}

func Gather(Addr string, Spec int) error {
    Ht, err := heartbeat.NewTast("gather", Spec)
    if err != nil {
        return err
    }
    Ht.Start(func() error {
        Buffer, _ := (&Gather{}).Exec()
        
        R, err := helper.Request(header.METHOD, header.SCHEMA + Addr + "/gather", string(Buffer))
        if err != nil {
            log.Printf("%v", err)
        }
        defer R.Body.Close()
        Body, _ := ioutil.ReadAll(R.Body)
        
        var Answer header.Answer
        err = json.Unmarshal(Body, &Answer)
        if err != nil {
            log.Println(err)
        }
        if Answer.Code == header.FAILURE {
            log.Println(Answer)
        }
        return nil
    })
    return nil
}