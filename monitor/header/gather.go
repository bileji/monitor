package header

import "monitor/monitor/collector"

type Gather struct {
    Cpu     *collector.Cpu
    Docker  *collector.Docker
    Memory  *collector.Memory
    Network *collector.Network
}
