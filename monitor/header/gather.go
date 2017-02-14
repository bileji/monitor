package header

import (
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