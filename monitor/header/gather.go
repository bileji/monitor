package header

import (
    "encoding/json"
    "monitor/monitor/collector"
)

type Gather struct {
    Cpu      *collector.Cpu         `json:"cpu"`
    Docker   *collector.Docker      `json:"docker"`
    Memory   *collector.Memory      `json:"memory"`
    Network  *collector.Network     `json:"network"`
    Disk     *collector.Disk        `json:"disk"`
    Created  int64                  `json:"created"`
    Modified int64                  `json:"modified"`
}

func (g *Gather) Exec() ([]byte, error) {
    return json.Marshal(Gather{
        Cpu:     collector.Cpu{}.Gather(),
        Docker:  collector.Docker{}.Gather(),
        Memory:  collector.Memory{}.Gather(),
        Disk:    collector.Disk{}.Gather(),
        Network: collector.Network{}.Gather(),
    })
}