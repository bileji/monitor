package collector

import (
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/load"
)

type Cpu struct {
    Load      *load.AvgStat `json:"load"`
    TimesStat []cpu.TimesStat `json:"times_stat"`
}

func (c Cpu) Gather() *Cpu {
    var err error
    if c.Load, err = load.Avg(); err != nil {
        c.Load = &load.AvgStat{}
    }
    if c.TimesStat, err = cpu.Times(false); err != nil {
        c.TimesStat = []cpu.TimesStat{}
    }
    return &c
}