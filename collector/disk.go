package collector

import "github.com/shirou/gopsutil/disk"

type Disk struct {
    UseAge     *disk.UsageStat `json:"use_age"`
    IOCounters map[string]disk.IOCountersStat `json:"io_counters"`
}

func (d Disk) Exec() *Disk {
    var err error
    if d.UseAge, err = disk.Usage("/data"); err != nil {
        d.UseAge = *disk.UsageStat{}
    }
    if d.IOCounters, err = disk.IOCounters(); err != nil {
        d.IOCounters = make(map[string]disk.IOCountersStat)
    }
    return &d
}