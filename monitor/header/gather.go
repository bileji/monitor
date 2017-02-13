package header

import "monitor/monitor/collector"

type Gather struct {
    Cpu     *collector.Cpu      `json:"cpu"`
    Docker  *collector.Docker   `json:"docker"`
    Memory  *collector.Memory   `json:"memory"`
    Network *collector.Network  `json:"network"`
}
