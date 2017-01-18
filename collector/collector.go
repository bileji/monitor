package collector

import (
    "encoding/json"
)

type Collector struct {
    Cpu     *Cpu `json:"cpu"`
    Docker  *Docker `json:"docker"`
    Memory  *Memory `json:"memory"`
    Network *Network `json:"network"`
    Disk    *Disk `json:"disk"`
}

func Start() ([]byte, error) {
    return json.Marshal(Collector{
        Cpu: Cpu{}.Exec(),
        Docker: Docker{}.Exec(),
        Memory: Memory{}.Exec(),
        Network: Network{}.Exec(),
        Disk: Disk{}.Exec(),
    })
}