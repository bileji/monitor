package configures

import (
    "fmt"
    "strings"
    "path/filepath"
    "github.com/spf13/viper"
)

type Server struct {
    PidFile    string
    SocketFile string
    LogFile    string
}

type Client struct {
    SocketFile string
}

type Database struct {
    Host     string
    Port     int8
    AuthDB   string
    Username string
    Password string
}

type Conf struct {
    Server  Server
    Client  Client
    MongoDB Database
}

func Initialize(Viper *viper.Viper, Path string) *Conf {
    Dir, File := filepath.Split(Path)
    Ext := filepath.Ext(File)
    
    Viper.SetConfigName(strings.Replace(File, Ext, "", 1))
    Viper.SetConfigType(strings.Trim(Ext, "."))
    Viper.AddConfigPath(Dir)
    
    if Viper.ReadInConfig() != nil {
        fmt.Println("No config file found. Using built-in defaults.")
    }
    
    return &Conf{
        Server: Server{
            PidFile:  Viper.Get("server.pid_file"),
            SocketFile: Viper.Get("server.socket_file"),
            LogFile: Viper.Get("server.log_file"),
        },
        Client: Client{
            SocketFile: Viper.Get("client.socket_file"),
        },
        MongoDB: Database{
            Host: Viper.Get("mongodb.host"),
            Port: Viper.Get("mongodb.port"),
            AuthDB: Viper.Get("mongodb.auth_db"),
            Username: Viper.Get("mongodb.username"),
            Password: Viper.Get("mongodb.password"),
        },
    }
}
