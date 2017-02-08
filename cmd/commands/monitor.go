package commands

import (
    "log"
    "net"
    "sync"
    "encoding/json"
    "monitor/monitor"
    "monitor/monitor/daemon"
    "monitor/cmd/protocols"
    "monitor/cmd/configures"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    Daemon bool
    
    ConfFile string
    
    PidFile string
    LogFile string
    
    Viper = viper.GetViper()
    
    Serve = func(cmd *cobra.Command, args []string) error {
        Conf := configures.Initialize(Viper, ConfFile)
        Daemon := &daemon.Daemon{
            PidFile: Conf.Server.PidFile,
            UnixFile: Conf.Server.UnixFile,
            LogFile: Conf.Server.LogFile,
        }
        
        if Conf.Server.Daemon {
            Daemon.Daemon(scheduler)
        } else {
            Daemon.UnixListen(scheduler)
        }
        
        return nil
    }
)

type serverStatus struct {
    sync.RWMutex
    Online bool
}

func (O *serverStatus) Get() bool {
    O.Lock()
    defer O.Unlock()
    return O.Online
}

func (O *serverStatus) Set(Status bool) {
    O.Lock()
    defer O.Unlock()
    O.Online = Status
}

var sStatus = serverStatus{Online: false}

func scheduler(Unix *net.UnixListener) {
    defer Unix.Close()
    
    for {
        if UnixConn, err := Unix.AcceptUnix(); err != nil {
            log.Printf("%v\r\n", err)
        } else {
            go func(Con *net.UnixConn) {
                for {
                    Buffer := make([]byte, protocols.READ_LENGTH)
                    Len, err := Con.Read(Buffer);
                    if err != nil {
                        Con.Close()
                        return
                    }
                    var Message protocols.Socket
                    json.Unmarshal(Buffer[0: Len], &Message)
                    dispatcher(Message, Con)
                }
            }(UnixConn)
        }
    }
}

// todo 接收到cli信息,然后处理
func dispatcher(Msg protocols.Socket, Con *net.UnixConn) {
    if Msg.Command == protocols.SERVER_INIT {
        if sStatus.Get() == false {
            var DB configures.Database
            json.Unmarshal(Msg.Body, &DB)
            M := &monitor.Monitor{}
            
            err := M.ServerInit(&monitor.WebServer{
                Addr: ":3647",
                Database: DB,
            })
            if err != nil {
                OutPut, _ := json.Marshal(protocols.OutPut{
                    Status: -1,
                    Body: []byte("failure"),
                })
                Con.Write(OutPut)
            } else {
                OutPut, _ := json.Marshal(protocols.OutPut{
                    Status: 0,
                    Body: []byte("success"),
                })
                Con.Write(OutPut)
            }
            sStatus.Set(true)
        } else {
            OutPut, _ := json.Marshal(protocols.OutPut{
                Status: -1,
                Body: []byte("server has been inited"),
            })
            Con.Write(OutPut)
        }
    }
}

func addCommand(Cmd *cobra.Command) {
    Cmd.AddCommand(ServerCmd)
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