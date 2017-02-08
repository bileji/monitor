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
    "monitor/utils"
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

type serverRole struct {
    sync.RWMutex
    // 0 未设置 1 server 2 slave
    Role int
}

func (O *serverRole) Get() int {
    O.Lock()
    defer O.Unlock()
    return O.Role
}

func (O *serverRole) Set(Role int) {
    O.Lock()
    defer O.Unlock()
    O.Role = Role
}

var Role = serverRole{Role: 0}

var WebServer = &monitor.WebServer{}

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
    const (
        // monitor角色
        RN int = 0  // 未设置
        RM int = 1  // master
        RS int = 2  // slave
    )
    
    // server init
    if Msg.Command == protocols.SERVER_INIT {
        if Role.Get() == RN {
            json.Unmarshal(Msg.Body, &WebServer.Database)
            WebServer.Addr = ":3647"
            WebServer.Token = utils.RandStr()
            
            if (&monitor.Monitor{}).ServerInit(WebServer) != nil {
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
            Role.Set(RM)
        } else {
            OutPut, _ := json.Marshal(protocols.OutPut{
                Status: -2,
                Body: []byte("server has been inited"),
            })
            Con.Write(OutPut)
        }
        return
    }
    
    // server token
    if Msg.Command == protocols.SERVER_TOKEN {
        if Role.Get() != RM {
            // todo
            return
        }
        OutPut, _ := json.Marshal(protocols.OutPut{
            Status: 0,
            Body: []byte(WebServer.Token),
        })
        Con.Write(OutPut)
        return
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
    
    MonitorCmd.UsageTemplate = func() {
        return `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine "[flags]"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
  {{ .CommandPath}} [command]{{end}}{{if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableSubCommands }}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
    }
    
    // add command
    addCommand(MonitorCmd)
}