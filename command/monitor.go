package command

import (
    "monitor/command/common"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var MonitorCmd = &common.Command{
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
    Children: []*cobra.Command{
        VersionCmd,
        //RoleCmd,
        //JoinCmd,
        //ServerCmd,
    },
}

func init() {
    
    Flags := MonitorCmd.Command.Flags()
    Flags.BoolVarP(&MonitorCmd.Flags["daemon"], "daemon", "d", false, "to start by daemon")
    Flags.StringVarP(&MonitorCmd.Flags["pid"], "pid", "", "", "full path of pid file")
    Flags.StringVarP(&MonitorCmd.Flags["log"], "log", "", "", "full path of log file")
    Flags.StringVarP(&MonitorCmd.Flags["config"], "config", "c", "/etc/monitor.toml", "configuration file specifying additional options")
    
    V := MonitorCmd.Viper
    V.BindPFlag("server.pid", Flags.Lookup("pid"))
    V.BindPFlag("server.log", Flags.Lookup("log"))
    V.BindPFlag("server.daemon", Flags.Lookup("daemon"))
    
    MonitorCmd.NewChildren()
}