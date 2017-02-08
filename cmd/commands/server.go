package commands

import (
    "strings"
    "strconv"
    "encoding/json"
    "github.com/spf13/cobra"
    "monitor/cmd/configures"
    "monitor/utils"
    "monitor/cmd/protocols"
)

var (
    authUri string
)

var initCmd = &cobra.Command{
    Use: "init",
    RunE: func(cmd *cobra.Command, args []string) error {
        Conf := configures.Initialize(Viper, ConfFile)
        
        Conn, err := utils.UnixSocket(Conf)
        if err != nil {
            return err
        }
        
        // todo 此处有一个bug永远会覆盖配置文件
        Conf.MongoDB = *(parseDBUri(authUri))
        
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
        
        return nil
    },
}

var ServerCmd = &cobra.Command{
    Use: "server",
    Aliases: []string{"serve"},
    //RunE: func(cmd *cobra.Command, args []string) {},
}

func parseDBUri(AuthUri string) *configures.Database {
    var Database = &configures.Database{}
    
    if arrStr := strings.Split(AuthUri, "/"); len(arrStr) >= 1 {
        if len(arrStr) > 1 {
            authInfo := strings.Split(arrStr[1], "@")
            if len(authInfo) > 1 {
                loginInfo := strings.Split(authInfo[1], ":")
                Database.Username = loginInfo[0]
                if len(loginInfo) > 1 {
                    Database.Password = loginInfo[1]
                }
            }
            Database.AuthDB = authInfo[0]
        }
        
        hostInfo := strings.Split(arrStr[0], ":")
        if len(hostInfo) > 1 {
            Port, _ := strconv.Atoi(hostInfo[1])
            Database.Port = int16(Port)
        }
        Database.Host = hostInfo[0]
    }
    
    return Database
}

func init() {
    Flags := ServerCmd.Flags()
    
    Flags.StringVarP(&authUri, "db_uri", "", "127.0.0.1:27017/vpn@shadowsocks:mlgR4evB", "auth uri of the mongodb")
    
    ServerCmd.AddCommand(initCmd)
    ServerCmd.AddCommand(tokenCmd)
}
