package collector

import (
    "os/exec"
    "monitor/monitor/collector/common"
    "strings"
    //"strconv"
    "fmt"
)

type Process struct {
    User    string
    Pid     int
    Cpu     float64
    Memory  float64
    Vsz     int
    Rss     int
    Tty     string
    Stat    string
    Start   string
    Time    string
    Command string
}

func (p Process) Get(Reg string) []Process {
    
    var Pros []Process
    Ps, err := exec.LookPath("ps")
    if err != nil {
        return Pros
    }
    Out, err := common.Invoke{}.Command(Ps, "aux")
    //Out, err := common.Invoke{}.Command(Ps, "aux", "|grep -E", Reg, "|grep -v grep")
    if err != nil {
        return Pros
    }
    Lines := strings.Split(string(Out), "\n")
    
    for _, Line := range Lines {
        fmt.Println(strings.Fields(Line))
        fmt.Println("+++++-------------++++++++")
    //    Info := strings.Split(Line, " ")
    //    Pid, _ := strconv.Atoi(Info[1])
    //    Cpu, _ := strconv.ParseFloat(Info[2], 64)
    //    Memory, _ := strconv.ParseFloat(Info[3], 64)
    //    Vsz, _ := strconv.Atoi(Info[4])
    //    Rss, _ := strconv.Atoi(Info[5])
    //    Pros = append(Pros, Process{
    //        User: Info[0],
    //        Pid: Pid,
    //        Cpu: Cpu,
    //        Memory: Memory,
    //        Vsz: Vsz,
    //        Rss: Rss,
    //        //Tty: Info[6],
    //        //Stat: Info[7],
    //        //Start: Info[8],
    //        //Time: Info[9],
    //        //Command: Info[10],
    //    })
    }
    
    return Pros
}

func (p Process) Gather(Reg string) []Process {
    return p.Get(Reg)
}
