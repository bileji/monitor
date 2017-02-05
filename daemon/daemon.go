package daemon

import (
    "os"
    "log"
    "fmt"
    "net"
    "os/exec"
    "syscall"
    "os/signal"
    "path/filepath"
)

type Daemon struct {
    LogFile  string
    PidFile  string
    UnixFile string
}

type Protocol struct {
    Name  string
    Value string
}

func (D *Daemon) Daemon(Routine func(chan []byte, *net.UnixListener)) {
    PidFile, err := os.OpenFile(D.PidFile, os.O_CREATE | os.O_RDWR, 0644)
    if err != nil {
        fmt.Printf("read pid file error: %v\r\n", err)
        return
    }
    
    if Info, _ := PidFile.Stat(); Info.Size() != 0 {
        fmt.Printf("pid file is exist: %s\r\n", D.PidFile)
        return
    }
    if os.Getppid() != 1 {
        file, _ := exec.LookPath(os.Args[0])
        path, _ := filepath.Abs(file)
        //args := append([]string{os.Args[0]}, os.Args[1:]...)
        args := append([]string{path}, os.Args[1:]...)
        os.StartProcess(os.Args[0], args, &os.ProcAttr{Dir: "/", Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
        return
    }
    fmt.Println(os.Environ())
    if _, err = PidFile.WriteString(fmt.Sprint(os.Getpid())); err != nil {
        fmt.Printf("fail write pid to %s: %v\r\n", D.PidFile, err)
        return
    }
    Signal := make(chan os.Signal, 1)
    signal.Notify(Signal, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
    
    LogFile, err := os.OpenFile(D.LogFile, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0644)
    if err != nil {
        fmt.Printf("create log error: %v\r\n", err)
        return
    }
    log.SetOutput(LogFile)
    
    go D.UnixListen(Routine)
    
    for {
        switch <-Signal {
        case syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
            if err := D.ClearFile(PidFile); err == nil {
                if err := os.Remove(D.UnixFile); err == nil {
                    log.Println("success to exit proc, bye bye!")
                } else {
                    log.Printf("fail to remove unix sock: %v\r\n", err)
                }
                LogFile.Close()
                os.Exit(1)
            } else {
                log.Printf("fail to remove process pid file: %v\r\n", err)
            }
        default:
            log.Println("unknow signal, this process will go on...")
        }
    }
}

func (D *Daemon) ClearFile(F *os.File) (error) {
    if err := F.Close(); err != nil {
        return err
    }
    if err := os.Remove(F.Name()); err != nil {
        return err
    }
    return nil
}

func (D *Daemon) UnixListen(Routine func(ch chan []byte, wr *net.UnixListener)) {
    UnixL, err := net.ListenUnix("unix", &net.UnixAddr{Name: D.UnixFile, Net: "unix"})
    if err != nil {
        log.Printf("listen unix error: %v", err)
    }
    defer UnixL.Close()
    
    C := make(chan []byte, 1)
    
    go Routine(C, UnixL)
    
    for {
        if Fd, err := UnixL.AcceptUnix(); err != nil {
            log.Printf("accect error: %v", err)
        } else {
            for {
                Buffer := make([]byte, 512)
                if Len, err := Fd.Read(Buffer); err != nil {
                    log.Printf("accect error: %v", err)
                } else {
                    C <- Buffer[0: Len]
                }
            }
        }
    }
}