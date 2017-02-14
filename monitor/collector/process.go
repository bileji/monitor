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
    User    string      `json:"user"`
    Pid     int         `json:"pid"`
    Cpu     float64     `json:"cpu"`
    Memory  float64     `json:"memory"`
    Vsz     int         `json:"vsz"`
    Rss     int         `json:"rss"`
    Tty     string      `json:"tty"`
    Stat    string      `json:"stat"`
    Start   string      `json:"start"`
    Time    string      `json:"time"`
    Command string      `json:"command"`
    Detail  []string    `json:"detail"`
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
        Fields := make([]string, 11)
        Slice := strings.Fields(Line)
        for i := 0; i <= 10; i++ {
            if i <= len(Slice) {
                Fields[i] = Slice[i]
            } else {
                Fields[i] = ""
            }
        }
        Pid, _ := strconv.Atoi(Fields[1])
        Cpu, _ := strconv.ParseFloat(Fields[2], 64)
        Memory, _ := strconv.ParseFloat(Fields[3], 64)
        Vsz, _ := strconv.Atoi(Fields[4])
        Rss, _ := strconv.Atoi(Fields[5])
        Pro := Process{
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
        }
        if len(Slice) > 11 {
            Pro.Detail = Slice[10:]
        }
        
        Pros = append(Pros, Pro)
    }
    
    return Pros
}

func (p Process) Gather(Reg string) []Process {
    return p.Get(Reg)
}
