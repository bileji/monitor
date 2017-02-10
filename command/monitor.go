package command

import (
    "monitor/command/common"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var MainCmd = &common.Command{
    Viper: viper.GetViper(),
    Command: &cobra.Command{
        Use: common.CMD_MAIN,
        Short: "Linux server status monitor",
        Long: "Powerful Linux server status monitor server",
        PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
            
            return nil
        },
        RunE: func(cmd *cobra.Command, args []string) error {
            
            return nil
        },
    },
    Children: []*common.Command{
        VersionCmd,
        //RoleCmd,
        JoinCmd,
        //ServerCmd,
    },
}

func init() {
    
    Flags := MainCmd.Command.Flags()
    Flags.BoolVarP(&MainCmd.Flags.Main.Daemon, "daemon", "d", false, "to start by daemon")
    Flags.StringVarP(&MainCmd.Flags.Main.Pid, "pid", "", "", "full path of pid file")
    Flags.StringVarP(&MainCmd.Flags.Main.Log, "log", "", "", "full path of log file")
    Flags.StringVarP(&MainCmd.Flags.Main.Config, "config", "c", "/etc/monitor.toml", "configuration file specifying additional options")
    
    V := MainCmd.Viper
    V.BindPFlag("server.pid", Flags.Lookup("pid"))
    V.BindPFlag("server.log", Flags.Lookup("log"))
    V.BindPFlag("server.daemon", Flags.Lookup("daemon"))
    
    MainCmd.NewChildren()
}