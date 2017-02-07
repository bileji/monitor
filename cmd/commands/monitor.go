package commands

import (
    "log"
    "net"
    "encoding/json"
    "monitor/monitor/daemon"
    "monitor/cmd/protocols"
    "monitor/cmd/configures"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

const (
    READ_LENGTH = 2048
)

var (
    Daemon bool
    
    ConfFile string
    
    PidFile string
    LogFile string
    
    Viper = viper.GetViper()
    
    Serve = func(cmd *cobra.Command, args []string) error {
        Conf := configures.Initialize(Viper, ConfFile)
        log.Println("rune")
        Daemon := &daemon.Daemon{
            PidFile: Conf.Server.PidFile,
            UnixFile: Conf.Server.UnixFile,
            LogFile: Conf.Server.LogFile,
        }
        
        if Conf.Server.Daemon {
            Daemon.Daemon(monitor)
        } else {
            Daemon.UnixListen(monitor)
        }
        
        return nil
    }
)
//type DBAuth struct {
//    Host     string `json:"host"`
//    Port     int `json:"port"`
//    Database string `json:"database"`
//    Username string `json:"username"`
//    Password string `json:"password"`
//}

func scheduler(Con *net.UnixConn) {
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

func monitor(Unix *net.UnixListener) {
    defer Unix.Close()
    
    for {
        if UnixConn, err := Unix.AcceptUnix(); err != nil {
            log.Printf("%v\r\n", err)
        } else {
            go scheduler(UnixConn)
        }
    }
}

func addCommand(Cmd *cobra.Command) {
    Cmd.AddCommand(VersionCmd)
}

var MonitorCmd = &cobra.Command{
    Use: "monitor",
    Short: "Linux server status monitor",
    Long: "Powerful Linux server status monitor server",
    RunE: Serve,
}

func init() {
    
    Flags := MonitorCmd.Flags()
    Flags.StringVarP(&ConfFile, "config", "c", "/etc/monitor.toml", "configuration file specifying additional options")
    
    Flags.BoolVarP(&Daemon, "daemon", "d", false, "to start the daemon way")
    Flags.StringVarP(&PidFile, "pid", "", "", "full path to pidfile")
    Flags.StringVarP(&LogFile, "log", "l", "", "log file")
    
    Viper.BindPFlag("server.daemon", Flags.Lookup("daemon"))
    Viper.BindPFlag("server.pid_file", Flags.Lookup("pid"))
    Viper.BindPFlag("server.log_file", Flags.Lookup("log"))
    
    // add command
    addCommand(MonitorCmd)
}