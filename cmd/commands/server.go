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
        
        // todo 此处有一个bug永远会覆盖配置文件
        //Conf.MongoDB = MongoDB
        
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
//
//func parseDBUri(AuthUri string) *configures.Database {
//    var Database = &configures.Database{}
//
//    if arrStr := strings.Split(AuthUri, "/"); len(arrStr) >= 1 {
//        if len(arrStr) > 1 {
//            authInfo := strings.Split(arrStr[1], "@")
//            if len(authInfo) > 1 {
//                loginInfo := strings.Split(authInfo[1], ":")
//                Database.Username = loginInfo[0]
//                if len(loginInfo) > 1 {
//                    Database.Password = loginInfo[1]
//                }
//            }
//            Database.AuthDB = authInfo[0]
//        }
//
//        hostInfo := strings.Split(arrStr[0], ":")
//        if len(hostInfo) > 1 {
//            Port, _ := strconv.Atoi(hostInfo[1])
//            Database.Port = int32(Port)
//        }
//        Database.Host = hostInfo[0]
//    }
//
//    return Database
//}

func init() {
    Flags := ServerCmd.Flags()
    //Flags.StringVarP(&authUri, "db_uri", "", "127.0.0.1:27017/vpn@shadowsocks:mlgR4evB", "auth uri of the mongodb")
    
    Flags.StringVarP(&MongoDB.Host, "host", "", "127.0.0.1", "mongodb host")
    Flags.Int32VarP(&MongoDB.Port, "port", "", 27017, "mongodb port")
    Flags.StringVarP(&MongoDB.AuthDB, "authDatabase", "", "admin", "auth database")
    Flags.StringVarP(&MongoDB.Username, "user", "", "", "username")
    Flags.StringVarP(&MongoDB.Password, "pwd", "", "", "password")
    
    Viper.BindPFlag("mongodb.host", Flags.Lookup("host"))
    Viper.BindPFlag("mongodb.port", Flags.Lookup("port"))
    Viper.BindPFlag("mongodb.auth_db", Flags.Lookup("authDatabase"))
    Viper.BindPFlag("mongodb.username", Flags.Lookup("user"))
    Viper.BindPFlag("mongodb.password", Flags.Lookup("pwd"))
    
    ServerCmd.AddCommand(initCmd)
    ServerCmd.AddCommand(tokenCmd)
}
