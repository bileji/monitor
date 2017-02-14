package collector

import (
    "os/exec"
    "strings"
    "monitor/monitor/collector/common"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/net"
)

type Network struct {
    PublicIP       string `json:"public_ip"`
    HostInfo       *host.InfoStat `json:"host_info"`
    InterfacesStat []net.InterfaceStat `json:"interfaces_stat"`
}

// echo `nc ns1.dnspod.net 6666`
func (n Network) GetPublicIP() string {
    nc, err := exec.LookPath("/usr/bin/nc")
    if err != nil {
        return ""
    }
    out, err := common.Invoke{}.Command(nc, "ns1.dnspod.net", "6666")
    if err != nil {
        return ""
    }
    lines := strings.Split(string(out), "\n")
    if len(lines) > 0 {
        return lines[0]
    }
    return ""
}

func (n Network) Gather() *Network {
    var err error
    n.PublicIP = n.GetPublicIP()
    if n.HostInfo, err = host.Info(); err != nil {
        n.HostInfo = &host.InfoStat{}
    }
    if n.InterfacesStat, err = net.Interfaces(); err != nil {
        n.InterfacesStat = []net.InterfaceStat{}
    }
    return &n
}
