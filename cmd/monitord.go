package main

import (
    "os"
    "fmt"
    "log"
    "net"
    "runtime"
    "encoding/json"
    "monitor/daemon"
    "monitor/cmd/protocols"
    "monitor/cmd/configures"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "monitor/cmd/commands"
)

const (
    READ_LENGTH = 2048
)

type DBAuth struct {
    Host     string `json:"host"`
    Port     int `json:"port"`
    Database string `json:"database"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func Scheduler(Con *net.UnixConn) {
    for {
        Buffer := make([]byte, READ_LENGTH)
        Len, err := Con.Read(Buffer);
        if err != nil {
            Con.Close()
            return
        }
        var Message protocols.Socket
        json.Unmarshal(Buffer[0: Len], &Message)
        // todo 接收到cli信息,然后处理
        log.Println(Message)
    }
}

func Monitor(Unix *net.UnixListener) {
    defer Unix.Close()
    
    for {
        if Fd, err := Unix.AcceptUnix(); err != nil {
            log.Printf("%v\r\n", err)
        } else {
            go Scheduler(Fd)
        }
    }
}

func main() {
    var (
        DaemonB bool
        
        ConfFile string
        
        PidFile string
        LogFile string
    )
    
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    Viper := viper.GetViper()
    
    RootCmd := &cobra.Command{
        Use: "monitord",
        Short: "Linux server status monitor daemon",
        Long: "######to do...",
        RunE:func(cmd *cobra.Command, args []string) error {
            
            Conf := configures.Initialize(Viper, ConfFile)
            
            Daemon := &daemon.Daemon{
                PidFile: Conf.Server.PidFile,
                UnixFile: Conf.Server.UnixFile,
                LogFile: Conf.Server.LogFile,
            }
            
            if DaemonB == false {
                Daemon.UnixListen(Monitor)
            }
            
            return nil
        },
    }
    
    Flags := RootCmd.Flags()
    
    // 配置文件路径
    Flags.StringVarP(&ConfFile, "config", "c", "/etc/monitor.toml", "configuration file specifying additional options")
    Flags.BoolVarP(&DaemonB, "daemon", "d", true, "to start the daemon way")
    
    Flags.StringVarP(&PidFile, "pid", "", "", "full path to pidfile")
    Flags.StringVarP(&LogFile, "log", "l", "", "log file")
    
    Viper.BindPFlag("server.daemon", Flags.Lookup("daemon"))
    Viper.BindPFlag("server.pid_file", Flags.Lookup("pid"))
    Viper.BindPFlag("server.log_file", Flags.Lookup("log"))
    
    RootCmd.AddCommand(commands.VersionCmd)
    
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}