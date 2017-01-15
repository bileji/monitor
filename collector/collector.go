package collector

import (
    "encoding/json"
)

type Collector struct {
    Cpu     *Cpu
    Docker  *Docker
    Memory  *Memory
    Network *Network
}

func Start() ([]byte, error) {
    return json.Marshal(Collector{
        Cpu: Cpu{}.Exec(),
        Docker: Docker{}.Exec(),
        Memory: Memory{}.Exec(),
        Network: Network{}.Exec(),
    })
}