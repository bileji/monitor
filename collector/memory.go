package collector

import (
    "github.com/shirou/gopsutil/mem"
)

type Memory struct {
    SwapMemory    *mem.SwapMemoryStat `json:"swap_memory"`
    VirtualMemory *mem.VirtualMemoryStat `json:"virtual_memory"`
}

func (m *Memory) Exec() *Memory {
    var err error
    m.VirtualMemory, err = mem.VirtualMemory()
    if err != nil {
        m.VirtualMemory = &mem.VirtualMemoryStat{}
    }
    m.SwapMemory, err = mem.SwapMemory()
    if err != nil {
        m.SwapMemory = &mem.SwapMemoryStat{}
    }
    return m
}