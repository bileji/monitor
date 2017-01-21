package daemon

import (
    "os"
    "log"
    "fmt"
    "syscall"
    "os/signal"
)

type Daemon struct {
    LogFile string
    PidFile string
}

func (D *Daemon) Daemon(routines... func()) {
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
        args := append([]string{os.Args[0]}, os.Args[1:]...)
        os.StartProcess(os.Args[0], args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
        return
    }
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
    
    for _, routine := range routines {
        go routine()
    }
    
    for {
        switch <-Signal {
        case syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
            if err := D.ClearFile(PidFile); err == nil {
                log.Println("success to exit proc, bye bye!")
                LogFile.Close()
                os.Exit(1)
            } else {
                log.Printf("fail to exit proc: %v\r\n", err)
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