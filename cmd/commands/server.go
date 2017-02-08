package commands

import (
    "encoding/json"
    "github.com/spf13/cobra"
    "monitor/cmd/configures"
    "monitor/utils"
    "monitor/cmd/protocols"
    "monitor/monitor"
)

var webServer = monitor.WebServer{}

var initCmd = &cobra.Command{
    Use: "init",
    Short: "initialize a server",
    Long: "Initialize a server",
    RunE: func(cmd *cobra.Command, args []string) error {
        Conf := configures.Initialize(Viper, ConfFile)
        
        Conn, err := utils.UnixSocket(Conf)
        if err != nil {
            return err
        }
        
        Body, _ := json.Marshal(webServer)
        
        Message, _ := json.Marshal(protocols.Socket{
            Command: protocols.SERVER_INIT,
            Body: Body,
            Timestamp: utils.UnixTime(),
        })
        Conn.Write(Message)
        
        utils.ParseOutPut(Conn)
        
        return nil
    },
}

var tokenCmd = &cobra.Command{
    Use: "token",
    Short: "show join token",
    Long: "Show join token",
    RunE: func(cmd *cobra.Command, args []string) error {
        Conf := configures.Initialize(Viper, ConfFile)
        
        Conn, err := utils.UnixSocket(Conf)
        if err != nil {
            return err
        }
        
        Message, _ := json.Marshal(protocols.Socket{
            Command: protocols.SERVER_TOKEN,
            Body: []byte(""),
            Timestamp: utils.UnixTime(),
        })
        Conn.Write(Message)
        
        utils.ParseOutPut(Conn)
        
        return nil
    },
}

var ServerCmd = &cobra.Command{
    Use: "server",
    Aliases: []string{"serve"},
    Short: "manage Server",
    Long: "Manage Server",
}

func init() {
    Flags := initCmd.Flags()
    
    Flags.StringVarP(&webServer.Database.Host, "host", "", "127.0.0.1", "mongodb host")
    Flags.Int32VarP(&webServer.Database.Port, "port", "", 27017, "mongodb port")
    Flags.StringVarP(&webServer.Database.AuthDB, "auth", "", "admin", "auth database")
    Flags.StringVarP(&webServer.Database.Username, "user", "", "", "username")
    Flags.StringVarP(&webServer.Database.Password, "pwd", "", "", "password")
    
    Flags.StringVarP(&webServer.Addr, "addr", "", ":3647", "web server address")
    
    Viper.BindPFlag("mongodb.host", Flags.Lookup("host"))
    Viper.BindPFlag("mongodb.port", Flags.Lookup("port"))
    Viper.BindPFlag("mongodb.auth_db", Flags.Lookup("auth"))
    Viper.BindPFlag("mongodb.username", Flags.Lookup("user"))
    Viper.BindPFlag("mongodb.password", Flags.Lookup("pwd"))
    
    ServerCmd.AddCommand(initCmd)
    ServerCmd.AddCommand(tokenCmd)
}
