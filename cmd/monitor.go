package main

import (
    "os"
    "fmt"
    "net"
    "encoding/json"
    "monitor/cmd/commands"
    "monitor/cmd/protocols"
    "monitor/cmd/configures"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/spf13/pflag"
)

var (
    ConfFile string
    
    Viper = viper.GetViper()
)

func Socket(Conf *configures.Conf) (*net.UnixConn, error) {
    os.Remove(Conf.Client.UnixFile)
    
    lAddr, err := net.ResolveUnixAddr("unix", Conf.Client.UnixFile)
    if err != nil {
        return nil, err
    }
    
    rAddr, err := net.ResolveUnixAddr("unix", Conf.Server.UnixFile)
    if err != nil {
        return nil, err
    }
    
    return net.DialUnix("unix", lAddr, rAddr)
}

func bindFlag(Flags *pflag.FlagSet) {
    Flags.StringVarP(&ConfFile, "config", "c", "/etc/monitor.toml", "configuration file specifying additional options")
}

func addCommand(Cmd *cobra.Command) {
    Cmd.AddCommand(commands.VersionCmd)
}

func main() {
    
    RootCmd := &cobra.Command{
        Use: "monitor",
        Short: "Linux server status monitor",
        Long: "Powerful Linux server status monitor client",
        RunE: func(cmd *cobra.Command, args []string) error {
            Conf := configures.Initialize(Viper, ConfFile)
            
            Con, err := Socket(Conf)
            if err != nil {
                return err
            }
            
            // TODO remove this example
            Message, _ := json.Marshal(protocols.Socket{
                Command: "test",
                Body: []byte(""),
                Timestamp: 1234567890,
            })
            fmt.Println(string(Message))
            Con.Write(Message)
            
            return nil
        },
    }
    
    bindFlag(RootCmd.Flags())
    
    addCommand(RootCmd)
    
    if err := RootCmd.Execute(); err != nil {
        fmt.Printf("%v", err)
        os.Exit(-1)
    }
}
