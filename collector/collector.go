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
    c := &Collector{
        Cpu: Cpu{}.Exec(),
        Network: Network{}.Exec(),
    }
    //c.Docker.Exec()
    //c.Memory.Exec()
    //c.Network.Exec()
    return json.Marshal(*c)
}