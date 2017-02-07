package main

import "monitor/cmd/commands"

func main() {
    commands.Monitor()
}
//
//import (
//    "os"
//    "fmt"
//    "encoding/json"
//    "monitor/utils"
//    "monitor/cmd/commands"
//    "monitor/cmd/protocols"
//    "monitor/cmd/configures"
//    "github.com/spf13/cobra"
//    "github.com/spf13/viper"
//    "github.com/spf13/pflag"
//)
//
//var (
//    ConfFile string
//
//    Viper = viper.GetViper()
//)
//
//func bindFlag(Flags *pflag.FlagSet) {
//    Flags.StringVarP(&ConfFile, "config", "c", "/etc/monitor.toml", "configuration file specifying additional options")
//}
//
//func addCommand(Cmd *cobra.Command) {
//    Cmd.AddCommand(commands.ServerCmd)
//    Cmd.AddCommand(commands.VersionCmd)
//}
//
//func main() {
//
//    RootCmd := &cobra.Command{
//        Use: "monitor",
//        Short: "Linux server status monitor",
//        Long: "Powerful Linux server status monitor client",
//        RunE: func(cmd *cobra.Command, args []string) error {
//            Conf := configures.Initialize(Viper, ConfFile)
//
//            Con, err := utils.UnixSocket(Conf)
//            if err != nil {
//                return err
//            }
//
//            // TODO remove this example
//            Message, _ := json.Marshal(protocols.Socket{
//                Command: "test",
//                Body: []byte(""),
//                Timestamp: 1234567890,
//            })
//            fmt.Println(string(Message))
//            Con.Write(Message)
//
//            return nil
//        },
//    }
//
//    bindFlag(RootCmd.Flags())
//
//    addCommand(RootCmd)
//
//    if err := RootCmd.Execute(); err != nil {
//        fmt.Printf("%v", err)
//        os.Exit(-1)
//    }
//}
