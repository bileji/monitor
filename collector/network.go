package collector

import (
    "os/exec"
    "strings"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/net"
    "monitor/common"
)

type Network struct {
    // echo `nc ns1.dnspod.net 6666`
    PublicIP       string `json:"public_ip"`
    HostInfo       *host.InfoStat `json:"host_info"`
    InterfacesStat []net.InterfaceStat `json:"interfaces_stat"`
}

func (n *Network) GetPublicIP() string {
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

func (n *Network) Exec() *Network {
    var err error
    n.HostInfo, err = host.Info()
    if err != nil {
        n.HostInfo = &host.InfoStat{}
    }
    
    n.PublicIP = n.GetPublicIP()
    
    n.InterfacesStat, err = net.Interfaces()
    if err != nil {
        n.InterfacesStat = []net.InterfaceStat{}
    }
    return n
}
