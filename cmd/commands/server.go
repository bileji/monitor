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
        
        Conf.MongoDB = parseDBUri(authUri)
        
        Body, _ := json.Marshal(Conf.MongoDB)
        
        Message, _ := json.Marshal(protocols.Socket{
            Command: "serverinit",
            Body: Body,
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
    var Database = &configures.Database{
        Username: "root",
        Password: "",
    }
    
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
}

//import (
//    "fmt"
//    "net"
//    "strings"
//    "strconv"
//    "github.com/spf13/cobra"
//    "gopkg.in/mgo.v2"
//    "monitor/monitor/daemon"
//    "monitor/monitor/webserver"
//)
//
//var (
//    serverPort int
//    serverInterface string
//
//    authUri string
//)
//
//type DBAuth struct {
//    Host     string
//    Port     int
//    Database string
//    Username string
//    Password string
//}
//
//var serverCmd = &cobra.Command{
//    Use: "server",
//    Aliases: []string{"serve"},
//    Short: "monitor the service side",
//    Long: ``,
//    Run: func(cmd *cobra.Command, args []string) {
//        auth := authUriToDBAuth(authUri)
//
//        fmt.Println(auth)
//
//        Session, err := mgo.Dial(auth.Host + ":" + strconv.Itoa(auth.Port))
//        if err != nil {
//            panic(err)
//        }
//        if Session.DB(auth.Database).Login(auth.Username, auth.Password) != nil {
//            panic(err)
//        }
//        Daemon := &daemon.Daemon{
//            PidFile: "/var/run/monitord.pid",
//            UnixFile: "/var/run/monitord.sock",
//            LogFile: "/var/log/monitord.log",
//        }
//        Daemon.Daemon(func(Unix *net.UnixListener) {
//            defer Session.Close()
//
//            go (&service.Master{
//                Addr: serverInterface + ":" + strconv.Itoa(serverPort),
//                DBHandler: Session.DB(auth.Database),
//            }).Listen()
//
//            // todo socket
//        })
//    },
//}
//
