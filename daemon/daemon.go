package daemon

import (
    "os"
    "log"
    "syscall"
    "errors"
    "strconv"
    "os/signal"
    "fmt"
)

type Daemon struct {
    LogFile    string
    PidFile    string
    PidHandler *os.File
}

func (d *Daemon) Start(ChDir, Close int) (int, error) {
    var err error
    
    File, err := os.Create(d.LogFile)
    if err != nil {
        fmt.Println("创建日志文件错误", err)
        return
    }
    log.SetOutput(File)
    
    // 判断是否已有程序启动 pid
    d.PidHandler, err = os.OpenFile(d.PidFile, os.O_RDWR | os.O_CREATE, 0644)
    if err != nil {
        return -1, err
    }
    Info, _ := d.PidHandler.Stat()
    if Info.Size() != 0 {
        return -1, errors.New("pid file is exist:" + d.PidFile)
    }
    
    // 已经以daemon启动
    if os.Getppid() != 1 {
        args := append([]string{os.Args[0]}, os.Args[1:]...)
        os.StartProcess(os.Args[0], args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
        return 0, nil
    }
    
    d.PidHandler.WriteString(strconv.Itoa(os.Getpid()))
    
    // 处理退出信号
    Signal := make(chan os.Signal, 1)
    signal.Notify(Signal, os.Interrupt, syscall.SIGUSR2)
    for {
        C := <-Signal
        fmt.Println(C)
        switch C {
        case os.Interrupt:
            d.Exit(d.PidHandler)
            log.Println("exit success")
        case syscall.SIGUSR2:
            log.Println("to do for user signal")
        }
    }
    return 0, nil
}

func (d *Daemon) Exit(F *os.File) {
    F.Close()
    os.Remove(F.Name())
}

//func (d *Daemon) Signal() {
//    // 处理退出信号
//    Signal := make(chan os.Signal, 1)
//    signal.Notify(Signal, os.Interrupt, syscall.SIGUSR2)
//    for {
//        C := <-Signal
//        fmt.Println(C)
//        log.Println(C)
//        switch C {
//        case os.Interrupt:
//            d.Exit(d.PidHandler)
//            log.Println("exit success")
//        case syscall.SIGUSR2:
//            log.Println("to do for user signal")
//        }
//    }
//}