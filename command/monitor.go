package command

import (
    "fmt"
    "monitor/command/common"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

//var MainCmd = &common.Command{
//    Viper: viper.GetViper(),
//    Command: &cobra.Command{
//        Use: common.CMD_MAIN,
//        Short: "Linux server status monitor",
//        Long: "Powerful Linux server status monitor server",
//        PersistentPreRunE: pre,
//        RunE: func(cmd *cobra.Command, args []string) error {
//            fmt.Println("run")
//            return nil
//        },
//    },
//    Children: []*common.Command{
//        VersionCmd,
//        //RoleCmd,
//        JoinCmd,
//        //ServerCmd,
//    },
//}

var MainCmd = &common.Command{
    Viper: viper.GetViper(),
    Children: []*common.Command{
        VersionCmd,
        //RoleCmd,
        JoinCmd,
        //ServerCmd,
    },
}

func init() {
    
    MainCmd.Command = &cobra.Command{
        Use: common.CMD_MAIN,
        Short: "Linux server status monitor",
        Long: "Powerful Linux server status monitor server",
        PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
            fmt.Println(common.ReadConf(MainCmd.Viper, MainCmd.Flags.Main.Config))
            return nil
        },
        RunE: func(cmd *cobra.Command, args []string) error {
            fmt.Println("run")
            return nil
        },
    }
    
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