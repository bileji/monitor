package slave

import (
    "encoding/json"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/load"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/docker"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/net"
)

type SysInfo struct {
    Cpu           []cpu.TimesStat `json:"cpu"`
    Load          *load.AvgStat `json:"load"`
    DiskMount     *disk.UsageStat `json:"disk_mount"`
    Docker        []docker.CgroupDockerStat `json:"docker"`
    VirtualMemory *mem.VirtualMemoryStat `json:"memory"`
    Host          *host.InfoStat `json:"host"`
    Net           []net.ConnectionStat `json:"net"`
}

func GetSysInfo() *SysInfo {

    c, err := cpu.Times(false)
    if err != nil {
        c = []cpu.TimesStat{}
    }

    vm, err := mem.VirtualMemory()
    if err != nil {
        vm = &mem.VirtualMemoryStat{}
    }

    lo, err := load.Avg()
    if err != nil {
        lo = &load.AvgStat{}
    }

    di, err := disk.Usage("/data")
    if err != nil {
        di = &disk.UsageStat{}
    }

    do, err := docker.GetDockerStat()
    if err != nil {
        do = []docker.CgroupDockerStat{}
    }

    ho, err := host.Info()
    if err != nil {
        ho = &host.InfoStat{}
    }

    ne, err := net.Connections("tcp")
    if err != nil {
        ne = []net.ConnectionStat{}
    }

    return &SysInfo{
        Cpu: c,
        VirtualMemory: vm,
        Load: lo,
        DiskMount: di,
        Docker: do,
        Host: ho,
        Net: ne,
    }
}

func (sys *SysInfo) ToString() string {
    bytes, _ := json.Marshal(sys)
    return string(bytes)
}


