package daemon

import (
    "os"
    "log"
    "syscall"
    "os/signal"
    "fmt"
)

type Daemon struct {
    LogFile string
    PidFile string
}

func (D *Daemon) Daemon() {
    
    File, err := os.OpenFile(D.LogFile, os.O_RDWR | os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("create log error: %v\r\n", err)
        return
    }
    log.SetOutput(File)
    
    File, err = os.OpenFile(D.PidFile, os.O_RDWR | os.O_CREATE, 0644)
    if err != nil {
        log.Printf("read pid file error: %v\r\n", err)
        return
    }
    Info, _ := File.Stat()
    if Info.Size() != 0 {
        log.Printf("pid file is exist: %s\r\n", D.LogFile)
        return
    }
    if os.Getppid() != 1 {
        args := append([]string{os.Args[0]}, os.Args[1:]...)
        os.StartProcess(os.Args[0], args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
        return
    }
    if _, err = File.WriteString(fmt.Sprint(os.Getpid())); err != nil {
        log.Printf("fail write pid to %s: %v", D.PidFile, err)
        return
    }
    Signal := make(chan os.Signal, 1)
    signal.Notify(Signal, syscall.SIGTERM, os.Interrupt, syscall.SIGUSR2)
    for {
        fmt.Println("+++++++go forever")
        switch <-Signal {
        case syscall.SIGTERM:
            fmt.Println("安全退出")
            Exit(File)
            os.Exit(1)
        case syscall.SIGUSR2:
            fmt.Println("自定义型号.")
        case os.Interrupt:
            fmt.Println("安全退出")
            Exit(File)
        }
    }
}

func Exit(F *os.File) {
    F.Close()
    os.Remove(F.Name())
}