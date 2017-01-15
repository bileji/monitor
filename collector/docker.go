package collector

import (
    "github.com/shirou/gopsutil/docker"
)

type Docker struct {
    DockersStat []docker.CgroupDockerStat `json:"dockers_stat"`
}

func (d Docker) Exec() *Docker {
    var err error
    d.DockersStat, err = docker.GetDockerStat()
    if err != nil {
        d.DockersStat = []docker.CgroupDockerStat{}
    }
    return &d
}
