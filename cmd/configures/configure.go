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
            PidFile:  Viper.GetString("server.pid_file"),
            SocketFile: Viper.GetString("server.socket_file"),
            LogFile: Viper.GetString("server.log_file"),
        },
        Client: Client{
            SocketFile: Viper.GetString("client.socket_file"),
        },
        MongoDB: Database{
            Host: Viper.GetString("mongodb.host"),
            Port: int8(Viper.GetInt("mongodb.port")),
            AuthDB: Viper.GetString("mongodb.auth_db"),
            Username: Viper.GetString("mongodb.username"),
            Password: Viper.GetString("mongodb.password"),
        },
    }
}
