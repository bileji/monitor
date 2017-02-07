package model

import (
    "encoding/json"
    "monitor/collector"
)

type Gather struct {
    Cpu     *collector.Cpu `json:"cpu"`
    Docker  *collector.Docker `json:"docker"`
    Memory  *collector.Memory `json:"memory"`
    Network *collector.Network `json:"network"`
    Disk    *collector.Disk `json:"disk"`
}

func (g Gather)Exec() ([]byte, error) {
    return json.Marshal(Gather{
        Cpu: collector.Cpu{}.Exec(),
        Docker: collector.Docker{}.Exec(),
        Memory: collector.Memory{}.Exec(),
        Network: collector.Network{}.Exec(),
        Disk: collector.Disk{}.Exec(),
    })
}