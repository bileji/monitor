package daemon

import (
    "os"
    "log"
    "fmt"
    "net"
    "strings"
    "errors"
    "os/exec"
    "syscall"
    "os/signal"
    "path/filepath"
)

type Daemon struct {
    LogFile  string
    PidFile  string
    UnixFile string
    
    Pid      *os.File
    Log      *os.File
}

type Protocol struct {
    Name  string
    Value string
}

func currentPath(arg string) string {
    File, err := exec.LookPath(arg)
    if err != nil {
        return arg
    }
    Path, err := filepath.Abs(File)
    if err != nil {
        return arg
    }
    return Path
}

func (D *Daemon) CreatePidFile() error {
    var err error
    D.Pid, err = os.OpenFile(D.PidFile, os.O_CREATE | os.O_RDWR, 0644)
    if err != nil {
        return err
    }
    
    if Info, _ := D.Pid.Stat(); Info.Size() != 0 {
        return errors.New("pid file is exist: " + D.PidFile)
    }
    return nil
}

func (D *Daemon) WritePidFile() error {
    _, err := D.Pid.WriteString(fmt.Sprint(os.Getpid()))
    return err
}

func (D *Daemon) Daemon(Routine func(*net.UnixListener)) {
    if err := D.CreatePidFile(); err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    
    if os.Getppid() != 1 {
        os.Args[0] = currentPath(os.Args[0])
        args := append([]string{os.Args[0]}, os.Args[1:]...)
        os.StartProcess(os.Args[0], args, &os.ProcAttr{Dir: "/", Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
        return
    }
    if err := D.WritePidFile(); err != nil {
        fmt.Printf("fail write pid to %s: %v\n", D.PidFile, err)
        return
    }
    
    var err error
    D.Log, err = os.OpenFile(D.LogFile, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0644)
    if err != nil {
        fmt.Printf("create log error: %v\n", err)
        return
    }
    log.SetOutput(D.Log)
    
    go D.UnixListen(Routine)
    
    D.Signal()
}

func (D *Daemon) Signal() {
    var Println = func(Str... string) {
        if D.Log == nil {
            fmt.Println(strings.Join(Str, ""))
        } else {
            log.Println(strings.Join(Str, ""))
        }
    }
    
    var PrintF = func(Format string, Inter interface{}) {
        if D.Log == nil {
            fmt.Printf(Format, Inter)
        } else {
            log.Printf(Format, Inter)
        }
    }
    
    Signal := make(chan os.Signal, 1)
    signal.Notify(Signal, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
    for {
        switch <-Signal {
        case syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
            if err := D.ClearPidFile(); err == nil {
                if err := os.Remove(D.UnixFile); err == nil {
                    Println("success to exit process!")
                } else {
                    PrintF("fail to remove unix sock: %v\n", err)
                }
                
                os.Exit(1)
            } else {
                PrintF("fail to remove process pid file: %v\n", err)
            }
        default:
            Println("unknow signal, this process will go on...")
        }
    }
}

func (D *Daemon) ClearPidFile() (error) {
    if err := D.Pid.Close(); err != nil {
        return err
    }
    if err := os.Remove(D.Pid.Name()); err != nil {
        return err
    }
    return nil
}

func (D *Daemon) UnixListen(Routine func(*net.UnixListener)) {
    os.Remove(D.UnixFile)
    UnixL, err := net.ListenUnix("unix", &net.UnixAddr{Name: D.UnixFile, Net: "unix"})
    if err != nil {
        log.Printf("%v\n", err)
    }
    
    Routine(UnixL)
}