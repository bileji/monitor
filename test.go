package main

import (
    "os"
    "github.com/VividCortex/godaemon"
    "syscall"
    //"bufio"
    "time"
    //"github.com/Sirupsen/logrus"
    "strconv"
)

func Daemon(ProgramName string, PidFile string) {
    var (
        err error
        File *os.File
    )
    
    if godaemon.Stage() == godaemon.StageParent {
        File, err = os.OpenFile(PidFile, os.O_CREATE | os.O_RDWR, 0644)
        if err != nil {
            os.Exit(1)
        }
        defer File.Close()
        Info, err := File.Stat()
        
        if err != nil {
            os.Exit(1)
        }
        if Info.Size() > 0 {
            //todo print message
            os.Exit(1)
        }
        
        if syscall.Flock(int(File.Fd()), syscall.LOCK_EX) != nil {
            os.Exit(0)
        }
        
    }
    
    godaemon.MakeDaemon(&godaemon.DaemonAttr{
        ProgramName: ProgramName,
        Files: []**os.File{&File},
    })
    
    File.WriteString(strconv.Itoa(os.Getpid()))
}

func main() {
    
    Daemon("test", "test.pid")
    
    for {
        //logrus.Println("2s")
        time.Sleep(time.Second * 1)
    }
}

//func main()  {
//    var (
//        f   *os.File
//        err error
//    )
//
//    if godaemon.Stage() == godaemon.StageParent {
//        f, err = os.OpenFile("./log", os.O_RDWR | os.O_CREATE, 0644)
//        defer f.Close()
//        if err != nil {
//            os.Exit(1)
//        }
//        err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
//        if err != nil {
//            os.Exit(1)
//        }
//    }
//
//    Out, _,  err := godaemon.MakeDaemon(&godaemon.DaemonAttr{
//        ProgramName: "test",
//        CaptureOutput: true,
//        Files: []**os.File{&f},
//    })
//
//
//    for {
//        fmt.Fprintln(os.Stdout, "out")
//
//        Read := bufio.NewReader(Out)
//        Buf := make([]byte, 30)
//        Read.Read(Buf)
//        f.Write(Buf)
//
//        time.Sleep(time.Second * 2)
//    }
//
//    logrus.WithFields(logrus.Fields{})
//}