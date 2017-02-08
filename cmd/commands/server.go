package commands

import (
    "encoding/json"
    "github.com/spf13/cobra"
    "monitor/cmd/configures"
    "monitor/utils"
    "monitor/cmd/protocols"
)

var MongoDB = configures.Database{}

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
        
        Body, _ := json.Marshal(Conf.MongoDB)
        
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
    
    Flags.StringVarP(&MongoDB.Host, "host", "", "127.0.0.1", "mongodb host")
    Flags.Int32VarP(&MongoDB.Port, "port", "p", 27017, "mongodb port")
    Flags.StringVarP(&MongoDB.AuthDB, "auth", "a", "admin", "auth database")
    Flags.StringVarP(&MongoDB.Username, "user", "u", "", "username")
    Flags.StringVarP(&MongoDB.Password, "pwd", "", "", "password")
    
    Viper.BindPFlag("mongodb.host", Flags.Lookup("host"))
    Viper.BindPFlag("mongodb.port", Flags.Lookup("port"))
    Viper.BindPFlag("mongodb.auth_db", Flags.Lookup("authDatabase"))
    Viper.BindPFlag("mongodb.username", Flags.Lookup("user"))
    Viper.BindPFlag("mongodb.password", Flags.Lookup("pwd"))
    
    ServerCmd.AddCommand(initCmd)
    ServerCmd.AddCommand(tokenCmd)
}
