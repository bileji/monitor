package collector

import (
    "os/exec"
    "monitor/monitor/collector/common"
    "strings"
    "strconv"
    "regexp"
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
    if err != nil {
        return Pros
    }
    Lines := strings.Split(string(Out), "\n")
    
    for _, Line := range Lines {
        Str := regexp.MustCompile(Reg).FindAllString(Line, -1)
        if len(Str) <= 0 {
            continue
        }
        var Fields [10]string
        Slice := strings.Fields(Line)
        for i := 0; i < 10; i++ {
            if i <= len(Slice) {
                Fields = append(Fields, Slice[i])
            } else {
                Fields = append(Fields, "")
            }
        }
        fmt.Println(Fields)
        Pid, _ := strconv.Atoi(Fields[1])
        Cpu, _ := strconv.ParseFloat(Fields[2], 64)
        Memory, _ := strconv.ParseFloat(Fields[3], 64)
        Vsz, _ := strconv.Atoi(Fields[4])
        Rss, _ := strconv.Atoi(Fields[5])
        Pros = append(Pros, Process{
            User: Fields[0],
            Pid: Pid,
            Cpu: Cpu,
            Memory: Memory,
            Vsz: Vsz,
            Rss: Rss,
            Tty: Fields[6],
            Stat: Fields[7],
            Start: Fields[8],
            Time: Fields[9],
            Command: Fields[10],
        })
    }
    
    return Pros
}

func (p Process) Gather(Reg string) []Process {
    return p.Get(Reg)
}
