package commands

import (
    "log"
    "net"
    "fmt"
    "encoding/json"
    "monitor/monitor"
    "monitor/monitor/daemon"
    "monitor/cmd/protocols"
    "monitor/cmd/configures"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "monitor/utils"
)

type filePath struct {
    Log  string
    Conf string
    Pid  string
}

type MonitorCmd struct {
    Daemon    bool
    
    RunE      func(cmd *cobra.Command, args []string) error
    
    File      *filePath
    Viper     *viper.Viper
    WebRole   *utils.SyncVar
    WebServer *monitor.WebServer
}

func (mc *MonitorCmd) Scheduler(Listener *net.UnixListener) {
    defer Listener.Close()
    
    var socket = func(Con *net.UnixConn) {
        for {
            Buffer := make([]byte, protocols.READ_LENGTH)
            Len, err := Con.Read(Buffer);
            if err != nil {
                Con.Close()
                return
            }
            var Message protocols.Socket
            json.Unmarshal(Buffer[0: Len], &Message)
            mc.Dispatcher(Message, Con)
        }
    }
    
    for {
        if UnixConn, err := Listener.AcceptUnix(); err != nil {
            log.Printf("%v\r\n", err)
        } else {
            go socket(UnixConn)
        }
    }
}

func (mc *MonitorCmd) Dispatcher(Msg protocols.Socket, Conn *net.UnixConn) {
    const (
        RN int = 0  // 未设置
        RM int = 1  // 管理
        RS int = 2  // 节点
    )
    
    // 查看角色
    if Msg.Command == protocols.ROLE {
        var OutPut []byte
        switch mc.WebRole.Get() {
        case RN:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: 0,
                Body: []byte("uninitialized"),
            })
        case RM:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: 0,
                Body: []byte("manager"),
            })
        case RS:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: 0,
                Body: []byte("node"),
            })
        default:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: 0,
                Body: []byte("undefined"),
            })
        }
        Conn.Write(OutPut)
        return
    }
    
    // 初始化manager
    if Msg.Command == protocols.SERVER_INIT {
        var OutPut []byte
        switch mc.WebRole.Get() {
        case RN:
            json.Unmarshal(Msg.Body, &mc.WebServer)
            
            if len(mc.WebServer.Token) <= 0 {
                mc.WebServer.Token = utils.RandStr()
            }
            
            err := (&monitor.Monitor{}).ServerInit(mc.WebServer)
            if err == nil {
                OutPut, _ = json.Marshal(protocols.OutPut{
                    Status: 0,
                    Body: []byte("success"),
                })
                mc.WebRole.Set(RM)
                break
            }
            
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: -1,
                Body: []byte(fmt.Sprint(err)),
            })
        case RM:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: -2,
                Body: []byte("server has been initialized as manager"),
            })
        case RS:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: -2,
                Body: []byte("server has been initialized as node"),
            })
        default:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: -2,
                Body: []byte("unknown identity"),
            })
        }
        
        Conn.Write(OutPut)
        return
    }
    
    // 查看令牌
    if Msg.Command == protocols.SERVER_TOKEN {
        var OutPut []byte
        switch mc.WebRole.Get() {
        case RN:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: 0,
                Body: []byte("uninitialized"),
            })
        case RS:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: 0,
                Body: []byte("monitor role is node"),
            })
        case RM:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: 0,
                Body: []byte("monitor join --addr " + mc.WebServer.Addr + " --token " + mc.WebServer.Token),
            })
        default:
            OutPut, _ = json.Marshal(protocols.OutPut{
                Status: 0,
                Body: []byte("unknown role"),
            })
        }
        Conn.Write(OutPut)
        return
    }
}

func (mc *MonitorCmd) AddCommand(Cmd *cobra.Command) {
    Cmd.AddCommand(RoleCmd)
    Cmd.AddCommand(ServerCmd)
    Cmd.AddCommand(VersionCmd)
}

var Manager = &MonitorCmd{
    File: &filePath{},
    Viper: viper.GetViper(),
    WebRole: &utils.SyncVar{},
    WebServer: &monitor.WebServer{},
}

var MainCmd = &cobra.Command{
    Use: "monitor",
    Short: "Linux server status monitor",
    Long: "Powerful Linux server status monitor server",
    RunE: func(cmd *cobra.Command, args []string) error {
        Conf := configures.Initialize(Manager.Viper, Manager.File.Conf)
        Daemon := &daemon.Daemon{
            PidFile:  Conf.Server.PidFile,
            UnixFile: Conf.Server.UnixFile,
            LogFile:  Conf.Server.LogFile,
        }
        
        if Conf.Server.Daemon {
            Daemon.Daemon(Manager.Scheduler)
            return nil
        }
        
        // todo write pid file
        Daemon.UnixListen(Manager.Scheduler)
        return nil
    },
}

func init() {
    
    Flags := MainCmd.Flags()
    Flags.StringVarP(&Manager.File.Conf, "config", "c", "/etc/monitor.toml", "configuration file specifying additional options")
    
    Flags.BoolVarP(&Manager.Daemon, "daemon", "d", false, "to start the daemon way")
    Flags.StringVarP(&Manager.File.Pid, "pid", "", "", "full path to pidfile")
    Flags.StringVarP(&Manager.File.Log, "log", "l", "", "log file")
    
    Manager.Viper.BindPFlag("server.daemon", Flags.Lookup("daemon"))
    Manager.Viper.BindPFlag("server.pid_file", Flags.Lookup("pid"))
    Manager.Viper.BindPFlag("server.log_file", Flags.Lookup("log"))
    
    MainCmd.SetUsageTemplate(utils.UsageTemplate())
    
    Manager.AddCommand(MainCmd)
}