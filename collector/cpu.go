package collector

import (
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/load"
)

type Cpu struct {
    Load      *load.AvgStat `json:"load"`
    TimesStat []cpu.TimesStat `json:"times_stat"`
}

func (c Cpu) Exec() Cpu {
    var err error
    c.Load, err = load.Avg()
    if err != nil {
        c.Load = &load.AvgStat{}
    }
    c.TimesStat, err = cpu.Times(false)
    if err != nil {
        c.TimesStat = []cpu.TimesStat{}
    }
    return c
}