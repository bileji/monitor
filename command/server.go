package command

import (
    "strings"
    "strconv"
    "github.com/spf13/cobra"
    "gopkg.in/mgo.v2"
    "monitor/daemon"
    "monitor/service"
)

var (
    serverPort int
    serverInterface string
    
    authUri string
)

type DBAuth struct {
    Host     string
    Port     int
    Database string
    Username string
    Password string
}

var serverCmd = &cobra.Command{
    Use: "server",
    Aliases: []string{"serve"},
    Short: "monitor the service side",
    Long: ``,
    Run: func(cmd *cobra.Command, args []string) {
        auth := authUriToDBAuth(authUri)
        Session, err := mgo.Dial(auth.Host + ":" + strconv.Itoa(auth.Port))
        if err != nil {
            panic(err)
        }
        if Session.DB(auth.Database).Login(auth.Username, auth.Password) != nil {
            panic(err)
        }
        Daemon := &daemon.Daemon{
            PidFile: "/var/run/monitord.pid",
            UnixFile: "/var/run/monitord.sock",
            LogFile: "/var/log/monitord.log",
        }
        Daemon.Daemon(func(ch chan []byte) {
            defer Session.Close()
            
            go (&service.Master{
                Addr: serverInterface + ":" + strconv.Itoa(serverPort),
                DBHandler: Session.DB(auth.Database),
            }).Listen()
            
            // todo socket
        })
    },
}

func authUriToDBAuth(AuthUri string) DBAuth {
    var dbAuth = DBAuth{
        Username: "root",
        Password: "",
    }
    
    if arrStr := strings.Split(AuthUri, "/"); len(arrStr) >= 1 {
        if len(arrStr) > 1 {
            authInfo := strings.Split(arrStr[1], "@")
            if len(authInfo) > 1 {
                loginInfo := strings.Split(authInfo[1], ":")
                dbAuth.Username = loginInfo[0]
                if len(loginInfo) > 1 {
                    dbAuth.Password = loginInfo[1]
                }
            }
            dbAuth.Database = authInfo[0]
        }
        
        hostInfo := strings.Split(arrStr[0], ":")
        if len(hostInfo) > 1 {
            dbAuth.Port, _ = strconv.Atoi(hostInfo[1])
        }
        dbAuth.Host = hostInfo[0]
    }
    
    return dbAuth
}

func init() {
    serverCmd.Flags().IntVarP(&serverPort, "port", "p", 3413, "port on which the server will listen")
    serverCmd.Flags().StringVarP(&serverInterface, "bind", "b", "0.0.0.0", "interface to which the server will bind")
    
    serverCmd.Flags().StringVarP(&authUri, "db_auth", "a", "127.0.0.1:27017/vpn@shadowsocks:mlgR4evB", "auth uri of the mongodb")
}