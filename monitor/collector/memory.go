package collector

import (
    "github.com/shirou/gopsutil/mem"
)

type Memory struct {
    SwapMemory    *mem.SwapMemoryStat `json:"swap_memory"`
    VirtualMemory *mem.VirtualMemoryStat `json:"virtual_memory"`
}

func (m Memory) Exec() *Memory {
    var err error
    if m.SwapMemory, err = mem.SwapMemory(); err != nil {
        m.SwapMemory = &mem.SwapMemoryStat{}
    }
    if m.VirtualMemory, err = mem.VirtualMemory(); err != nil {
        m.VirtualMemory = &mem.VirtualMemoryStat{}
    }
    return &m
}